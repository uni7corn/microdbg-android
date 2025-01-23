package internal

import (
	"hash/fnv"
	"sync"
	"unsafe"

	android "github.com/wnxd/microdbg-android"
	"github.com/wnxd/microdbg-android/internal/jni"
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg/debugger"
)

type Environ struct {
	APK     Package
	JNI     android.JNIEnv
	cf      gava.ClassFactory
	dbg     debugger.Debugger
	vm      gava.FakeJavaVM
	ex      java.IThrowable
	classes sync.Map
	methods sync.Map
	fields  sync.Map
}

func (env *Environ) Init(dbg debugger.Debugger) (err error) {
	env.cf = gava.NewClassFactory(nil)
	env.dbg = dbg
	env.vm, err = gava.NewJavaVM(dbg, jni.NewJavaVM(env))
	return
}

func (env *Environ) Close() error {
	if env.APK != nil {
		env.APK.Close()
	}
	return env.vm.Close()
}

func (env *Environ) JavaVM() java.JavaVM {
	return env.vm
}

func (env *Environ) Package() android.Package {
	return env.APK
}

func (env *Environ) ClassFactory() gava.ClassFactory {
	return env.cf
}

func (env *Environ) DefineClass(name string, loader java.IObject, buf []java.JByte) java.IClass {
	if env.JNI == nil {
		return nil
	}
	return env.cf.WrapClass(must(env.JNI.DefineClass(env, name, loader, buf)))
}

func (env *Environ) FindClass(name string) java.IClass {
	h := fnv1a(name)
	if val, ok := env.classes.Load(h); ok {
		return val.(java.IClass)
	} else if env.JNI == nil {
		return nil
	}
	cls, err := env.JNI.FindClass(env, name)
	if err == nil {
		cls = env.cf.WrapClass(cls)
		env.classes.Store(h, cls)
	}
	return cls
}

func (env *Environ) Throw(obj java.IThrowable) java.JInt {
	env.ex = obj
	return java.JNI_OK
}

func (env *Environ) ThrowNew(clazz java.IClass, message string) java.JInt {
	if env.JNI == nil {
		return java.JNI_ERR
	}
	return must(env.JNI.ThrowNew(env, clazz, message))
}

func (env *Environ) ExceptionOccurred() java.IThrowable {
	return env.ex
}

func (env *Environ) ExceptionDescribe() {
}

func (env *Environ) ExceptionClear() {
	env.ex = nil
}

func (env *Environ) FatalError(string) {
}

func (env *Environ) AllocObject(clazz java.IClass) java.IObject {
	return clazz.NewInstance()
}

func (env *Environ) NewObject(clazz java.IClass, method java.IMethod, args ...any) java.IObject {
	return method.Call(clazz, args...)
}

func (env *Environ) GetMethod(clazz java.IClass, name string, sig string) java.IMethod {
	h := uint64(clazz.HashCode()) ^ fnv1a(name) ^ fnv1a(sig)
	if val, ok := env.methods.Load(h); ok {
		return val.(java.IMethod)
	}
	fake := clazz.(gava.FakeClass)
	if method := fake.FindMethod(name, sig); method != nil {
		return method
	} else if env.JNI == nil {
		return nil
	}
	method, err := env.JNI.GetMethod(env, clazz, name, sig)
	if err == nil {
		env.methods.Store(h, method)
	}
	return method
}

func (env *Environ) CallObjectMethod(obj java.IObject, method java.IMethod, args ...any) java.IObject {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		return gava.ToObject[java.IObject](call.CallMethod(method, args...))
	}
	return method.Call(obj, args...)
}

func (env *Environ) CallBooleanMethod(obj java.IObject, method java.IMethod, args ...any) java.JBoolean {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		return call.CallMethod(method, args...).(java.JBoolean)
	}
	return method.CallPrimitive(obj, args...).(java.JBoolean)
}

func (env *Environ) CallByteMethod(obj java.IObject, method java.IMethod, args ...any) java.JByte {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		return call.CallMethod(method, args...).(java.JByte)
	}
	return method.CallPrimitive(obj, args...).(java.JByte)
}

func (env *Environ) CallCharMethod(obj java.IObject, method java.IMethod, args ...any) java.JChar {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		return call.CallMethod(method, args...).(java.JChar)
	}
	return method.CallPrimitive(obj, args...).(java.JChar)
}

func (env *Environ) CallShortMethod(obj java.IObject, method java.IMethod, args ...any) java.JShort {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		return call.CallMethod(method, args...).(java.JShort)
	}
	return method.CallPrimitive(obj, args...).(java.JShort)
}

func (env *Environ) CallIntMethod(obj java.IObject, method java.IMethod, args ...any) java.JInt {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		return call.CallMethod(method, args...).(java.JInt)
	}
	return method.CallPrimitive(obj, args...).(java.JInt)
}

func (env *Environ) CallLongMethod(obj java.IObject, method java.IMethod, args ...any) java.JLong {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		return call.CallMethod(method, args...).(java.JLong)
	}
	return method.CallPrimitive(obj, args...).(java.JLong)
}

func (env *Environ) CallFloatMethod(obj java.IObject, method java.IMethod, args ...any) java.JFloat {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		return call.CallMethod(method, args...).(java.JFloat)
	}
	return method.CallPrimitive(obj, args...).(java.JFloat)
}

func (env *Environ) CallDoubleMethod(obj java.IObject, method java.IMethod, args ...any) java.JDouble {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		return call.CallMethod(method, args...).(java.JDouble)
	}
	return method.CallPrimitive(obj, args...).(java.JDouble)
}

func (env *Environ) CallVoidMethod(obj java.IObject, method java.IMethod, args ...any) {
	if call, ok := obj.(interface {
		CallMethod(java.IMethod, ...any) any
	}); ok {
		call.CallMethod(method, args...)
	} else {
		method.CallPrimitive(obj, args...)
	}
}

func (env *Environ) CallNonvirtualObjectMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) java.IObject {
	return method.Call(obj, args...)
}

func (env *Environ) CallNonvirtualBooleanMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) java.JBoolean {
	return method.CallPrimitive(obj, args...).(java.JBoolean)
}

func (env *Environ) CallNonvirtualByteMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) java.JByte {
	return method.CallPrimitive(obj, args...).(java.JByte)
}

func (env *Environ) CallNonvirtualCharMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) java.JChar {
	return method.CallPrimitive(obj, args...).(java.JChar)
}

func (env *Environ) CallNonvirtualShortMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) java.JShort {
	return method.CallPrimitive(obj, args...).(java.JShort)
}

func (env *Environ) CallNonvirtualIntMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) java.JInt {
	return method.CallPrimitive(obj, args...).(java.JInt)
}

func (env *Environ) CallNonvirtualLongMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) java.JLong {
	return method.CallPrimitive(obj, args...).(java.JLong)
}

func (env *Environ) CallNonvirtualFloatMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) java.JFloat {
	return method.CallPrimitive(obj, args...).(java.JFloat)
}

func (env *Environ) CallNonvirtualDoubleMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) java.JDouble {
	return method.CallPrimitive(obj, args...).(java.JDouble)
}

func (env *Environ) CallNonvirtualVoidMethod(obj java.IObject, clazz java.IClass, method java.IMethod, args ...any) {
	method.CallPrimitive(obj, args...)
}

func (env *Environ) GetField(clazz java.IClass, name string, sig string) java.IField {
	h := uint64(clazz.HashCode()) ^ ^fnv1a(name) ^ fnv1a(sig)
	if val, ok := env.fields.Load(h); ok {
		return val.(java.IField)
	} else if env.JNI == nil {
		return nil
	}
	field, err := env.JNI.GetField(env, clazz, name, sig)
	if err == nil {
		env.fields.Store(h, field)
	}
	return field
}

func (env *Environ) GetObjectField(obj java.IObject, field java.IField) java.IObject {
	return field.Get(obj)
}

func (env *Environ) GetBooleanField(obj java.IObject, field java.IField) java.JBoolean {
	return field.GetPrimitive(obj).(java.JBoolean)
}

func (env *Environ) GetByteField(obj java.IObject, field java.IField) java.JByte {
	return field.GetPrimitive(obj).(java.JByte)
}

func (env *Environ) GetCharField(obj java.IObject, field java.IField) java.JChar {
	return field.GetPrimitive(obj).(java.JChar)
}

func (env *Environ) GetShortField(obj java.IObject, field java.IField) java.JShort {
	return field.GetPrimitive(obj).(java.JShort)
}

func (env *Environ) GetIntField(obj java.IObject, field java.IField) java.JInt {
	return field.GetPrimitive(obj).(java.JInt)
}

func (env *Environ) GetLongField(obj java.IObject, field java.IField) java.JLong {
	return field.GetPrimitive(obj).(java.JLong)
}

func (env *Environ) GetFloatField(obj java.IObject, field java.IField) java.JFloat {
	return field.GetPrimitive(obj).(java.JFloat)
}

func (env *Environ) GetDoubleField(obj java.IObject, field java.IField) java.JDouble {
	return field.GetPrimitive(obj).(java.JDouble)
}

func (env *Environ) SetObjectField(obj java.IObject, field java.IField, value java.IObject) {
	field.Set(obj, value)
}

func (env *Environ) SetBooleanField(obj java.IObject, field java.IField, value java.JBoolean) {
	field.SetPrimitive(obj, value)
}

func (env *Environ) SetByteField(obj java.IObject, field java.IField, value java.JByte) {
	field.SetPrimitive(obj, value)
}

func (env *Environ) SetCharField(obj java.IObject, field java.IField, value java.JChar) {
	field.SetPrimitive(obj, value)
}

func (env *Environ) SetShortField(obj java.IObject, field java.IField, value java.JShort) {
	field.SetPrimitive(obj, value)
}

func (env *Environ) SetIntField(obj java.IObject, field java.IField, value java.JInt) {
	field.SetPrimitive(obj, value)
}

func (env *Environ) SetLongField(obj java.IObject, field java.IField, value java.JLong) {
	field.SetPrimitive(obj, value)
}

func (env *Environ) SetFloatField(obj java.IObject, field java.IField, value java.JFloat) {
	field.SetPrimitive(obj, value)
}

func (env *Environ) SetDoubleField(obj java.IObject, field java.IField, value java.JDouble) {
	field.SetPrimitive(obj, value)
}

func (env *Environ) GetStaticMethod(clazz java.IClass, name string, sig string) java.IMethod {
	h := uint64(clazz.HashCode()) ^ ^fnv1a(name) ^ fnv1a(sig)
	if val, ok := env.methods.Load(h); ok {
		return val.(java.IMethod)
	}
	fake := clazz.(gava.FakeClass)
	if method := fake.FindMethod(name, sig); method != nil {
		return method
	} else if env.JNI == nil {
		return nil
	}
	method, err := env.JNI.GetStaticMethod(env, clazz, name, sig)
	if err == nil {
		env.methods.Store(h, method)
	}
	return method
}

func (env *Environ) CallStaticObjectMethod(clazz java.IClass, method java.IMethod, args ...any) java.IObject {
	return method.Call(clazz, args...)
}

func (env *Environ) CallStaticBooleanMethod(clazz java.IClass, method java.IMethod, args ...any) java.JBoolean {
	return method.CallPrimitive(clazz, args...).(java.JBoolean)
}

func (env *Environ) CallStaticByteMethod(clazz java.IClass, method java.IMethod, args ...any) java.JByte {
	return method.CallPrimitive(clazz, args...).(java.JByte)
}

func (env *Environ) CallStaticCharMethod(clazz java.IClass, method java.IMethod, args ...any) java.JChar {
	return method.CallPrimitive(clazz, args...).(java.JChar)
}

func (env *Environ) CallStaticShortMethod(clazz java.IClass, method java.IMethod, args ...any) java.JShort {
	return method.CallPrimitive(clazz, args...).(java.JShort)
}

func (env *Environ) CallStaticIntMethod(clazz java.IClass, method java.IMethod, args ...any) java.JInt {
	return method.CallPrimitive(clazz, args...).(java.JInt)
}

func (env *Environ) CallStaticLongMethod(clazz java.IClass, method java.IMethod, args ...any) java.JLong {
	return method.CallPrimitive(clazz, args...).(java.JLong)
}

func (env *Environ) CallStaticFloatMethod(clazz java.IClass, method java.IMethod, args ...any) java.JFloat {
	return method.CallPrimitive(clazz, args...).(java.JFloat)
}

func (env *Environ) CallStaticDoubleMethod(clazz java.IClass, method java.IMethod, args ...any) java.JDouble {
	return method.CallPrimitive(clazz, args...).(java.JDouble)
}

func (env *Environ) CallStaticVoidMethod(clazz java.IClass, method java.IMethod, args ...any) {
	method.CallPrimitive(clazz, args...)
}

func (env *Environ) GetStaticField(clazz java.IClass, name string, sig string) java.IField {
	h := uint64(clazz.HashCode()) ^ ^fnv1a(name) ^ fnv1a(sig)
	if val, ok := env.fields.Load(h); ok {
		return val.(java.IField)
	} else if env.JNI == nil {
		return nil
	}
	field, err := env.JNI.GetStaticField(env, clazz, name, sig)
	if err == nil {
		env.fields.Store(h, field)
	}
	return field
}

func (env *Environ) GetStaticObjectField(clazz java.IClass, field java.IField) java.IObject {
	return field.Get(clazz)
}

func (env *Environ) GetStaticBooleanField(clazz java.IClass, field java.IField) java.JBoolean {
	return field.GetPrimitive(clazz).(java.JBoolean)
}

func (env *Environ) GetStaticByteField(clazz java.IClass, field java.IField) java.JByte {
	return field.GetPrimitive(clazz).(java.JByte)
}

func (env *Environ) GetStaticCharField(clazz java.IClass, field java.IField) java.JChar {
	return field.GetPrimitive(clazz).(java.JChar)
}

func (env *Environ) GetStaticShortField(clazz java.IClass, field java.IField) java.JShort {
	return field.GetPrimitive(clazz).(java.JShort)
}

func (env *Environ) GetStaticIntField(clazz java.IClass, field java.IField) java.JInt {
	return field.GetPrimitive(clazz).(java.JInt)
}

func (env *Environ) GetStaticLongField(clazz java.IClass, field java.IField) java.JLong {
	return field.GetPrimitive(clazz).(java.JLong)
}

func (env *Environ) GetStaticFloatField(clazz java.IClass, field java.IField) java.JFloat {
	return field.GetPrimitive(clazz).(java.JFloat)
}

func (env *Environ) GetStaticDoubleField(clazz java.IClass, field java.IField) java.JDouble {
	return field.GetPrimitive(clazz).(java.JDouble)
}

func (env *Environ) SetStaticObjectField(clazz java.IClass, field java.IField, value java.IObject) {
	field.Set(clazz, value)
}

func (env *Environ) SetStaticBooleanField(clazz java.IClass, field java.IField, value java.JBoolean) {
	field.SetPrimitive(clazz, value)
}

func (env *Environ) SetStaticByteField(clazz java.IClass, field java.IField, value java.JByte) {
	field.SetPrimitive(clazz, value)
}

func (env *Environ) SetStaticCharField(clazz java.IClass, field java.IField, value java.JChar) {
	field.SetPrimitive(clazz, value)
}

func (env *Environ) SetStaticShortField(clazz java.IClass, field java.IField, value java.JShort) {
	field.SetPrimitive(clazz, value)
}

func (env *Environ) SetStaticIntField(clazz java.IClass, field java.IField, value java.JInt) {
	field.SetPrimitive(clazz, value)
}

func (env *Environ) SetStaticLongField(clazz java.IClass, field java.IField, value java.JLong) {
	field.SetPrimitive(clazz, value)
}

func (env *Environ) SetStaticFloatField(clazz java.IClass, field java.IField, value java.JFloat) {
	field.SetPrimitive(clazz, value)
}

func (env *Environ) SetStaticDoubleField(clazz java.IClass, field java.IField, value java.JDouble) {
	field.SetPrimitive(clazz, value)
}

func (env *Environ) NewString(chars []java.JChar) java.IString {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewString(env, chars))
}

func (env *Environ) NewStringUTF(bytes string) java.IString {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewStringUTF(env, bytes))
}

func (env *Environ) NewObjectArray(length java.JSize, elementClass java.IClass, initialElement java.IObject) java.IGenericArray[java.IObject] {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewObjectArray(env, length, elementClass, initialElement))
}

func (env *Environ) NewBooleanArray(length java.JSize) java.IGenericArray[java.JBoolean] {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewBooleanArray(env, length))
}

func (env *Environ) NewByteArray(length java.JSize) java.IGenericArray[java.JByte] {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewByteArray(env, length))
}

func (env *Environ) NewCharArray(length java.JSize) java.IGenericArray[java.JChar] {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewCharArray(env, length))
}

func (env *Environ) NewShortArray(length java.JSize) java.IGenericArray[java.JShort] {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewShortArray(env, length))
}

func (env *Environ) NewIntArray(length java.JSize) java.IGenericArray[java.JInt] {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewIntArray(env, length))
}

func (env *Environ) NewLongArray(length java.JSize) java.IGenericArray[java.JLong] {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewLongArray(env, length))
}

func (env *Environ) NewFloatArray(length java.JSize) java.IGenericArray[java.JFloat] {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewFloatArray(env, length))
}

func (env *Environ) NewDoubleArray(length java.JSize) java.IGenericArray[java.JDouble] {
	if env.JNI == nil {
		return nil
	}
	return must(env.JNI.NewDoubleArray(env, length))
}

func (env *Environ) RegisterNatives(clazz java.IClass, methods []java.JNINativeMethod) java.JInt {
	fake := clazz.(gava.FakeClass)
	for i := range methods {
		method := &methods[i]
		fakeMethod := fake.DefineMethod(method.Name, method.Signature, gava.Modifier_NATIVE)
		fakeMethod.BindCall(newNativeMethod(env.dbg, env.vm, uint64(method.FnPtr), fakeMethod.GetReturnType().DescriptorString().String()))
	}
	return java.JNI_OK
}

func (env *Environ) UnregisterNatives(clazz java.IClass) java.JInt {
	fake := clazz.(gava.FakeClass)
	fake.ClearNativeMethods()
	return java.JNI_OK
}

func (env *Environ) ExceptionCheck() java.JBoolean {
	return env.ex != nil
}

func must[V any](r V, _ error) V {
	return r
}

func fnv1a(str string) uint64 {
	h := fnv.New64a()
	h.Write(unsafe.Slice(unsafe.StringData(str), len(str)))
	return h.Sum64()
}
