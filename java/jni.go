package java

import (
	"io"
	"iter"
	"reflect"
	"strings"
	"sync"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg-linux/gcc"
	"github.com/wnxd/microdbg/debugger"
)

type FakeJavaVM interface {
	io.Closer
	java.JavaVM
	AttachJNIEnv(debugger.Debugger) (FakeJNIEnv, error)
	DetachJNIEnv(FakeJNIEnv)
	GetJNIEnv(FakeJNIEnv) java.JNIEnv
}

type FakeJNIEnv uintptr

type Ref uintptr

type invokeInterface struct {
	reserved0 uintptr
	reserved1 uintptr
	reserved2 uintptr

	DestroyJavaVM               uintptr
	AttachCurrentThread         uintptr
	DetachCurrentThread         uintptr
	GetEnv                      uintptr
	AttachCurrentThreadAsDaemon uintptr
}

type nativeInterface struct {
	reserved0 uintptr
	reserved1 uintptr
	reserved2 uintptr
	reserved3 uintptr

	GetVersion                    uintptr
	DefineClass                   uintptr
	FindClass                     uintptr
	FromReflectedMethod           uintptr
	FromReflectedField            uintptr
	ToReflectedMethod             uintptr
	GetSuperclass                 uintptr
	IsAssignableFrom              uintptr
	ToReflectedField              uintptr
	Throw                         uintptr
	ThrowNew                      uintptr
	ExceptionOccurred             uintptr
	ExceptionDescribe             uintptr
	ExceptionClear                uintptr
	FatalError                    uintptr
	PushLocalFrame                uintptr
	PopLocalFrame                 uintptr
	NewGlobalRef                  uintptr
	DeleteGlobalRef               uintptr
	DeleteLocalRef                uintptr
	IsSameObject                  uintptr
	NewLocalRef                   uintptr
	EnsureLocalCapacity           uintptr
	AllocObject                   uintptr
	NewObject                     uintptr
	NewObjectV                    uintptr
	NewObjectA                    uintptr
	GetObjectClass                uintptr
	IsInstanceOf                  uintptr
	GetMethodID                   uintptr
	CallObjectMethod              uintptr
	CallObjectMethodV             uintptr
	CallObjectMethodA             uintptr
	CallBooleanMethod             uintptr
	CallBooleanMethodV            uintptr
	CallBooleanMethodA            uintptr
	CallByteMethod                uintptr
	CallByteMethodV               uintptr
	CallByteMethodA               uintptr
	CallCharMethod                uintptr
	CallCharMethodV               uintptr
	CallCharMethodA               uintptr
	CallShortMethod               uintptr
	CallShortMethodV              uintptr
	CallShortMethodA              uintptr
	CallIntMethod                 uintptr
	CallIntMethodV                uintptr
	CallIntMethodA                uintptr
	CallLongMethod                uintptr
	CallLongMethodV               uintptr
	CallLongMethodA               uintptr
	CallFloatMethod               uintptr
	CallFloatMethodV              uintptr
	CallFloatMethodA              uintptr
	CallDoubleMethod              uintptr
	CallDoubleMethodV             uintptr
	CallDoubleMethodA             uintptr
	CallVoidMethod                uintptr
	CallVoidMethodV               uintptr
	CallVoidMethodA               uintptr
	CallNonvirtualObjectMethod    uintptr
	CallNonvirtualObjectMethodV   uintptr
	CallNonvirtualObjectMethodA   uintptr
	CallNonvirtualBooleanMethod   uintptr
	CallNonvirtualBooleanMethodV  uintptr
	CallNonvirtualBooleanMethodA  uintptr
	CallNonvirtualByteMethod      uintptr
	CallNonvirtualByteMethodV     uintptr
	CallNonvirtualByteMethodA     uintptr
	CallNonvirtualCharMethod      uintptr
	CallNonvirtualCharMethodV     uintptr
	CallNonvirtualCharMethodA     uintptr
	CallNonvirtualShortMethod     uintptr
	CallNonvirtualShortMethodV    uintptr
	CallNonvirtualShortMethodA    uintptr
	CallNonvirtualIntMethod       uintptr
	CallNonvirtualIntMethodV      uintptr
	CallNonvirtualIntMethodA      uintptr
	CallNonvirtualLongMethod      uintptr
	CallNonvirtualLongMethodV     uintptr
	CallNonvirtualLongMethodA     uintptr
	CallNonvirtualFloatMethod     uintptr
	CallNonvirtualFloatMethodV    uintptr
	CallNonvirtualFloatMethodA    uintptr
	CallNonvirtualDoubleMethod    uintptr
	CallNonvirtualDoubleMethodV   uintptr
	CallNonvirtualDoubleMethodA   uintptr
	CallNonvirtualVoidMethod      uintptr
	CallNonvirtualVoidMethodV     uintptr
	CallNonvirtualVoidMethodA     uintptr
	GetFieldID                    uintptr
	GetObjectField                uintptr
	GetBooleanField               uintptr
	GetByteField                  uintptr
	GetCharField                  uintptr
	GetShortField                 uintptr
	GetIntField                   uintptr
	GetLongField                  uintptr
	GetFloatField                 uintptr
	GetDoubleField                uintptr
	SetObjectField                uintptr
	SetBooleanField               uintptr
	SetByteField                  uintptr
	SetCharField                  uintptr
	SetShortField                 uintptr
	SetIntField                   uintptr
	SetLongField                  uintptr
	SetFloatField                 uintptr
	SetDoubleField                uintptr
	GetStaticMethodID             uintptr
	CallStaticObjectMethod        uintptr
	CallStaticObjectMethodV       uintptr
	CallStaticObjectMethodA       uintptr
	CallStaticBooleanMethod       uintptr
	CallStaticBooleanMethodV      uintptr
	CallStaticBooleanMethodA      uintptr
	CallStaticByteMethod          uintptr
	CallStaticByteMethodV         uintptr
	CallStaticByteMethodA         uintptr
	CallStaticCharMethod          uintptr
	CallStaticCharMethodV         uintptr
	CallStaticCharMethodA         uintptr
	CallStaticShortMethod         uintptr
	CallStaticShortMethodV        uintptr
	CallStaticShortMethodA        uintptr
	CallStaticIntMethod           uintptr
	CallStaticIntMethodV          uintptr
	CallStaticIntMethodA          uintptr
	CallStaticLongMethod          uintptr
	CallStaticLongMethodV         uintptr
	CallStaticLongMethodA         uintptr
	CallStaticFloatMethod         uintptr
	CallStaticFloatMethodV        uintptr
	CallStaticFloatMethodA        uintptr
	CallStaticDoubleMethod        uintptr
	CallStaticDoubleMethodV       uintptr
	CallStaticDoubleMethodA       uintptr
	CallStaticVoidMethod          uintptr
	CallStaticVoidMethodV         uintptr
	CallStaticVoidMethodA         uintptr
	GetStaticFieldID              uintptr
	GetStaticObjectField          uintptr
	GetStaticBooleanField         uintptr
	GetStaticByteField            uintptr
	GetStaticCharField            uintptr
	GetStaticShortField           uintptr
	GetStaticIntField             uintptr
	GetStaticLongField            uintptr
	GetStaticFloatField           uintptr
	GetStaticDoubleField          uintptr
	SetStaticObjectField          uintptr
	SetStaticBooleanField         uintptr
	SetStaticByteField            uintptr
	SetStaticCharField            uintptr
	SetStaticShortField           uintptr
	SetStaticIntField             uintptr
	SetStaticLongField            uintptr
	SetStaticFloatField           uintptr
	SetStaticDoubleField          uintptr
	NewString                     uintptr
	GetStringLength               uintptr
	GetStringChars                uintptr
	ReleaseStringChars            uintptr
	NewStringUTF                  uintptr
	GetStringUTFLength            uintptr
	GetStringUTFChars             uintptr
	ReleaseStringUTFChars         uintptr
	GetArrayLength                uintptr
	NewObjectArray                uintptr
	GetObjectArrayElement         uintptr
	SetObjectArrayElement         uintptr
	NewBooleanArray               uintptr
	NewByteArray                  uintptr
	NewCharArray                  uintptr
	NewShortArray                 uintptr
	NewIntArray                   uintptr
	NewLongArray                  uintptr
	NewFloatArray                 uintptr
	NewDoubleArray                uintptr
	GetBooleanArrayElements       uintptr
	GetByteArrayElements          uintptr
	GetCharArrayElements          uintptr
	GetShortArrayElements         uintptr
	GetIntArrayElements           uintptr
	GetLongArrayElements          uintptr
	GetFloatArrayElements         uintptr
	GetDoubleArrayElements        uintptr
	ReleaseBooleanArrayElements   uintptr
	ReleaseByteArrayElements      uintptr
	ReleaseCharArrayElements      uintptr
	ReleaseShortArrayElements     uintptr
	ReleaseIntArrayElements       uintptr
	ReleaseLongArrayElements      uintptr
	ReleaseFloatArrayElements     uintptr
	ReleaseDoubleArrayElements    uintptr
	GetBooleanArrayRegion         uintptr
	GetByteArrayRegion            uintptr
	GetCharArrayRegion            uintptr
	GetShortArrayRegion           uintptr
	GetIntArrayRegion             uintptr
	GetLongArrayRegion            uintptr
	GetFloatArrayRegion           uintptr
	GetDoubleArrayRegion          uintptr
	SetBooleanArrayRegion         uintptr
	SetByteArrayRegion            uintptr
	SetCharArrayRegion            uintptr
	SetShortArrayRegion           uintptr
	SetIntArrayRegion             uintptr
	SetLongArrayRegion            uintptr
	SetFloatArrayRegion           uintptr
	SetDoubleArrayRegion          uintptr
	RegisterNatives               uintptr
	UnregisterNatives             uintptr
	MonitorEnter                  uintptr
	MonitorExit                   uintptr
	GetJavaVM                     uintptr
	GetStringRegion               uintptr
	GetStringUTFRegion            uintptr
	GetPrimitiveArrayCritical     uintptr
	ReleasePrimitiveArrayCritical uintptr
	GetStringCritical             uintptr
	ReleaseStringCritical         uintptr
	NewWeakGlobalRef              uintptr
	DeleteWeakGlobalRef           uintptr
	ExceptionCheck                uintptr
	NewDirectByteBuffer           uintptr
	GetDirectBufferAddress        uintptr
	GetDirectBufferCapacity       uintptr
	GetObjectRefType              uintptr
}

type fakeJavaVM struct {
	functions *invokeInterface
}

type fakeJNIEnv struct {
	functions *nativeInterface
}

type javaVM struct {
	fake uintptr
	*vm  `encoding:"ignore"`
}

type vm struct {
	java.JavaVM
	releases []func() error
	env      *env
}

type env struct {
	fake     fakeJNIEnv
	releases []func() error
	maps     sync.Map
	mems     sync.Map
}

type jniEnv struct {
	java.JNIEnv
	release func() error
}

var currentEnvKey int

func NewJavaVM(dbg debugger.Debugger, handler java.JavaVM) (FakeJavaVM, error) {
	env, err := newJNIEnv(dbg)
	if err != nil {
		return nil, err
	}
	fake := fakeJavaVM{functions: new(invokeInterface)}
	jvm := javaVM{vm: &vm{JavaVM: handler, env: env}}
	typ := reflect.TypeOf(jvm.vm)
	funcs := reflect.ValueOf(fake.functions).Elem()
	for field := range rangeField(reflect.TypeFor[invokeInterface]()) {
		if strings.HasPrefix(field.Name, "reserved") {
			continue
		}
		method, _ := typ.MethodByName(field.Name + "_")
		ctrl, err := dbg.AddControl(func(ctx debugger.Context, data any) {
			data.(func(*vm, debugger.Context, any))(jvm.vm, ctx, data)
		}, method.Func.Interface())
		if err != nil {
			jvm.Close()
			return nil, err
		}
		funcs.FieldByIndex(field.Index).SetUint(ctrl.Addr())
		jvm.releases = append(jvm.releases, ctrl.Close)
	}
	addrs, err := dbg.MemImport(fake)
	if err != nil {
		jvm.Close()
		return nil, err
	}
	jvm.fake = uintptr(addrs[0])
	jvm.releases = append(jvm.releases, func() error {
		for _, addr := range addrs {
			dbg.MemFree(addr)
		}
		return nil
	})
	return jvm, nil
}

func newJNIEnv(dbg debugger.Debugger) (*env, error) {
	jni := &env{fake: fakeJNIEnv{functions: new(nativeInterface)}}
	typ := reflect.TypeOf(jni)
	funcs := reflect.ValueOf(jni.fake.functions).Elem()
	for field := range rangeField(reflect.TypeFor[nativeInterface]()) {
		if strings.HasPrefix(field.Name, "reserved") {
			continue
		}
		method, _ := typ.MethodByName(field.Name + "_")
		ctrl, err := dbg.AddControl(func(ctx debugger.Context, data any) {
			data.(func(*env, debugger.Context, any))(jni, ctx, data)
		}, method.Func.Interface())
		if err != nil {
			jni.Close()
			return nil, err
		}
		funcs.FieldByIndex(field.Index).SetUint(ctrl.Addr())
		jni.releases = append(jni.releases, ctrl.Close)
	}
	return jni, nil
}

func (vm *vm) Close() error {
	vm.env.Close()
	for i := len(vm.releases) - 1; i >= 0; i-- {
		vm.releases[i]()
	}
	return nil
}

func (vm *vm) AttachJNIEnv(dbg debugger.Debugger) (FakeJNIEnv, error) {
	var env java.JNIEnv
	vm.AttachCurrentThread(&env, nil)
	return vm.env.getFake(dbg, env)
}

func (vm *vm) DetachJNIEnv(fake FakeJNIEnv) {
	vm.env.clearFake(fake)
	vm.DetachCurrentThread()
}

func (vm *vm) GetJNIEnv(fake FakeJNIEnv) java.JNIEnv {
	if env, ok := vm.env.getEnv(fake); ok {
		return env.(*jniEnv).JNIEnv
	}
	return nil
}

func (vm *vm) DestroyJavaVM_(ctx debugger.Context, _ any) {
	ctx.RetWrite(vm.DestroyJavaVM())
	ctx.Return()
}

func (vm *vm) AttachCurrentThread_(ctx debugger.Context, _ any) {
	_, ok := ctx.LocalLoad(&currentEnvKey)
	if ok {
		ctx.RetWrite(java.JInt(java.JNI_EEXIST))
		ctx.Return()
		return
	}
	var penv, args uintptr
	ctx.ArgExtract(debugger.Calling_Default, nil, &penv, &args)
	var env java.JNIEnv
	ctx.RetWrite(vm.AttachCurrentThread(&env, args))
	dbg := ctx.Debugger()
	if fake, err := vm.env.getFake(dbg, env); err == nil {
		dbg.MemWrite(uint64(penv), fake)
		ctx.LocalStore(&currentEnvKey, fake)
	}
	ctx.Return()
}

func (vm *vm) DetachCurrentThread_(ctx debugger.Context, _ any) {
	if val, ok := ctx.LocalLoad(&currentEnvKey); ok {
		ctx.LocalDelete(&currentEnvKey)
		fake := val.(FakeJNIEnv)
		vm.env.clearFake(fake)
	}
	ctx.RetWrite(vm.DetachCurrentThread())
	ctx.Return()
}

func (vm *vm) GetEnv_(ctx debugger.Context, _ any) {
	var penv uintptr
	var version java.JInt
	ctx.ArgExtract(debugger.Calling_Default, nil, &penv, &version)
	dbg := ctx.Debugger()
	if addr, ok := ctx.LocalLoad(&currentEnvKey); ok {
		dbg.MemWrite(uint64(penv), addr)
		ctx.RetWrite(java.JInt(java.JNI_OK))
	} else {
		var env java.JNIEnv
		ctx.RetWrite(vm.GetEnv(&env, version))
		if fake, err := vm.env.getFake(dbg, env); err == nil {
			dbg.MemWrite(uint64(penv), fake)
			ctx.LocalStore(&currentEnvKey, fake)
		}
	}
	ctx.Return()
}

func (vm *vm) AttachCurrentThreadAsDaemon_(ctx debugger.Context, _ any) {
	_, ok := ctx.LocalLoad(&currentEnvKey)
	if ok {
		ctx.RetWrite(java.JInt(java.JNI_EEXIST))
		ctx.Return()
		return
	}
	var penv, args uintptr
	ctx.ArgExtract(debugger.Calling_Default, nil, &penv, &args)
	var env java.JNIEnv
	ctx.RetWrite(vm.AttachCurrentThreadAsDaemon(&env, args))
	dbg := ctx.Debugger()
	if fake, err := vm.env.getFake(dbg, env); err == nil {
		dbg.MemWrite(uint64(penv), fake)
		ctx.LocalStore(&currentEnvKey, fake)
	}
	ctx.Return()
}

func (env *env) Close() error {
	env.maps.Range(func(key, value any) bool {
		value.(io.Closer).Close()
		return true
	})
	env.maps.Clear()
	for i := len(env.releases) - 1; i >= 0; i-- {
		env.releases[i]()
	}
	env.releases = nil
	return nil
}

func (env *env) getFake(dbg debugger.Debugger, handler java.JNIEnv) (FakeJNIEnv, error) {
	if handler == nil {
		return 0, debugger.ErrArgumentInvalid
	}
	var fake FakeJNIEnv
	env.maps.Range(func(key, value any) bool {
		if value.(*jniEnv).JNIEnv == handler {
			fake = key.(FakeJNIEnv)
			return false
		}
		return true
	})
	if fake != 0 {
		return fake, nil
	}
	addrs, err := dbg.MemImport(env.fake)
	if err != nil {
		return 0, err
	}
	fake = FakeJNIEnv(addrs[0])
	env.maps.Store(fake, &jniEnv{
		JNIEnv: handler,
		release: func() error {
			for _, addr := range addrs {
				dbg.MemFree(addr)
			}
			return nil
		},
	})
	return fake, nil
}

func (env *env) clearFake(fake FakeJNIEnv) {
	if val, ok := env.maps.LoadAndDelete(fake); ok {
		val.(io.Closer).Close()
	}
}

func (env *env) getEnv(fake FakeJNIEnv) (java.JNIEnv, bool) {
	if val, ok := env.maps.Load(fake); ok {
		return val.(java.JNIEnv), true
	}
	return nil, false
}

func (env *env) GetVersion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	ctx.ArgExtract(debugger.Calling_Default, &fake)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetVersion())
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) DefineClass_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var name string
	var loader Ref
	var buf uintptr
	var len java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &name, &loader, &buf, &len)
	if handler, ok := env.getEnv(fake); ok {
		data := make([]java.JByte, len)
		ctx.ToPointer(uint64(buf)).MemReadPtr(uint64(len), unsafe.Pointer(unsafe.SliceData(data)))
		ctx.RetWrite(handler.DefineClass(name, loader, data))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) FindClass_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var name string
	ctx.ArgExtract(debugger.Calling_Default, &fake, &name)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.FindClass(name))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) FromReflectedMethod_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var method Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &method)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.FromReflectedMethod(method))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) FromReflectedField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var field Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &field)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.FromReflectedField(field))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) ToReflectedMethod_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var cls, methodID Ref
	var isStatic java.JBoolean
	ctx.ArgExtract(debugger.Calling_Default, &cls, &methodID, &isStatic)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.ToReflectedMethod(cls, methodID, isStatic))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetSuperclass_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetSuperclass(clazz))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) IsAssignableFrom_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz1, clazz2 Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz1, &clazz2)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.IsAssignableFrom(clazz1, clazz2))
	} else {
		ctx.RetWrite(false)
	}
	ctx.Return()
}

func (env *env) ToReflectedField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var cls, fieldID Ref
	var isStatic java.JBoolean
	ctx.ArgExtract(debugger.Calling_Default, &fake, &cls, &fieldID, &isStatic)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.ToReflectedField(cls, fieldID, isStatic))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) Throw_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.Throw(obj))
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) ThrowNew_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz Ref
	var msg string
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &msg)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.ThrowNew(clazz, msg))
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) ExceptionOccurred_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	ctx.ArgExtract(debugger.Calling_Default, &fake)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.ExceptionOccurred())
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) ExceptionDescribe_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	ctx.ArgExtract(debugger.Calling_Default, &fake)
	if handler, ok := env.getEnv(fake); ok {
		handler.ExceptionDescribe()
	}
	ctx.Return()
}

func (env *env) ExceptionClear_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	ctx.ArgExtract(debugger.Calling_Default, &fake)
	if handler, ok := env.getEnv(fake); ok {
		handler.ExceptionClear()
	}
	ctx.Return()
}

func (env *env) FatalError_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var msg string
	ctx.ArgExtract(debugger.Calling_Default, &fake, &msg)
	if handler, ok := env.getEnv(fake); ok {
		handler.FatalError(msg)
	}
	ctx.Return()
}

func (env *env) PushLocalFrame_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var capacity java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &capacity)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.PushLocalFrame(capacity))
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) PopLocalFrame_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var result Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &result)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.PopLocalFrame(result))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewGlobalRef_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var lobj Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &lobj)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewGlobalRef(lobj))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) DeleteGlobalRef_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var gref Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &gref)
	if handler, ok := env.getEnv(fake); ok {
		handler.DeleteGlobalRef(gref)
	}
	ctx.Return()
}

func (env *env) DeleteLocalRef_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj)
	if handler, ok := env.getEnv(fake); ok {
		handler.DeleteLocalRef(obj)
	}
	ctx.Return()
}

func (env *env) IsSameObject_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj1, obj2 Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj1, &obj2)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.IsSameObject(obj1, obj2))
	} else {
		ctx.RetWrite(false)
	}
	ctx.Return()
}

func (env *env) NewLocalRef_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var ref Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &ref)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewLocalRef(ref))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) EnsureLocalCapacity_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var capacity java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &capacity)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.EnsureLocalCapacity(capacity))
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) AllocObject_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.AllocObject(clazz))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewObject_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewObjectV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewObjectV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.NewObjectV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewObjectA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewObjectA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetObjectClass_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetObjectClass(obj))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) IsInstanceOf_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.IsInstanceOf(obj, clazz))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetMethodID_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz Ref
	var name, sig string
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &name, &sig)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetMethodID(clazz, name, sig))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallObjectMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallObjectMethodV(obj, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallObjectMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallObjectMethodV(obj, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallObjectMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallObjectMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallBooleanMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallBooleanMethodV(obj, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallBooleanMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallBooleanMethodV(obj, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallBooleanMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallBooleanMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallByteMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallByteMethodV(obj, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallByteMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallByteMethodV(obj, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallByteMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallByteMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallCharMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallCharMethodV(obj, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallCharMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallCharMethodV(obj, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallCharMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallCharMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallShortMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallShortMethodV(obj, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallShortMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallShortMethodV(obj, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallShortMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallShortMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallIntMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallIntMethodV(obj, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallIntMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallIntMethodV(obj, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallIntMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallIntMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallLongMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallLongMethodV(obj, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallLongMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallLongMethodV(obj, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallLongMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallLongMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallFloatMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallFloatMethodV(obj, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallFloatMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallFloatMethodV(obj, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallFloatMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallFloatMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallDoubleMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallDoubleMethodV(obj, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallDoubleMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallDoubleMethodV(obj, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallDoubleMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallDoubleMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallVoidMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, methodID Ref
	args.Extract(&fake, &obj, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		handler.CallVoidMethodV(obj, methodID, args)
	}
	ctx.Return()
}

func (env *env) CallVoidMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		handler.CallVoidMethodV(obj, methodID, va)
	}
	ctx.Return()
}

func (env *env) CallVoidMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		handler.CallVoidMethodA(obj, methodID, newJValuePtr(ctx.ToPointer(uint64(args))))
	}
	ctx.Return()
}

func (env *env) CallNonvirtualObjectMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualObjectMethodV(obj, clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualObjectMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallNonvirtualObjectMethodV(obj, clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualObjectMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualObjectMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualBooleanMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualBooleanMethodV(obj, clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualBooleanMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallNonvirtualBooleanMethodV(obj, clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualBooleanMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualObjectMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualByteMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualByteMethodV(obj, clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualByteMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallNonvirtualByteMethodV(obj, clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualByteMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualByteMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualCharMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualCharMethodV(obj, clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualCharMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallNonvirtualCharMethodV(obj, clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualCharMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualCharMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualShortMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualShortMethodV(obj, clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualShortMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallNonvirtualShortMethodV(obj, clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualShortMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualShortMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualIntMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualIntMethodV(obj, clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualIntMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallNonvirtualIntMethodV(obj, clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualIntMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualIntMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualLongMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualLongMethodV(obj, clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualLongMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallNonvirtualLongMethodV(obj, clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualLongMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualLongMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualFloatMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualFloatMethodV(obj, clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualFloatMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallNonvirtualFloatMethodV(obj, clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualFloatMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualFloatMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualDoubleMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualDoubleMethodV(obj, clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualDoubleMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallNonvirtualDoubleMethodV(obj, clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualDoubleMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallNonvirtualDoubleMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualVoidMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	args.Extract(&fake, &obj, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		handler.CallNonvirtualVoidMethodV(obj, clazz, methodID, args)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualVoidMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		handler.CallNonvirtualVoidMethodV(obj, clazz, methodID, va)
	}
	ctx.Return()
}

func (env *env) CallNonvirtualVoidMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		handler.CallNonvirtualVoidMethodA(obj, clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args))))
	}
	ctx.Return()
}

func (env *env) GetFieldID_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz Ref
	var name, sig string
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &name, &sig)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetFieldID(clazz, name, sig))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetObjectField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetObjectField(obj, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetBooleanField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetBooleanField(obj, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetByteField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetByteField(obj, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetCharField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetCharField(obj, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetShortField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetShortField(obj, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetIntField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetIntField(obj, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetLongField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetLongField(obj, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetFloatField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetFloatField(obj, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetDoubleField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetDoubleField(obj, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) SetObjectField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID, value Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetObjectField(obj, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetBooleanField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	var value java.JBoolean
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetBooleanField(obj, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetByteField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	var value java.JByte
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetByteField(obj, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetCharField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	var value java.JChar
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetCharField(obj, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetShortField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	var value java.JShort
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetShortField(obj, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetIntField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	var value java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetIntField(obj, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetLongField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	var value java.JLong
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetLongField(obj, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetFloatField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	var value java.JFloat
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetFloatField(obj, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetDoubleField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj, fieldID Ref
	var value java.JDouble
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetDoubleField(obj, fieldID, value)
	}
	ctx.Return()
}

func (env *env) GetStaticMethodID_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz Ref
	var name, sig string
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &name, &sig)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticMethodID(clazz, name, sig))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticObjectMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticObjectMethodV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticObjectMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallStaticObjectMethodV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticObjectMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticObjectMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticBooleanMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticBooleanMethodV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticBooleanMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallStaticBooleanMethodV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticBooleanMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticBooleanMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticByteMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticByteMethodV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticByteMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallStaticByteMethodV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticByteMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticByteMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticCharMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticCharMethodV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticCharMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallStaticCharMethodV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticCharMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticCharMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticShortMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticShortMethodV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticShortMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallStaticShortMethodV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticShortMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticShortMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticIntMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticIntMethodV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticIntMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallStaticIntMethodV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticIntMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticIntMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticLongMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticLongMethodV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticLongMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallStaticLongMethodV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticLongMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticLongMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticFloatMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticFloatMethodV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticFloatMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallStaticFloatMethodV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticFloatMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticFloatMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticDoubleMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticDoubleMethodV(clazz, methodID, args))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticDoubleMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		ctx.RetWrite(handler.CallStaticDoubleMethodV(clazz, methodID, va))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticDoubleMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.CallStaticDoubleMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args)))))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) CallStaticVoidMethod_(ctx debugger.Context, _ any) {
	args, _ := ctx.GetArgs(debugger.Calling_Default)
	var fake FakeJNIEnv
	var clazz, methodID Ref
	args.Extract(&fake, &clazz, &methodID)
	if handler, ok := env.getEnv(fake); ok {
		handler.CallStaticVoidMethodV(clazz, methodID, args)
	}
	ctx.Return()
}

func (env *env) CallStaticVoidMethodV_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		va, _ := gcc.NewVaList(ctx.Debugger(), uint64(args))
		handler.CallStaticVoidMethodV(clazz, methodID, va)
	}
	ctx.Return()
}

func (env *env) CallStaticVoidMethodA_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, methodID Ref
	var args uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methodID, &args)
	if handler, ok := env.getEnv(fake); ok {
		handler.CallStaticVoidMethodA(clazz, methodID, newJValuePtr(ctx.ToPointer(uint64(args))))
	}
	ctx.Return()
}

func (env *env) GetStaticFieldID_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz Ref
	var name, sig string
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &name, &sig)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticFieldID(clazz, name, sig))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStaticObjectField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticObjectField(clazz, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStaticBooleanField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticBooleanField(clazz, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStaticByteField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticByteField(clazz, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStaticCharField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticCharField(clazz, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStaticShortField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticShortField(clazz, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStaticIntField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticIntField(clazz, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStaticLongField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticLongField(clazz, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStaticFloatField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticFloatField(clazz, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStaticDoubleField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStaticDoubleField(clazz, fieldID))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) SetStaticObjectField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID, value Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetStaticObjectField(clazz, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetStaticBooleanField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	var value java.JBoolean
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetStaticBooleanField(clazz, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetStaticByteField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	var value java.JByte
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetStaticByteField(clazz, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetStaticCharField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	var value java.JChar
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetStaticCharField(clazz, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetStaticShortField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	var value java.JShort
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetStaticShortField(clazz, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetStaticIntField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	var value java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetStaticIntField(clazz, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetStaticLongField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	var value java.JLong
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetStaticLongField(clazz, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetStaticFloatField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	var value java.JFloat
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetStaticFloatField(clazz, fieldID, value)
	}
	ctx.Return()
}

func (env *env) SetStaticDoubleField_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz, fieldID Ref
	var value java.JDouble
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &fieldID, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetStaticDoubleField(clazz, fieldID, value)
	}
	ctx.Return()
}

func (env *env) NewString_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var unicodeChars uintptr
	var size java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &unicodeChars, &size)
	if handler, ok := env.getEnv(fake); ok {
		chars := make([]java.JChar, size)
		ctx.ToPointer(uint64(unicodeChars)).MemReadPtr(uint64(size*2), unsafe.Pointer(unsafe.SliceData(chars)))
		ctx.RetWrite(handler.NewString(chars))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStringLength_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var str Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &str)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStringLength(str))
	} else {
		ctx.RetWrite(java.JSize(0))
	}
	ctx.Return()
}

func (env *env) GetStringChars_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var str Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &str)
	if handler, ok := env.getEnv(fake); ok {
		chars := handler.GetStringChars(str)
		size := uint64(len(chars) * 2)
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(chars)))
			env.mems.Store(addr, chars)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) ReleaseStringChars_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var str Ref
	var chars uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &str, &chars)
	addr := uint64(chars)
	if val, ok := env.mems.LoadAndDelete(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			handler.ReleaseStringChars(str, val.([]java.JChar))
		}
		ctx.Debugger().MemFree(addr)
	}
	ctx.Return()
}

func (env *env) NewStringUTF_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var bytes string
	ctx.ArgExtract(debugger.Calling_Default, &fake, &bytes)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewStringUTF(bytes))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetStringUTFLength_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var str Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &str)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetStringUTFLength(str))
	} else {
		ctx.RetWrite(java.JSize(0))
	}
	ctx.Return()
}

func (env *env) GetStringUTFChars_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var str Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &str)
	if handler, ok := env.getEnv(fake); ok {
		bytes := handler.GetStringUTFChars(str)
		size := uint64(len(bytes))
		addr, err := ctx.Debugger().MemAlloc(size + 1)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(bytes)))
			env.mems.Store(addr, bytes)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) ReleaseStringUTFChars_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var str Ref
	var bytes uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &str, &bytes)
	addr := uint64(bytes)
	if val, ok := env.mems.LoadAndDelete(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			handler.ReleaseStringUTFChars(str, val.([]byte))
		}
		ctx.Debugger().MemFree(addr)
	}
	ctx.Return()
}

func (env *env) GetArrayLength_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetArrayLength(array))
	} else {
		ctx.RetWrite(java.JSize(0))
	}
	ctx.Return()
}

func (env *env) NewObjectArray_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var length java.JSize
	var elementClass, initialElement Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &length, &elementClass, &initialElement)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewObjectArray(length, elementClass, initialElement))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetObjectArrayElement_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var index java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &index)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetObjectArrayElement(array, index))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) SetObjectArrayElement_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var index java.JSize
	var value Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &index, &value)
	if handler, ok := env.getEnv(fake); ok {
		handler.SetObjectArrayElement(array, index, value)
	}
	ctx.Return()
}

func (env *env) NewBooleanArray_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var length java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &length)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewBooleanArray(length))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewByteArray_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var length java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &length)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewByteArray(length))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewCharArray_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var length java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &length)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewCharArray(length))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewShortArray_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var length java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &length)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewShortArray(length))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewIntArray_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var length java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &length)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewIntArray(length))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewLongArray_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var length java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &length)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewLongArray(length))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewFloatArray_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var length java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &length)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewFloatArray(length))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) NewDoubleArray_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var length java.JSize
	ctx.ArgExtract(debugger.Calling_Default, &fake, &length)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewDoubleArray(length))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetBooleanArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		elems := handler.GetBooleanArrayElements(array)
		size := uint64(len(elems))
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(elems)))
			env.mems.Store(addr, elems)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetByteArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		elems := handler.GetByteArrayElements(array)
		size := uint64(len(elems))
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(elems)))
			env.mems.Store(addr, elems)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetCharArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		elems := handler.GetCharArrayElements(array)
		size := uint64(len(elems) * 2)
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(elems)))
			env.mems.Store(addr, elems)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetShortArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		elems := handler.GetShortArrayElements(array)
		size := uint64(len(elems) * 2)
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(elems)))
			env.mems.Store(addr, elems)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetIntArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		elems := handler.GetIntArrayElements(array)
		size := uint64(len(elems) * 4)
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(elems)))
			env.mems.Store(addr, elems)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetLongArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		elems := handler.GetLongArrayElements(array)
		size := uint64(len(elems) * 8)
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(elems)))
			env.mems.Store(addr, elems)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetFloatArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		elems := handler.GetFloatArrayElements(array)
		size := uint64(len(elems) * 4)
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(elems)))
			env.mems.Store(addr, elems)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetDoubleArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		elems := handler.GetDoubleArrayElements(array)
		size := uint64(len(elems) * 8)
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(elems)))
			env.mems.Store(addr, elems)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) ReleaseBooleanArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var elems uintptr
	var mode java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &elems, &mode)
	addr := uint64(elems)
	if val, ok := env.mems.Load(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			elems := val.([]java.JBoolean)
			if mode == 0 || mode == java.JNI_COMMIT {
				ctx.ToPointer(addr).MemReadPtr(uint64(len(elems)), unsafe.Pointer(unsafe.SliceData(elems)))
			}
			handler.ReleaseBooleanArrayElements(array, elems, mode)
		}
		if mode == 0 || mode == java.JNI_ABORT {
			env.mems.Delete(addr)
			ctx.Debugger().MemFree(addr)
		}
	}
	ctx.Return()
}

func (env *env) ReleaseByteArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var elems uintptr
	var mode java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &elems, &mode)
	addr := uint64(elems)
	if val, ok := env.mems.Load(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			elems := val.([]java.JByte)
			if mode == 0 || mode == java.JNI_COMMIT {
				ctx.ToPointer(addr).MemReadPtr(uint64(len(elems)), unsafe.Pointer(unsafe.SliceData(elems)))
			}
			handler.ReleaseByteArrayElements(array, elems, mode)
		}
		if mode == 0 || mode == java.JNI_ABORT {
			env.mems.Delete(addr)
			ctx.Debugger().MemFree(addr)
		}
	}
	ctx.Return()
}

func (env *env) ReleaseCharArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var elems uintptr
	var mode java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &elems, &mode)
	addr := uint64(elems)
	if val, ok := env.mems.Load(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			elems := val.([]java.JChar)
			if mode == 0 || mode == java.JNI_COMMIT {
				ctx.ToPointer(addr).MemReadPtr(uint64(len(elems)*2), unsafe.Pointer(unsafe.SliceData(elems)))
			}
			handler.ReleaseCharArrayElements(array, elems, mode)
		}
		if mode == 0 || mode == java.JNI_ABORT {
			env.mems.Delete(addr)
			ctx.Debugger().MemFree(addr)
		}
	}
	ctx.Return()
}

func (env *env) ReleaseShortArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var elems uintptr
	var mode java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &elems, &mode)
	addr := uint64(elems)
	if val, ok := env.mems.Load(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			elems := val.([]java.JShort)
			if mode == 0 || mode == java.JNI_COMMIT {
				ctx.ToPointer(addr).MemReadPtr(uint64(len(elems)*2), unsafe.Pointer(unsafe.SliceData(elems)))
			}
			handler.ReleaseShortArrayElements(array, elems, mode)
		}
		if mode == 0 || mode == java.JNI_ABORT {
			env.mems.Delete(addr)
			ctx.Debugger().MemFree(addr)
		}
	}
	ctx.Return()
}

func (env *env) ReleaseIntArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var elems uintptr
	var mode java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &elems, &mode)
	addr := uint64(elems)
	if val, ok := env.mems.Load(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			elems := val.([]java.JInt)
			if mode == 0 || mode == java.JNI_COMMIT {
				ctx.ToPointer(addr).MemReadPtr(uint64(len(elems)*4), unsafe.Pointer(unsafe.SliceData(elems)))
			}
			handler.ReleaseIntArrayElements(array, elems, mode)
		}
		if mode == 0 || mode == java.JNI_ABORT {
			env.mems.Delete(addr)
			ctx.Debugger().MemFree(addr)
		}
	}
	ctx.Return()
}

func (env *env) ReleaseLongArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var elems uintptr
	var mode java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &elems, &mode)
	addr := uint64(elems)
	if val, ok := env.mems.Load(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			elems := val.([]java.JLong)
			if mode == 0 || mode == java.JNI_COMMIT {
				ctx.ToPointer(addr).MemReadPtr(uint64(len(elems)*8), unsafe.Pointer(unsafe.SliceData(elems)))
			}
			handler.ReleaseLongArrayElements(array, elems, mode)
		}
		if mode == 0 || mode == java.JNI_ABORT {
			env.mems.Delete(addr)
			ctx.Debugger().MemFree(addr)
		}
	}
	ctx.Return()
}

func (env *env) ReleaseFloatArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var elems uintptr
	var mode java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &elems, &mode)
	addr := uint64(elems)
	if val, ok := env.mems.Load(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			elems := val.([]java.JFloat)
			if mode == 0 || mode == java.JNI_COMMIT {
				ctx.ToPointer(addr).MemReadPtr(uint64(len(elems)*4), unsafe.Pointer(unsafe.SliceData(elems)))
			}
			handler.ReleaseFloatArrayElements(array, elems, mode)
		}
		if mode == 0 || mode == java.JNI_ABORT {
			env.mems.Delete(addr)
			ctx.Debugger().MemFree(addr)
		}
	}
	ctx.Return()
}

func (env *env) ReleaseDoubleArrayElements_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var elems uintptr
	var mode java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &elems, &mode)
	addr := uint64(elems)
	if val, ok := env.mems.Load(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			elems := val.([]java.JDouble)
			if mode == 0 || mode == java.JNI_COMMIT {
				ctx.ToPointer(addr).MemReadPtr(uint64(len(elems)*8), unsafe.Pointer(unsafe.SliceData(elems)))
			}
			handler.ReleaseDoubleArrayElements(array, elems, mode)
		}
		if mode == 0 || mode == java.JNI_ABORT {
			env.mems.Delete(addr)
			ctx.Debugger().MemFree(addr)
		}
	}
	ctx.Return()
}

func (env *env) GetBooleanArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JBoolean, len)
		handler.GetBooleanArrayRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) GetByteArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JByte, len)
		handler.GetByteArrayRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) GetCharArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JChar, len)
		handler.GetCharArrayRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len*2), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) GetShortArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JShort, len)
		handler.GetShortArrayRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len*2), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) GetIntArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JInt, len)
		handler.GetIntArrayRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len*4), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) GetLongArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JLong, len)
		handler.GetLongArrayRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len*8), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) GetFloatArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JFloat, len)
		handler.GetFloatArrayRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len*4), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) GetDoubleArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JDouble, len)
		handler.GetDoubleArrayRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len*8), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) SetBooleanArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JBoolean, len)
		ctx.ToPointer(uint64(buf)).MemReadPtr(uint64(len), unsafe.Pointer(unsafe.SliceData(elems)))
		handler.SetBooleanArrayRegion(array, start, elems)
	}
	ctx.Return()
}

func (env *env) SetByteArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JByte, len)
		ctx.ToPointer(uint64(buf)).MemReadPtr(uint64(len), unsafe.Pointer(unsafe.SliceData(elems)))
		handler.SetByteArrayRegion(array, start, elems)
	}
	ctx.Return()
}

func (env *env) SetCharArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JChar, len)
		ctx.ToPointer(uint64(buf)).MemReadPtr(uint64(len*2), unsafe.Pointer(unsafe.SliceData(elems)))
		handler.SetCharArrayRegion(array, start, elems)
	}
	ctx.Return()
}

func (env *env) SetShortArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JShort, len)
		ctx.ToPointer(uint64(buf)).MemReadPtr(uint64(len*2), unsafe.Pointer(unsafe.SliceData(elems)))
		handler.SetShortArrayRegion(array, start, elems)
	}
	ctx.Return()
}

func (env *env) SetIntArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JInt, len)
		ctx.ToPointer(uint64(buf)).MemReadPtr(uint64(len*4), unsafe.Pointer(unsafe.SliceData(elems)))
		handler.SetIntArrayRegion(array, start, elems)
	}
	ctx.Return()
}

func (env *env) SetLongArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JLong, len)
		ctx.ToPointer(uint64(buf)).MemReadPtr(uint64(len*8), unsafe.Pointer(unsafe.SliceData(elems)))
		handler.SetLongArrayRegion(array, start, elems)
	}
	ctx.Return()
}

func (env *env) SetFloatArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JFloat, len)
		ctx.ToPointer(uint64(buf)).MemReadPtr(uint64(len*4), unsafe.Pointer(unsafe.SliceData(elems)))
		handler.SetFloatArrayRegion(array, start, elems)
	}
	ctx.Return()
}

func (env *env) SetDoubleArrayRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JDouble, len)
		ctx.ToPointer(uint64(buf)).MemReadPtr(uint64(len*8), unsafe.Pointer(unsafe.SliceData(elems)))
		handler.SetDoubleArrayRegion(array, start, elems)
	}
	ctx.Return()
}

func (env *env) RegisterNatives_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz Ref
	var methods uintptr
	var nMethods java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz, &methods, &nMethods)
	if handler, ok := env.getEnv(fake); ok {
		arr := make([]java.JNINativeMethod, nMethods)
		ctx.Debugger().MemExtract(uint64(methods), arr)
		ctx.RetWrite(handler.RegisterNatives(clazz, arr))
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) UnregisterNatives_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var clazz Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &clazz)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.UnregisterNatives(clazz))
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) MonitorEnter_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.MonitorEnter(obj))
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) MonitorExit_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.MonitorExit(obj))
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) GetJavaVM_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var pvm uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &pvm)
	if handler, ok := env.getEnv(fake); ok {
		var vm java.JavaVM
		ctx.RetWrite(handler.GetJavaVM(&vm))
		fake, ok := vm.(FakeJavaVM)
		dbg := ctx.Debugger()
		if !ok {
			var err error
			if fake, err = NewJavaVM(dbg, vm); err == nil {
				env.releases = append(env.releases, fake.Close)
			}
		}
		dbg.MemWrite(uint64(pvm), fake)
	} else {
		ctx.RetWrite(java.JInt(java.JNI_ERR))
	}
	ctx.Return()
}

func (env *env) GetStringRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]java.JChar, len)
		handler.GetStringRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len*2), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) GetStringUTFRegion_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var start, len java.JSize
	var buf uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &start, &len, &buf)
	if handler, ok := env.getEnv(fake); ok {
		elems := make([]byte, len)
		handler.GetStringUTFRegion(array, start, elems)
		ctx.ToPointer(uint64(buf)).MemWritePtr(uint64(len), unsafe.Pointer(unsafe.SliceData(elems)))
	}
	ctx.Return()
}

func (env *env) GetPrimitiveArrayCritical_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array)
	if handler, ok := env.getEnv(fake); ok {
		raw := handler.GetPrimitiveArrayCritical(array)
		addr, err := ctx.Debugger().MemAlloc(uint64(len(raw)))
		if err == nil {
			ctx.ToPointer(addr).MemWrite(raw)
			env.mems.Store(addr, raw)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) ReleasePrimitiveArrayCritical_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var array Ref
	var carray uintptr
	var mode java.JInt
	ctx.ArgExtract(debugger.Calling_Default, &fake, &array, &carray, &mode)
	addr := uint64(carray)
	if val, ok := env.mems.Load(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			raw := val.([]byte)
			if mode == 0 || mode == java.JNI_COMMIT {
				ctx.ToPointer(addr).MemReadPtr(uint64(len(raw)), unsafe.Pointer(unsafe.SliceData(raw)))
			}
			handler.ReleasePrimitiveArrayCritical(array, raw, mode)
		}
		if mode == 0 || mode == java.JNI_ABORT {
			ctx.Debugger().MemFree(uint64(carray))
		}
	}
	ctx.Return()
}

func (env *env) GetStringCritical_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var str Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &str)
	if handler, ok := env.getEnv(fake); ok {
		chars := handler.GetStringCritical(str)
		size := uint64(len(chars) * 2)
		addr, err := ctx.Debugger().MemAlloc(size)
		if err == nil {
			ctx.ToPointer(addr).MemWritePtr(size, unsafe.Pointer(unsafe.SliceData(chars)))
			env.mems.Store(addr, chars)
		}
		ctx.RetWrite(uintptr(addr))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) ReleaseStringCritical_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var str Ref
	var carray uintptr
	ctx.ArgExtract(debugger.Calling_Default, &fake, &str, &carray)
	addr := uint64(carray)
	if val, ok := env.mems.LoadAndDelete(addr); ok {
		if handler, ok := env.getEnv(fake); ok {
			handler.ReleaseStringCritical(str, val.([]java.JChar))
		}
		ctx.Debugger().MemFree(addr)
	}
	ctx.Return()
}

func (env *env) NewWeakGlobalRef_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewWeakGlobalRef(obj))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) DeleteWeakGlobalRef_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj)
	if handler, ok := env.getEnv(fake); ok {
		handler.DeleteWeakGlobalRef(obj)
	}
	ctx.Return()
}

func (env *env) ExceptionCheck_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	ctx.ArgExtract(debugger.Calling_Default, &fake)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.ExceptionCheck())
	} else {
		ctx.RetWrite(false)
	}
	ctx.Return()
}

func (env *env) NewDirectByteBuffer_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var address java.AnyPtr
	var capacity java.JLong
	ctx.ArgExtract(debugger.Calling_Default, &fake, &address, &capacity)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.NewDirectByteBuffer(address, capacity))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetDirectBufferAddress_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var buf Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &buf)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetDirectBufferAddress(buf))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetDirectBufferCapacity_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var buf Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &buf)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetDirectBufferCapacity(buf))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *env) GetObjectRefType_(ctx debugger.Context, _ any) {
	var fake FakeJNIEnv
	var obj Ref
	ctx.ArgExtract(debugger.Calling_Default, &fake, &obj)
	if handler, ok := env.getEnv(fake); ok {
		ctx.RetWrite(handler.GetObjectRefType(obj))
	} else {
		ctx.RetWrite(nil)
	}
	ctx.Return()
}

func (env *jniEnv) Close() error {
	return env.release()
}

func rangeField(typ reflect.Type) iter.Seq[reflect.StructField] {
	return func(yield func(reflect.StructField) bool) {
		count := typ.NumField()
		for i := 0; i < count; i++ {
			if !yield(typ.Field(i)) {
				break
			}
		}
	}
}
