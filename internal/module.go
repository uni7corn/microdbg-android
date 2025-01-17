package internal

import (
	"context"
	"debug/elf"
	"io"

	android "github.com/wnxd/microdbg-android"
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
	addr += m.BaseAddr()
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
