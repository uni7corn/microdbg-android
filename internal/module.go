package internal

import (
	"context"
	"debug/elf"
	"io"
	"strings"

	android "github.com/wnxd/microdbg-android"
	"github.com/wnxd/microdbg-android/internal/jni"
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
	loader "github.com/wnxd/microdbg-loader/elf"
	"github.com/wnxd/microdbg/debugger"
	"github.com/wnxd/microdbg/emulator"
)

type module struct {
	debugger.Module
	art android.Runtime
}

func ModuleOf(m debugger.Module, art android.Runtime) android.Module {
	return &module{Module: m, art: art}
}

func (m *module) Close() error {
	m.art.Debugger().Unload(m.Module)
	return m.Module.Close()
}

func (m *module) FindSymbol(name string) (android.Symbol, error) {
	addr, err := m.Module.FindSymbol(name)
	if err != nil {
		return nil, err
	}
	return NewSymbol(m.art.Debugger(), name, addr), nil
}

func (m *module) Symbols(yield func(android.Symbol) bool) {
	if it, ok := m.Module.(debugger.SymbolIter); ok {
		it.Symbols(func(sym debugger.Symbol) bool {
			return yield(NewSymbol(m.art.Debugger(), sym.Name, sym.Value))
		})
	}
}

func (m *module) CallEntry(ctx context.Context) error {
	sym := NewSymbol(m.art.Debugger(), "start", m.Module.EntryAddr())
	return sym.Call(ctx, debugger.Calling_Default, nil)
}

func (m *module) CallOnLoad(ctx context.Context) (java.JInt, error) {
	sym, err := m.FindSymbol("JNI_OnLoad")
	if err != nil {
		return 0, err
	}
	var r java.JInt
	err = sym.Call(ctx, debugger.Calling_Default, &r, m.art.JavaVM(), nil)
	return r, err
}

func (m *module) FindNativeMethod(vm java.JavaVM, clazz java.IClass, name, sig string) (android.NativeMethod, error) {
	bind := func(method android.NativeMethod) android.NativeMethod {
		return method
	}
	fake, ok := clazz.(gava.FakeClass)
	if ok {
		method := fake.FindMethod(name, sig)
		if method != nil && gava.IsNative(method) {
			return method.CallPrimitive, nil
		}
		bind = func(method android.NativeMethod) android.NativeMethod {
			fake.DefineMethod(name, sig, gava.Modifier_NATIVE).BindCall(method)
			return method
		}
	}
	prefix := nameEscape(clazz, name)
	for sym := range m.Symbols {
		name := sym.Name()
		if !strings.HasPrefix(name, prefix) {
			continue
		}
		args := name[len(prefix):]
		if len(args) == 0 {
			_, ret, _ := strings.Cut(sig, ")")
			return bind(newNativeMethod(m.art.Debugger(), vm, sym.Address(), ret)), nil
		} else if !strings.HasPrefix(args, "__") {
			continue
		}
		arg, ret, _ := strings.Cut(sig, ")")
		if sigEscape(arg[1:]) == args[2:] {
			return bind(newNativeMethod(m.art.Debugger(), vm, sym.Address(), ret)), nil
		}
	}
	return nil, android.ErrMethodNotFound
}

func (m *module) Dump(w io.Writer) error {
	start, size := m.Module.Region()
	_, err := io.Copy(w, io.NewSectionReader(emulator.ToPointer(m.art.Emulator(), start), 0, int64(size)))
	return err
}

func AndroidReloc(dbg debugger.Debugger, module loader.Module) {
	const (
		DT_ANDROID_REL = elf.DT_LOOS + 2 + iota
		DT_ANDROID_RELSZ
		DT_ANDROID_RELA
		DT_ANDROID_RELASZ
	)

	sz := module.DynValue(DT_ANDROID_RELSZ)
	for i, v := range module.DynValue(DT_ANDROID_REL) {
		sr := io.NewSectionReader(dbg.ToPointer(module.BaseAddr()), int64(v), int64(sz[i]))
		var magic [4]byte
		sr.Read(magic[:])
		if string(magic[:]) != "APS2" {
			continue
		}
		switch module.Class() {
		case elf.ELFCLASS32:
			androidRel32(module, sr)
		case elf.ELFCLASS64:
			androidRel64(module, sr)
		}
	}
	sz = module.DynValue(DT_ANDROID_RELASZ)
	for i, v := range module.DynValue(DT_ANDROID_RELA) {
		sr := io.NewSectionReader(dbg.ToPointer(module.BaseAddr()), int64(v), int64(sz[i]))
		var magic [4]byte
		sr.Read(magic[:])
		if string(magic[:]) != "APS2" {
			continue
		}
		switch module.Class() {
		case elf.ELFCLASS32:
			androidRela32(module, sr)
		case elf.ELFCLASS64:
			androidRela64(module, sr)
		}
	}
}

func androidRel32(module loader.Module, r io.Reader) {
	for rela := range newRelocIter[uint32](r) {
		module.Reloc(elf.Rel32{Off: rela.Off, Info: rela.Info})
	}
}

func androidRel64(module loader.Module, r io.Reader) {
	for rela := range newRelocIter[uint64](r) {
		module.Reloc(elf.Rel64{Off: rela.Off, Info: rela.Info})
	}
}

func androidRela32(module loader.Module, r io.Reader) {
	for rela := range newRelocIter[uint32](r) {
		module.Reloc(elf.Rela32{Off: rela.Off, Info: rela.Info, Addend: int32(rela.Addend)})
	}
}

func androidRela64(module loader.Module, r io.Reader) {
	for rela := range newRelocIter[uint64](r) {
		module.Reloc(elf.Rela64{Off: rela.Off, Info: rela.Info, Addend: int64(rela.Addend)})
	}
}

func newNativeMethod(dbg debugger.Debugger, jvm java.JavaVM, addr uint64, returnType string) android.NativeMethod {
	vm, ok := jvm.(gava.FakeJavaVM)
	if !ok {
		return nil
	}
	sym := NewSymbol(dbg, "", addr)
	var cast func(env jni.JNIEnv, v gava.JValue) any
	switch returnType {
	case "Z":
		cast = func(env jni.JNIEnv, v gava.JValue) any { return v.JBoolean() }
	case "B":
		cast = func(env jni.JNIEnv, v gava.JValue) any { return v.JByte() }
	case "C":
		cast = func(env jni.JNIEnv, v gava.JValue) any { return v.JChar() }
	case "S":
		cast = func(env jni.JNIEnv, v gava.JValue) any { return v.JShort() }
	case "I":
		cast = func(env jni.JNIEnv, v gava.JValue) any { return v.JInt() }
	case "J":
		cast = func(env jni.JNIEnv, v gava.JValue) any { return v.JLong() }
	case "F":
		cast = func(env jni.JNIEnv, v gava.JValue) any { return v.JFloat() }
	case "D":
		cast = func(env jni.JNIEnv, v gava.JValue) any { return v.JDouble() }
	case "V":
		cast = func(env jni.JNIEnv, v gava.JValue) any { return nil }
	default:
		cast = func(env jni.JNIEnv, v gava.JValue) any { return env.GetObject(v.JObject()) }
	}
	return func(obj java.IObject, args ...any) any {
		fake, err := vm.AttachJNIEnv(dbg)
		if err != nil {
			panic(err)
		}
		defer vm.DetachJNIEnv(fake)
		env := vm.GetJNIEnv(fake).(jni.JNIEnv)
		for i := range args {
			if obj, ok := args[i].(java.IObject); ok {
				args[i] = env.ObjectRef(obj)
			}
		}
		var r gava.JValue
		err = sym.Call(context.TODO(), debugger.Calling_Default, &r, append([]any{fake, env.ObjectRef(obj)}, args...)...)
		if err != nil {
			panic(err)
		}
		return cast(env, r)
	}
}

func nameEscape(clazz java.IClass, name string) string {
	cn := clazz.GetName().String()
	cn = strings.ReplaceAll(cn, "_", "_1")
	cn = strings.ReplaceAll(cn, ".", "_")
	name = strings.ReplaceAll(name, "_", "_1")
	return "Java_" + cn + "_" + name
}

func sigEscape(sig string) string {
	sig = strings.ReplaceAll(sig, ";", "_2")
	sig = strings.ReplaceAll(sig, "[", "_3")
	return sig
}
