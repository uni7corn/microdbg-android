package jni

import (
	"sync"
	"unicode/utf16"
	"unsafe"

	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

const nilRef = gava.Ref(0)

type JNIEnv interface {
	java.JNIEnv
	ObjectRef(java.IObject) gava.Ref
	GetObject(java.JObject) java.IObject
}

type AndroidEnv interface {
	JavaVM() java.JavaVM
	DefineClass(string, java.IObject, []java.JByte) java.IClass
	FindClass(string) java.IClass
	Throw(java.IThrowable) java.JInt
	ThrowNew(java.IClass, string) java.JInt
	ExceptionOccurred() java.IThrowable
	ExceptionDescribe()
	ExceptionClear()
	FatalError(string)
	AllocObject(java.IClass) java.IObject
	NewObject(java.IClass, java.IMethod, ...any) java.IObject
	GetMethod(java.IClass, string, string) java.IMethod
	CallObjectMethod(java.IObject, java.IMethod, ...any) java.IObject
	CallBooleanMethod(java.IObject, java.IMethod, ...any) java.JBoolean
	CallByteMethod(java.IObject, java.IMethod, ...any) java.JByte
	CallCharMethod(java.IObject, java.IMethod, ...any) java.JChar
	CallShortMethod(java.IObject, java.IMethod, ...any) java.JShort
	CallIntMethod(java.IObject, java.IMethod, ...any) java.JInt
	CallLongMethod(java.IObject, java.IMethod, ...any) java.JLong
	CallFloatMethod(java.IObject, java.IMethod, ...any) java.JFloat
	CallDoubleMethod(java.IObject, java.IMethod, ...any) java.JDouble
	CallVoidMethod(java.IObject, java.IMethod, ...any)
	CallNonvirtualObjectMethod(java.IObject, java.IClass, java.IMethod, ...any) java.IObject
	CallNonvirtualBooleanMethod(java.IObject, java.IClass, java.IMethod, ...any) java.JBoolean
	CallNonvirtualByteMethod(java.IObject, java.IClass, java.IMethod, ...any) java.JByte
	CallNonvirtualCharMethod(java.IObject, java.IClass, java.IMethod, ...any) java.JChar
	CallNonvirtualShortMethod(java.IObject, java.IClass, java.IMethod, ...any) java.JShort
	CallNonvirtualIntMethod(java.IObject, java.IClass, java.IMethod, ...any) java.JInt
	CallNonvirtualLongMethod(java.IObject, java.IClass, java.IMethod, ...any) java.JLong
	CallNonvirtualFloatMethod(java.IObject, java.IClass, java.IMethod, ...any) java.JFloat
	CallNonvirtualDoubleMethod(java.IObject, java.IClass, java.IMethod, ...any) java.JDouble
	CallNonvirtualVoidMethod(java.IObject, java.IClass, java.IMethod, ...any)
	GetField(java.IClass, string, string) java.IField
	GetObjectField(java.IObject, java.IField) java.IObject
	GetBooleanField(java.IObject, java.IField) java.JBoolean
	GetByteField(java.IObject, java.IField) java.JByte
	GetCharField(java.IObject, java.IField) java.JChar
	GetShortField(java.IObject, java.IField) java.JShort
	GetIntField(java.IObject, java.IField) java.JInt
	GetLongField(java.IObject, java.IField) java.JLong
	GetFloatField(java.IObject, java.IField) java.JFloat
	GetDoubleField(java.IObject, java.IField) java.JDouble
	SetObjectField(java.IObject, java.IField, java.IObject)
	SetBooleanField(java.IObject, java.IField, java.JBoolean)
	SetByteField(java.IObject, java.IField, java.JByte)
	SetCharField(java.IObject, java.IField, java.JChar)
	SetShortField(java.IObject, java.IField, java.JShort)
	SetIntField(java.IObject, java.IField, java.JInt)
	SetLongField(java.IObject, java.IField, java.JLong)
	SetFloatField(java.IObject, java.IField, java.JFloat)
	SetDoubleField(java.IObject, java.IField, java.JDouble)
	GetStaticMethod(java.IClass, string, string) java.IMethod
	CallStaticObjectMethod(java.IClass, java.IMethod, ...any) java.IObject
	CallStaticBooleanMethod(java.IClass, java.IMethod, ...any) java.JBoolean
	CallStaticByteMethod(java.IClass, java.IMethod, ...any) java.JByte
	CallStaticCharMethod(java.IClass, java.IMethod, ...any) java.JChar
	CallStaticShortMethod(java.IClass, java.IMethod, ...any) java.JShort
	CallStaticIntMethod(java.IClass, java.IMethod, ...any) java.JInt
	CallStaticLongMethod(java.IClass, java.IMethod, ...any) java.JLong
	CallStaticFloatMethod(java.IClass, java.IMethod, ...any) java.JFloat
	CallStaticDoubleMethod(java.IClass, java.IMethod, ...any) java.JDouble
	CallStaticVoidMethod(java.IClass, java.IMethod, ...any)
	GetStaticField(java.IClass, string, string) java.IField
	GetStaticObjectField(java.IClass, java.IField) java.IObject
	GetStaticBooleanField(java.IClass, java.IField) java.JBoolean
	GetStaticByteField(java.IClass, java.IField) java.JByte
	GetStaticCharField(java.IClass, java.IField) java.JChar
	GetStaticShortField(java.IClass, java.IField) java.JShort
	GetStaticIntField(java.IClass, java.IField) java.JInt
	GetStaticLongField(java.IClass, java.IField) java.JLong
	GetStaticFloatField(java.IClass, java.IField) java.JFloat
	GetStaticDoubleField(java.IClass, java.IField) java.JDouble
	SetStaticObjectField(java.IClass, java.IField, java.IObject)
	SetStaticBooleanField(java.IClass, java.IField, java.JBoolean)
	SetStaticByteField(java.IClass, java.IField, java.JByte)
	SetStaticCharField(java.IClass, java.IField, java.JChar)
	SetStaticShortField(java.IClass, java.IField, java.JShort)
	SetStaticIntField(java.IClass, java.IField, java.JInt)
	SetStaticLongField(java.IClass, java.IField, java.JLong)
	SetStaticFloatField(java.IClass, java.IField, java.JFloat)
	SetStaticDoubleField(java.IClass, java.IField, java.JDouble)
	NewString([]java.JChar) java.IString
	NewStringUTF(string) java.IString
	NewObjectArray(java.JSize, java.IClass, java.IObject) java.IGenericArray[java.IObject]
	NewBooleanArray(java.JSize) java.IGenericArray[java.JBoolean]
	NewByteArray(java.JSize) java.IGenericArray[java.JByte]
	NewCharArray(java.JSize) java.IGenericArray[java.JChar]
	NewShortArray(java.JSize) java.IGenericArray[java.JShort]
	NewIntArray(java.JSize) java.IGenericArray[java.JInt]
	NewLongArray(java.JSize) java.IGenericArray[java.JLong]
	NewFloatArray(java.JSize) java.IGenericArray[java.JFloat]
	NewDoubleArray(java.JSize) java.IGenericArray[java.JDouble]
	RegisterNatives(java.IClass, []java.JNINativeMethod) java.JInt
	UnregisterNatives(java.IClass) java.JInt
	ExceptionCheck() java.JBoolean
}

type javaVM struct {
	env *jniEnv
}

type jniEnv struct {
	adEnv     AndroidEnv
	localRef  sync.Map
	globalRef *sync.Map
	weakRef   *sync.Map
	methods   sync.Map
	fields    sync.Map
}

func NewJavaVM(env AndroidEnv) java.JavaVM {
	vm := &javaVM{env: &jniEnv{adEnv: env, globalRef: new(sync.Map), weakRef: new(sync.Map)}}
	return vm
}

func (vm *javaVM) DestroyJavaVM() java.JInt {
	return java.JNI_OK
}

func (vm *javaVM) AttachCurrentThread(penv *java.JNIEnv, _ any) java.JInt {
	*penv = vm.newJNIEnv()
	return java.JNI_OK
}

func (vm *javaVM) DetachCurrentThread() java.JInt {
	return java.JNI_OK
}

func (vm *javaVM) GetEnv(penv *java.JNIEnv, _ java.JInt) java.JInt {
	*penv = vm.env
	return java.JNI_OK
}

func (vm *javaVM) AttachCurrentThreadAsDaemon(penv *java.JNIEnv, args any) java.JInt {
	return vm.AttachCurrentThread(penv, args)
}

func (vm *javaVM) newJNIEnv() *jniEnv {
	return &jniEnv{adEnv: vm.env.adEnv, globalRef: vm.env.globalRef, weakRef: vm.env.weakRef}
}

func (env *jniEnv) GetVersion() java.JInt {
	return java.JNI_VERSION_1_6
}

func (env *jniEnv) DefineClass(name string, loader java.JObject, buf []java.JByte) java.JClass {
	return env.ObjectRef(env.adEnv.DefineClass(name, env.GetObject(loader), buf))
}

func (env *jniEnv) FindClass(name string) java.JClass {
	return env.ObjectRef(env.adEnv.FindClass(name))
}

func (env *jniEnv) FromReflectedMethod(method java.JObject) java.JMethodID {
	m, _ := env.GetObject(method).(java.IMethod)
	return env.methodRef(m)
}

func (env *jniEnv) FromReflectedField(field java.JObject) java.JFieldID {
	f, _ := env.GetObject(field).(java.IField)
	return env.fieldRef(f)
}

func (env *jniEnv) ToReflectedMethod(cls java.JClass, methodID java.JMethodID, isStatic java.JBoolean) java.JObject {
	method := env.getMethod(methodID)
	if method == nil {
		return nil
	} else if isStatic != gava.IsStatic(method) {
		return nil
	}
	return env.ObjectRef(method)
}

func (env *jniEnv) GetSuperclass(clazz java.JClass) java.JClass {
	cls := env.getClass(clazz)
	if cls == nil {
		return nil
	}
	return env.ObjectRef(cls.GetSuperclass())
}

func (env *jniEnv) IsAssignableFrom(clazz1, clazz2 java.JClass) java.JBoolean {
	cls1 := env.getClass(clazz1)
	if cls1 == nil {
		return false
	}
	cls2 := env.getClass(clazz2)
	if cls2 == nil {
		return false
	}
	return cls1.IsAssignableFrom(cls2)
}

func (env *jniEnv) ToReflectedField(cls java.JClass, fieldID java.JFieldID, isStatic java.JBoolean) java.JObject {
	field := env.getField(fieldID)
	if field == nil {
		return nil
	} else if isStatic != gava.IsStatic(field) {
		return nil
	}
	return env.ObjectRef(field)
}

func (env *jniEnv) Throw(obj java.JThrowable) java.JInt {
	ex, _ := env.GetObject(obj).(java.IThrowable)
	return env.adEnv.Throw(ex)
}

func (env *jniEnv) ThrowNew(clazz java.JClass, message string) java.JInt {
	return env.adEnv.ThrowNew(env.getClass(clazz), message)
}

func (env *jniEnv) ExceptionOccurred() java.JThrowable {
	return env.ObjectRef(env.adEnv.ExceptionOccurred())
}

func (env *jniEnv) ExceptionDescribe() {
	env.adEnv.ExceptionDescribe()
}

func (env *jniEnv) ExceptionClear() {
	env.adEnv.ExceptionClear()
}

func (env *jniEnv) FatalError(msg string) {
	env.adEnv.FatalError(msg)
}

func (env *jniEnv) PushLocalFrame(java.JInt) java.JInt {
	return java.JNI_OK
}

func (env *jniEnv) PopLocalFrame(result java.JObject) java.JObject {
	return result
}

func (env *jniEnv) NewGlobalRef(obj java.JObject) java.JObject {
	o := env.GetObject(obj)
	if o == nil {
		return nil
	}
	ref := gava.Ref(o.HashCode()<<2 | java.JInt(java.JNIGlobalRefType))
	env.globalRef.Store(ref, o)
	return ref
}

func (env *jniEnv) DeleteGlobalRef(globalRef java.JObject) {
	env.globalRef.Delete(globalRef)
}

func (env *jniEnv) DeleteLocalRef(localRef java.JObject) {
	env.localRef.Delete(localRef)
}

func (env *jniEnv) IsSameObject(ref1, ref2 java.JObject) java.JBoolean {
	return env.GetObject(ref1) == env.GetObject(ref2)
}

func (env *jniEnv) NewLocalRef(ref java.JObject) java.JObject {
	return env.ObjectRef(env.GetObject(ref))
}

func (env *jniEnv) EnsureLocalCapacity(java.JInt) java.JInt {
	return java.JNI_OK
}

func (env *jniEnv) AllocObject(clazz java.JClass) java.JObject {
	return env.ObjectRef(env.adEnv.AllocObject(env.getClass(clazz)))
}

func (env *jniEnv) NewObject(clazz java.JClass, methodID java.JMethodID, args ...any) java.JObject {
	return env.ObjectRef(env.adEnv.NewObject(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...))
}

func (env *jniEnv) NewObjectV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JObject {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.ObjectRef(env.adEnv.NewObject(env.getClass(clazz), method, a...))
	}
	return nil
}

func (env *jniEnv) NewObjectA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JObject {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.ObjectRef(env.adEnv.NewObject(env.getClass(clazz), method, a...))
	}
	return nil
}

func (env *jniEnv) GetObjectClass(obj java.JObject) java.JClass {
	o := env.GetObject(obj)
	if o == nil {
		return nil
	}
	return env.ObjectRef(o.GetClass())
}

func (env *jniEnv) IsInstanceOf(obj java.JObject, clazz java.JClass) java.JBoolean {
	cls := env.getClass(clazz)
	if cls == nil {
		return false
	}
	return cls.IsInstance(env.GetObject(obj))
}

func (env *jniEnv) GetMethodID(clazz java.JClass, name, sig string) java.JMethodID {
	return env.methodRef(env.adEnv.GetMethod(env.getClass(clazz), name, sig))
}

func (env *jniEnv) CallObjectMethod(obj java.JObject, methodID java.JMethodID, args ...any) java.JObject {
	return env.ObjectRef(env.adEnv.CallObjectMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...))
}

func (env *jniEnv) CallObjectMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) java.JObject {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.ObjectRef(env.adEnv.CallObjectMethod(env.GetObject(obj), method, a...))
	}
	return nil
}

func (env *jniEnv) CallObjectMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JObject {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.ObjectRef(env.adEnv.CallObjectMethod(env.GetObject(obj), method, a...))
	}
	return nil
}

func (env *jniEnv) CallBooleanMethod(obj java.JObject, methodID java.JMethodID, args ...any) java.JBoolean {
	return env.adEnv.CallBooleanMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallBooleanMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) java.JBoolean {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallBooleanMethod(env.GetObject(obj), method, a...)
	}
	return false
}

func (env *jniEnv) CallBooleanMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JBoolean {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallBooleanMethod(env.GetObject(obj), method, a...)
	}
	return false
}

func (env *jniEnv) CallByteMethod(obj java.JObject, methodID java.JMethodID, args ...any) java.JByte {
	return env.adEnv.CallByteMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallByteMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) java.JByte {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallByteMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallByteMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JByte {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallByteMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallCharMethod(obj java.JObject, methodID java.JMethodID, args ...any) java.JChar {
	return env.adEnv.CallCharMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallCharMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) java.JChar {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallCharMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallCharMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JChar {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallCharMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallShortMethod(obj java.JObject, methodID java.JMethodID, args ...any) java.JShort {
	return env.adEnv.CallShortMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallShortMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) java.JShort {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallShortMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallShortMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JShort {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallShortMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallIntMethod(obj java.JObject, methodID java.JMethodID, args ...any) java.JInt {
	return env.adEnv.CallIntMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallIntMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) java.JInt {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallIntMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallIntMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JInt {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallIntMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallLongMethod(obj java.JObject, methodID java.JMethodID, args ...any) java.JLong {
	return env.adEnv.CallLongMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallLongMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) java.JLong {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallLongMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallLongMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JLong {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallLongMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallFloatMethod(obj java.JObject, methodID java.JMethodID, args ...any) java.JFloat {
	return env.adEnv.CallFloatMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallFloatMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) java.JFloat {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallFloatMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallFloatMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JFloat {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallFloatMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallDoubleMethod(obj java.JObject, methodID java.JMethodID, args ...any) java.JDouble {
	return env.adEnv.CallDoubleMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallDoubleMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) java.JDouble {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallDoubleMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallDoubleMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JDouble {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallDoubleMethod(env.GetObject(obj), method, a...)
	}
	return 0
}

func (env *jniEnv) CallVoidMethod(obj java.JObject, methodID java.JMethodID, args ...any) {
	env.adEnv.CallVoidMethod(env.GetObject(obj), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallVoidMethodV(obj java.JObject, methodID java.JMethodID, args java.VaList) {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		env.adEnv.CallVoidMethod(env.GetObject(obj), method, a...)
	}
}

func (env *jniEnv) CallVoidMethodA(obj java.JObject, methodID java.JMethodID, args java.TypePtr[java.JValue]) {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		env.adEnv.CallVoidMethod(env.GetObject(obj), method, a...)
	}
}

func (env *jniEnv) CallNonvirtualObjectMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) java.JObject {
	return env.ObjectRef(env.adEnv.CallNonvirtualObjectMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...))
}

func (env *jniEnv) CallNonvirtualObjectMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JObject {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.ObjectRef(env.adEnv.CallNonvirtualObjectMethod(env.GetObject(obj), env.getClass(clazz), method, a...))
	}
	return nil
}

func (env *jniEnv) CallNonvirtualObjectMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JObject {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.ObjectRef(env.adEnv.CallNonvirtualObjectMethod(env.GetObject(obj), env.getClass(clazz), method, a...))
	}
	return nil
}

func (env *jniEnv) CallNonvirtualBooleanMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) java.JBoolean {
	return env.adEnv.CallNonvirtualBooleanMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallNonvirtualBooleanMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JBoolean {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallNonvirtualBooleanMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return false
}

func (env *jniEnv) CallNonvirtualBooleanMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JBoolean {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallNonvirtualBooleanMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return false
}

func (env *jniEnv) CallNonvirtualByteMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) java.JByte {
	return env.adEnv.CallNonvirtualByteMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallNonvirtualByteMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JByte {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallNonvirtualByteMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualByteMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JByte {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallNonvirtualByteMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualCharMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) java.JChar {
	return env.adEnv.CallNonvirtualCharMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallNonvirtualCharMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JChar {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallNonvirtualCharMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualCharMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JChar {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallNonvirtualCharMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualShortMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) java.JShort {
	return env.adEnv.CallNonvirtualShortMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallNonvirtualShortMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JShort {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallNonvirtualShortMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualShortMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JShort {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallNonvirtualShortMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualIntMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) java.JInt {
	return env.adEnv.CallNonvirtualIntMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallNonvirtualIntMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JInt {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallNonvirtualIntMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualIntMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JInt {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallNonvirtualIntMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualLongMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) java.JLong {
	return env.adEnv.CallNonvirtualLongMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallNonvirtualLongMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JLong {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallNonvirtualLongMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualLongMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JLong {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallNonvirtualLongMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualFloatMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) java.JFloat {
	return env.adEnv.CallNonvirtualFloatMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallNonvirtualFloatMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JFloat {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallNonvirtualFloatMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualFloatMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JFloat {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallNonvirtualFloatMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualDoubleMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) java.JDouble {
	return env.adEnv.CallNonvirtualDoubleMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallNonvirtualDoubleMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JDouble {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallNonvirtualDoubleMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualDoubleMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JDouble {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallNonvirtualDoubleMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallNonvirtualVoidMethod(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args ...any) {
	env.adEnv.CallNonvirtualVoidMethod(env.GetObject(obj), env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallNonvirtualVoidMethodV(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.VaList) {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		env.adEnv.CallNonvirtualVoidMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
}

func (env *jniEnv) CallNonvirtualVoidMethodA(obj java.JObject, clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		env.adEnv.CallNonvirtualVoidMethod(env.GetObject(obj), env.getClass(clazz), method, a...)
	}
}

func (env *jniEnv) GetFieldID(clazz java.JClass, name, sig string) java.JFieldID {
	return env.fieldRef(env.adEnv.GetField(env.getClass(clazz), name, sig))
}

func (env *jniEnv) GetObjectField(obj java.JObject, fieldID java.JFieldID) java.JObject {
	return env.ObjectRef(env.adEnv.GetObjectField(env.GetObject(obj), env.getField(fieldID)))
}

func (env *jniEnv) GetBooleanField(obj java.JObject, fieldID java.JFieldID) java.JBoolean {
	return env.adEnv.GetBooleanField(env.GetObject(obj), env.getField(fieldID))
}

func (env *jniEnv) GetByteField(obj java.JObject, fieldID java.JFieldID) java.JByte {
	return env.adEnv.GetByteField(env.GetObject(obj), env.getField(fieldID))
}

func (env *jniEnv) GetCharField(obj java.JObject, fieldID java.JFieldID) java.JChar {
	return env.adEnv.GetCharField(env.GetObject(obj), env.getField(fieldID))
}

func (env *jniEnv) GetShortField(obj java.JObject, fieldID java.JFieldID) java.JShort {
	return env.adEnv.GetShortField(env.GetObject(obj), env.getField(fieldID))
}

func (env *jniEnv) GetIntField(obj java.JObject, fieldID java.JFieldID) java.JInt {
	return env.adEnv.GetIntField(env.GetObject(obj), env.getField(fieldID))
}

func (env *jniEnv) GetLongField(obj java.JObject, fieldID java.JFieldID) java.JLong {
	return env.adEnv.GetLongField(env.GetObject(obj), env.getField(fieldID))
}

func (env *jniEnv) GetFloatField(obj java.JObject, fieldID java.JFieldID) java.JFloat {
	return env.adEnv.GetFloatField(env.GetObject(obj), env.getField(fieldID))
}

func (env *jniEnv) GetDoubleField(obj java.JObject, fieldID java.JFieldID) java.JDouble {
	return env.adEnv.GetDoubleField(env.GetObject(obj), env.getField(fieldID))
}

func (env *jniEnv) SetObjectField(obj java.JObject, fieldID java.JFieldID, value java.JObject) {
	env.adEnv.SetObjectField(env.GetObject(obj), env.getField(fieldID), env.GetObject(value))
}

func (env *jniEnv) SetBooleanField(obj java.JObject, fieldID java.JFieldID, value java.JBoolean) {
	env.adEnv.SetBooleanField(env.GetObject(obj), env.getField(fieldID), value)
}

func (env *jniEnv) SetByteField(obj java.JObject, fieldID java.JFieldID, value java.JByte) {
	env.adEnv.SetByteField(env.GetObject(obj), env.getField(fieldID), value)
}

func (env *jniEnv) SetCharField(obj java.JObject, fieldID java.JFieldID, value java.JChar) {
	env.adEnv.SetCharField(env.GetObject(obj), env.getField(fieldID), value)
}

func (env *jniEnv) SetShortField(obj java.JObject, fieldID java.JFieldID, value java.JShort) {
	env.adEnv.SetShortField(env.GetObject(obj), env.getField(fieldID), value)
}

func (env *jniEnv) SetIntField(obj java.JObject, fieldID java.JFieldID, value java.JInt) {
	env.adEnv.SetIntField(env.GetObject(obj), env.getField(fieldID), value)
}

func (env *jniEnv) SetLongField(obj java.JObject, fieldID java.JFieldID, value java.JLong) {
	env.adEnv.SetLongField(env.GetObject(obj), env.getField(fieldID), value)
}

func (env *jniEnv) SetFloatField(obj java.JObject, fieldID java.JFieldID, value java.JFloat) {
	env.adEnv.SetFloatField(env.GetObject(obj), env.getField(fieldID), value)
}

func (env *jniEnv) SetDoubleField(obj java.JObject, fieldID java.JFieldID, value java.JDouble) {
	env.adEnv.SetDoubleField(env.GetObject(obj), env.getField(fieldID), value)
}

func (env *jniEnv) GetStaticMethodID(clazz java.JClass, name, sig string) java.JMethodID {
	return env.methodRef(env.adEnv.GetStaticMethod(env.getClass(clazz), name, sig))
}

func (env *jniEnv) CallStaticObjectMethod(clazz java.JClass, methodID java.JMethodID, args ...any) java.JObject {
	return env.ObjectRef(env.adEnv.CallStaticObjectMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...))
}

func (env *jniEnv) CallStaticObjectMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JObject {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.ObjectRef(env.adEnv.CallStaticObjectMethod(env.getClass(clazz), method, a...))
	}
	return nil
}

func (env *jniEnv) CallStaticObjectMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JObject {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.ObjectRef(env.adEnv.CallStaticObjectMethod(env.getClass(clazz), method, a...))
	}
	return nil
}

func (env *jniEnv) CallStaticBooleanMethod(clazz java.JClass, methodID java.JMethodID, args ...any) java.JBoolean {
	return env.adEnv.CallStaticBooleanMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallStaticBooleanMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JBoolean {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallStaticBooleanMethod(env.getClass(clazz), method, a...)
	}
	return false
}

func (env *jniEnv) CallStaticBooleanMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JBoolean {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallStaticBooleanMethod(env.getClass(clazz), method, a...)
	}
	return false
}

func (env *jniEnv) CallStaticByteMethod(clazz java.JClass, methodID java.JMethodID, args ...any) java.JByte {
	return env.adEnv.CallStaticByteMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallStaticByteMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JByte {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallStaticByteMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticByteMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JByte {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallStaticByteMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticCharMethod(clazz java.JClass, methodID java.JMethodID, args ...any) java.JChar {
	return env.adEnv.CallStaticCharMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallStaticCharMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JChar {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallStaticCharMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticCharMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JChar {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallStaticCharMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticShortMethod(clazz java.JClass, methodID java.JMethodID, args ...any) java.JShort {
	return env.adEnv.CallStaticShortMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallStaticShortMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JShort {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallStaticShortMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticShortMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JShort {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallStaticShortMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticIntMethod(clazz java.JClass, methodID java.JMethodID, args ...any) java.JInt {
	return env.adEnv.CallStaticIntMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallStaticIntMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JInt {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallStaticIntMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticIntMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JInt {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallStaticIntMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticLongMethod(clazz java.JClass, methodID java.JMethodID, args ...any) java.JLong {
	return env.adEnv.CallStaticLongMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallStaticLongMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JLong {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallStaticLongMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticLongMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JLong {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallStaticLongMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticFloatMethod(clazz java.JClass, methodID java.JMethodID, args ...any) java.JFloat {
	return env.adEnv.CallStaticFloatMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallStaticFloatMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JFloat {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallStaticFloatMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticFloatMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JFloat {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallStaticFloatMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticDoubleMethod(clazz java.JClass, methodID java.JMethodID, args ...any) java.JDouble {
	return env.adEnv.CallStaticDoubleMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallStaticDoubleMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) java.JDouble {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		return env.adEnv.CallStaticDoubleMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticDoubleMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) java.JDouble {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		return env.adEnv.CallStaticDoubleMethod(env.getClass(clazz), method, a...)
	}
	return 0
}

func (env *jniEnv) CallStaticVoidMethod(clazz java.JClass, methodID java.JMethodID, args ...any) {
	env.adEnv.CallStaticVoidMethod(env.getClass(clazz), env.getMethod(methodID), env.deRefArgs(args)...)
}

func (env *jniEnv) CallStaticVoidMethodV(clazz java.JClass, methodID java.JMethodID, args java.VaList) {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractVaList(method, args); err == nil {
		env.adEnv.CallStaticVoidMethod(env.getClass(clazz), method, a...)
	}
}

func (env *jniEnv) CallStaticVoidMethodA(clazz java.JClass, methodID java.JMethodID, args java.TypePtr[java.JValue]) {
	if method := env.getMethod(methodID); method == nil {
	} else if a, err := env.extractPtr(method, args); err == nil {
		env.adEnv.CallStaticVoidMethod(env.getClass(clazz), method, a...)
	}
}

func (env *jniEnv) GetStaticFieldID(clazz java.JClass, name, sig string) java.JFieldID {
	return env.fieldRef(env.adEnv.GetStaticField(env.getClass(clazz), name, sig))
}

func (env *jniEnv) GetStaticObjectField(clazz java.JClass, fieldID java.JFieldID) java.JObject {
	return env.ObjectRef(env.adEnv.GetStaticObjectField(env.getClass(clazz), env.getField(fieldID)))
}

func (env *jniEnv) GetStaticBooleanField(clazz java.JClass, fieldID java.JFieldID) java.JBoolean {
	return env.adEnv.GetStaticBooleanField(env.getClass(clazz), env.getField(fieldID))
}

func (env *jniEnv) GetStaticByteField(clazz java.JClass, fieldID java.JFieldID) java.JByte {
	return env.adEnv.GetStaticByteField(env.getClass(clazz), env.getField(fieldID))
}

func (env *jniEnv) GetStaticCharField(clazz java.JClass, fieldID java.JFieldID) java.JChar {
	return env.adEnv.GetStaticCharField(env.getClass(clazz), env.getField(fieldID))
}

func (env *jniEnv) GetStaticShortField(clazz java.JClass, fieldID java.JFieldID) java.JShort {
	return env.adEnv.GetStaticShortField(env.getClass(clazz), env.getField(fieldID))
}

func (env *jniEnv) GetStaticIntField(clazz java.JClass, fieldID java.JFieldID) java.JInt {
	return env.adEnv.GetStaticIntField(env.getClass(clazz), env.getField(fieldID))
}

func (env *jniEnv) GetStaticLongField(clazz java.JClass, fieldID java.JFieldID) java.JLong {
	return env.adEnv.GetStaticLongField(env.getClass(clazz), env.getField(fieldID))
}

func (env *jniEnv) GetStaticFloatField(clazz java.JClass, fieldID java.JFieldID) java.JFloat {
	return env.adEnv.GetStaticFloatField(env.getClass(clazz), env.getField(fieldID))
}

func (env *jniEnv) GetStaticDoubleField(clazz java.JClass, fieldID java.JFieldID) java.JDouble {
	return env.adEnv.GetStaticDoubleField(env.getClass(clazz), env.getField(fieldID))
}

func (env *jniEnv) SetStaticObjectField(clazz java.JClass, fieldID java.JFieldID, value java.JObject) {
	env.adEnv.SetStaticObjectField(env.getClass(clazz), env.getField(fieldID), env.GetObject(value))
}

func (env *jniEnv) SetStaticBooleanField(clazz java.JClass, fieldID java.JFieldID, value java.JBoolean) {
	env.adEnv.SetBooleanField(env.getClass(clazz), env.getField(fieldID), value)
}

func (env *jniEnv) SetStaticByteField(clazz java.JClass, fieldID java.JFieldID, value java.JByte) {
	env.adEnv.SetStaticByteField(env.getClass(clazz), env.getField(fieldID), value)
}

func (env *jniEnv) SetStaticCharField(clazz java.JClass, fieldID java.JFieldID, value java.JChar) {
	env.adEnv.SetStaticCharField(env.getClass(clazz), env.getField(fieldID), value)
}

func (env *jniEnv) SetStaticShortField(clazz java.JClass, fieldID java.JFieldID, value java.JShort) {
	env.adEnv.SetStaticShortField(env.getClass(clazz), env.getField(fieldID), value)
}

func (env *jniEnv) SetStaticIntField(clazz java.JClass, fieldID java.JFieldID, value java.JInt) {
	env.adEnv.SetStaticIntField(env.getClass(clazz), env.getField(fieldID), value)
}

func (env *jniEnv) SetStaticLongField(clazz java.JClass, fieldID java.JFieldID, value java.JLong) {
	env.adEnv.SetStaticLongField(env.getClass(clazz), env.getField(fieldID), value)
}

func (env *jniEnv) SetStaticFloatField(clazz java.JClass, fieldID java.JFieldID, value java.JFloat) {
	env.adEnv.SetStaticFloatField(env.getClass(clazz), env.getField(fieldID), value)
}

func (env *jniEnv) SetStaticDoubleField(clazz java.JClass, fieldID java.JFieldID, value java.JDouble) {
	env.adEnv.SetStaticDoubleField(env.getClass(clazz), env.getField(fieldID), value)
}

func (env *jniEnv) NewString(chars []java.JChar) java.JString {
	return env.ObjectRef(env.adEnv.NewString(chars))
}

func (env *jniEnv) GetStringLength(str java.JString) java.JSize {
	s := env.getString(str)
	if s == nil {
		return 0
	}
	return s.Length()
}

func (env *jniEnv) GetStringChars(str java.JString) []java.JChar {
	s := env.getString(str)
	if s == nil {
		return nil
	}
	return utf16.Encode([]rune(s.String()))
}

func (env *jniEnv) ReleaseStringChars(java.JString, []java.JChar) {
}

func (env *jniEnv) NewStringUTF(bytes string) java.JString {
	return env.ObjectRef(env.adEnv.NewStringUTF(bytes))
}

func (env *jniEnv) GetStringUTFLength(str java.JString) java.JSize {
	s := env.getString(str)
	if s == nil {
		return 0
	}
	return s.Length()
}

func (env *jniEnv) GetStringUTFChars(str java.JString) []byte {
	s := env.getString(str)
	if s == nil {
		return nil
	}
	gs := s.String()
	return unsafe.Slice(unsafe.StringData(gs), len(gs))
}

func (env *jniEnv) ReleaseStringUTFChars(java.JString, []byte) {
}

func (env *jniEnv) GetArrayLength(array java.JArray) java.JSize {
	arr := env.getArray(array)
	if arr == nil {
		return 0
	}
	return arr.Length()
}

func (env *jniEnv) NewObjectArray(length java.JSize, elementClass java.JClass, initialElement java.JObject) java.JGenericArray[java.JObject] {
	return env.ObjectRef(env.adEnv.NewObjectArray(length, env.getClass(elementClass), env.GetObject(initialElement)))
}

func (env *jniEnv) GetObjectArrayElement(array java.JGenericArray[java.JObject], index java.JSize) java.JObject {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.IObject]); ok {
		return env.ObjectRef(arr.Get(index))
	}
	return nil
}

func (env *jniEnv) SetObjectArrayElement(array java.JGenericArray[java.JObject], index java.JSize, value java.JObject) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.IObject]); ok {
		arr.Set(index, env.GetObject(value))
	}
}

func (env *jniEnv) NewBooleanArray(length java.JSize) java.JGenericArray[java.JBoolean] {
	return env.ObjectRef(env.adEnv.NewBooleanArray(length))
}

func (env *jniEnv) NewByteArray(length java.JSize) java.JGenericArray[java.JByte] {
	return env.ObjectRef(env.adEnv.NewByteArray(length))
}

func (env *jniEnv) NewCharArray(length java.JSize) java.JGenericArray[java.JChar] {
	return env.ObjectRef(env.adEnv.NewCharArray(length))
}

func (env *jniEnv) NewShortArray(length java.JSize) java.JGenericArray[java.JShort] {
	return env.ObjectRef(env.adEnv.NewShortArray(length))
}

func (env *jniEnv) NewIntArray(length java.JSize) java.JGenericArray[java.JInt] {
	return env.ObjectRef(env.adEnv.NewIntArray(length))
}

func (env *jniEnv) NewLongArray(length java.JSize) java.JGenericArray[java.JLong] {
	return env.ObjectRef(env.adEnv.NewLongArray(length))
}

func (env *jniEnv) NewFloatArray(length java.JSize) java.JGenericArray[java.JFloat] {
	return env.ObjectRef(env.adEnv.NewFloatArray(length))
}

func (env *jniEnv) NewDoubleArray(length java.JSize) java.JGenericArray[java.JDouble] {
	return env.ObjectRef(env.adEnv.NewDoubleArray(length))
}

func (env *jniEnv) GetBooleanArrayElements(array java.JGenericArray[java.JBoolean]) []java.JBoolean {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JBoolean]); ok {
		return arr.Elements()
	}
	return nil
}

func (env *jniEnv) GetByteArrayElements(array java.JGenericArray[java.JByte]) []java.JByte {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JByte]); ok {
		return arr.Elements()
	}
	return nil
}

func (env *jniEnv) GetCharArrayElements(array java.JGenericArray[java.JChar]) []java.JChar {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JChar]); ok {
		return arr.Elements()
	}
	return nil
}

func (env *jniEnv) GetShortArrayElements(array java.JGenericArray[java.JShort]) []java.JShort {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JShort]); ok {
		return arr.Elements()
	}
	return nil
}

func (env *jniEnv) GetIntArrayElements(array java.JGenericArray[java.JInt]) []java.JInt {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JInt]); ok {
		return arr.Elements()
	}
	return nil
}

func (env *jniEnv) GetLongArrayElements(array java.JGenericArray[java.JLong]) []java.JLong {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JLong]); ok {
		return arr.Elements()
	}
	return nil
}

func (env *jniEnv) GetFloatArrayElements(array java.JGenericArray[java.JFloat]) []java.JFloat {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JFloat]); ok {
		return arr.Elements()
	}
	return nil
}

func (env *jniEnv) GetDoubleArrayElements(array java.JGenericArray[java.JDouble]) []java.JDouble {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JDouble]); ok {
		return arr.Elements()
	}
	return nil
}

func (env *jniEnv) ReleaseBooleanArrayElements(array java.JGenericArray[java.JBoolean], elems []java.JBoolean, mode java.JInt) {
	if mode == 0 || mode == java.JNI_COMMIT {
		if arr, ok := env.getArray(array).(java.IGenericArray[java.JBoolean]); ok {
			copy(arr.Elements(), elems)
		}
	}
}

func (env *jniEnv) ReleaseByteArrayElements(array java.JGenericArray[java.JByte], elems []java.JByte, mode java.JInt) {
	if mode == 0 || mode == java.JNI_COMMIT {
		if arr, ok := env.getArray(array).(java.IGenericArray[java.JByte]); ok {
			copy(arr.Elements(), elems)
		}
	}
}

func (env *jniEnv) ReleaseCharArrayElements(array java.JGenericArray[java.JChar], elems []java.JChar, mode java.JInt) {
	if mode == 0 || mode == java.JNI_COMMIT {
		if arr, ok := env.getArray(array).(java.IGenericArray[java.JChar]); ok {
			copy(arr.Elements(), elems)
		}
	}
}

func (env *jniEnv) ReleaseShortArrayElements(array java.JGenericArray[java.JShort], elems []java.JShort, mode java.JInt) {
	if mode == 0 || mode == java.JNI_COMMIT {
		if arr, ok := env.getArray(array).(java.IGenericArray[java.JShort]); ok {
			copy(arr.Elements(), elems)
		}
	}
}

func (env *jniEnv) ReleaseIntArrayElements(array java.JGenericArray[java.JInt], elems []java.JInt, mode java.JInt) {
	if mode == 0 || mode == java.JNI_COMMIT {
		if arr, ok := env.getArray(array).(java.IGenericArray[java.JInt]); ok {
			copy(arr.Elements(), elems)
		}
	}
}

func (env *jniEnv) ReleaseLongArrayElements(array java.JGenericArray[java.JLong], elems []java.JLong, mode java.JInt) {
	if mode == 0 || mode == java.JNI_COMMIT {
		if arr, ok := env.getArray(array).(java.IGenericArray[java.JLong]); ok {
			copy(arr.Elements(), elems)
		}
	}
}

func (env *jniEnv) ReleaseFloatArrayElements(array java.JGenericArray[java.JFloat], elems []java.JFloat, mode java.JInt) {
	if mode == 0 || mode == java.JNI_COMMIT {
		if arr, ok := env.getArray(array).(java.IGenericArray[java.JFloat]); ok {
			copy(arr.Elements(), elems)
		}
	}
}

func (env *jniEnv) ReleaseDoubleArrayElements(array java.JGenericArray[java.JDouble], elems []java.JDouble, mode java.JInt) {
	if mode == 0 || mode == java.JNI_COMMIT {
		if arr, ok := env.getArray(array).(java.IGenericArray[java.JDouble]); ok {
			copy(arr.Elements(), elems)
		}
	}
}

func (env *jniEnv) GetBooleanArrayRegion(array java.JGenericArray[java.JBoolean], start java.JSize, buf []java.JBoolean) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JBoolean]); ok {
		copy(buf, arr.Elements()[start:])
	}
}

func (env *jniEnv) GetByteArrayRegion(array java.JGenericArray[java.JByte], start java.JSize, buf []java.JByte) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JByte]); ok {
		copy(buf, arr.Elements()[start:])
	}
}

func (env *jniEnv) GetCharArrayRegion(array java.JGenericArray[java.JChar], start java.JSize, buf []java.JChar) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JChar]); ok {
		copy(buf, arr.Elements()[start:])
	}
}

func (env *jniEnv) GetShortArrayRegion(array java.JGenericArray[java.JShort], start java.JSize, buf []java.JShort) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JShort]); ok {
		copy(buf, arr.Elements()[start:])
	}
}

func (env *jniEnv) GetIntArrayRegion(array java.JGenericArray[java.JInt], start java.JSize, buf []java.JInt) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JInt]); ok {
		copy(buf, arr.Elements()[start:])
	}
}

func (env *jniEnv) GetLongArrayRegion(array java.JGenericArray[java.JLong], start java.JSize, buf []java.JLong) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JLong]); ok {
		copy(buf, arr.Elements()[start:])
	}
}

func (env *jniEnv) GetFloatArrayRegion(array java.JGenericArray[java.JFloat], start java.JSize, buf []java.JFloat) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JFloat]); ok {
		copy(buf, arr.Elements()[start:])
	}
}

func (env *jniEnv) GetDoubleArrayRegion(array java.JGenericArray[java.JDouble], start java.JSize, buf []java.JDouble) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JDouble]); ok {
		copy(buf, arr.Elements()[start:])
	}
}

func (env *jniEnv) SetBooleanArrayRegion(array java.JGenericArray[java.JBoolean], start java.JSize, buf []java.JBoolean) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JBoolean]); ok {
		copy(arr.Elements()[start:], buf)
	}
}

func (env *jniEnv) SetByteArrayRegion(array java.JGenericArray[java.JByte], start java.JSize, buf []java.JByte) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JByte]); ok {
		copy(arr.Elements()[start:], buf)
	}
}

func (env *jniEnv) SetCharArrayRegion(array java.JGenericArray[java.JChar], start java.JSize, buf []java.JChar) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JChar]); ok {
		copy(arr.Elements()[start:], buf)
	}
}

func (env *jniEnv) SetShortArrayRegion(array java.JGenericArray[java.JShort], start java.JSize, buf []java.JShort) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JShort]); ok {
		copy(arr.Elements()[start:], buf)
	}
}

func (env *jniEnv) SetIntArrayRegion(array java.JGenericArray[java.JInt], start java.JSize, buf []java.JInt) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JInt]); ok {
		copy(arr.Elements()[start:], buf)
	}
}

func (env *jniEnv) SetLongArrayRegion(array java.JGenericArray[java.JLong], start java.JSize, buf []java.JLong) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JLong]); ok {
		copy(arr.Elements()[start:], buf)
	}
}

func (env *jniEnv) SetFloatArrayRegion(array java.JGenericArray[java.JFloat], start java.JSize, buf []java.JFloat) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JFloat]); ok {
		copy(arr.Elements()[start:], buf)
	}
}

func (env *jniEnv) SetDoubleArrayRegion(array java.JGenericArray[java.JDouble], start java.JSize, buf []java.JDouble) {
	if arr, ok := env.getArray(array).(java.IGenericArray[java.JDouble]); ok {
		copy(arr.Elements()[start:], buf)
	}
}

func (env *jniEnv) RegisterNatives(clazz java.JClass, methods []java.JNINativeMethod) java.JInt {
	return env.adEnv.RegisterNatives(env.getClass(clazz), methods)
}

func (env *jniEnv) UnregisterNatives(clazz java.JClass) java.JInt {
	return env.adEnv.UnregisterNatives(env.getClass(clazz))
}

func (env *jniEnv) MonitorEnter(java.JObject) java.JInt {
	return java.JNI_OK
}

func (env *jniEnv) MonitorExit(java.JObject) java.JInt {
	return java.JNI_OK
}

func (env *jniEnv) GetJavaVM(vm *java.JavaVM) java.JInt {
	*vm = env.adEnv.JavaVM()
	return java.JNI_OK
}

func (env *jniEnv) GetStringRegion(str java.JString, start java.JSize, buf []java.JChar) {
	if s := env.getString(str); s != nil {
		copy(buf, utf16.Encode([]rune(s.String()))[start:])
	}
}

func (env *jniEnv) GetStringUTFRegion(str java.JString, start java.JSize, buf []byte) {
	if s := env.getString(str); s != nil {
		ss := s.String()
		copy(buf, unsafe.Slice(unsafe.StringData(ss), len(ss))[start:])
	}
}

func (env *jniEnv) GetPrimitiveArrayCritical(array java.JArray) []byte {
	switch arr := env.getArray(array).(type) {
	case java.IGenericArray[java.JBoolean]:
		elems := arr.Elements()
		return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems))
	case java.IGenericArray[java.JByte]:
		elems := arr.Elements()
		return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems))
	case java.IGenericArray[java.JChar]:
		elems := arr.Elements()
		return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*2)
	case java.IGenericArray[java.JShort]:
		elems := arr.Elements()
		return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*2)
	case java.IGenericArray[java.JInt]:
		elems := arr.Elements()
		return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*4)
	case java.IGenericArray[java.JLong]:
		elems := arr.Elements()
		return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*8)
	case java.IGenericArray[java.JFloat]:
		elems := arr.Elements()
		return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*4)
	case java.IGenericArray[java.JDouble]:
		elems := arr.Elements()
		return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*8)
	}
	return nil
}

func (env *jniEnv) ReleasePrimitiveArrayCritical(array java.JArray, raw []byte, mode java.JInt) {
	if mode == 0 || mode == java.JNI_COMMIT {
		switch arr := env.getArray(array).(type) {
		case java.IGenericArray[java.JBoolean]:
			elems := arr.Elements()
			copy(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)), raw)
		case java.IGenericArray[java.JByte]:
			elems := arr.Elements()
			copy(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)), raw)
		case java.IGenericArray[java.JChar]:
			elems := arr.Elements()
			copy(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*2), raw)
		case java.IGenericArray[java.JShort]:
			elems := arr.Elements()
			copy(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*2), raw)
		case java.IGenericArray[java.JInt]:
			elems := arr.Elements()
			copy(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*4), raw)
		case java.IGenericArray[java.JLong]:
			elems := arr.Elements()
			copy(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*8), raw)
		case java.IGenericArray[java.JFloat]:
			elems := arr.Elements()
			copy(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*4), raw)
		case java.IGenericArray[java.JDouble]:
			elems := arr.Elements()
			copy(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(elems))), len(elems)*8), raw)
		}
	}
}

func (env *jniEnv) GetStringCritical(str java.JString) []java.JChar {
	return env.GetStringChars(str)
}

func (env *jniEnv) ReleaseStringCritical(java.JString, []java.JChar) {
}

func (env *jniEnv) NewWeakGlobalRef(obj java.JObject) java.JWeak {
	if typ := env.GetObjectRefType(obj); typ == java.JNILocalRefType || typ == java.JNIGlobalRefType {
		ref := obj.(gava.Ref) | gava.Ref(java.JNIWeakGlobalRefType)
		env.weakRef.Store(ref, obj)
		return ref
	}
	return nil
}

func (env *jniEnv) DeleteWeakGlobalRef(obj java.JWeak) {
	env.weakRef.Delete(obj)
}

func (env *jniEnv) ExceptionCheck() java.JBoolean {
	return env.adEnv.ExceptionCheck()
}

func (env *jniEnv) NewDirectByteBuffer(java.AnyPtr, java.JLong) java.JObject {
	panic("[NewDirectByteBuffer] Not implemented")
}

func (env *jniEnv) GetDirectBufferAddress(java.JObject) java.AnyPtr {
	panic("[GetDirectBufferAddress] Not implemented")
}

func (env *jniEnv) GetDirectBufferCapacity(java.JObject) java.JLong {
	panic("[GetDirectBufferCapacity] Not implemented")
}

func (env *jniEnv) GetObjectRefType(obj java.JObject) java.JObjectRefType {
	ref, _ := obj.(gava.Ref)
	return java.JObjectRefType(ref & 3)
}

func (env *jniEnv) ObjectRef(obj java.IObject) gava.Ref {
	if getPtr(obj) == nil {
		return nilRef
	}
	ref := gava.Ref(obj.HashCode()<<2|java.JInt(java.JNILocalRefType)) & 0xFFFFFFFF
	env.localRef.Store(ref, obj)
	return ref
}

func (env *jniEnv) GetObject(ref java.JObject) java.IObject {
	switch env.GetObjectRefType(ref) {
	case java.JNILocalRefType:
		if val, ok := env.localRef.Load(ref); ok {
			return val.(java.IObject)
		}
	case java.JNIGlobalRefType:
		if val, ok := env.globalRef.Load(ref); ok {
			return val.(java.IObject)
		}
	case java.JNIWeakGlobalRefType:
		if val, ok := env.weakRef.Load(ref); ok {
			return env.GetObject(val)
		}
	}
	return nil
}

func (env *jniEnv) getClass(ref java.JClass) java.IClass {
	cls, _ := env.GetObject(ref).(java.IClass)
	return cls
}

func (env *jniEnv) getString(ref java.JString) java.IString {
	str, _ := env.GetObject(ref).(java.IString)
	return str
}

func (env *jniEnv) getArray(ref java.JArray) java.IArray {
	arr, _ := env.GetObject(ref).(java.IArray)
	return arr
}

func (env *jniEnv) getMethod(ref java.JMethodID) java.IMethod {
	if val, ok := env.methods.Load(ref); ok {
		return val.(java.IMethod)
	}
	return nil
}

func (env *jniEnv) getField(ref java.JFieldID) java.IField {
	if val, ok := env.fields.Load(ref); ok {
		return val.(java.IField)
	}
	return nil
}

func (env *jniEnv) methodRef(method java.IMethod) gava.Ref {
	if getPtr(method) == nil {
		return nilRef
	}
	ref := gava.Ref(method.HashCode()<<1) & 0xFFFFFFFF
	env.methods.Store(ref, method)
	return ref
}

func (env *jniEnv) fieldRef(field java.IField) gava.Ref {
	if getPtr(field) == nil {
		return nilRef
	}
	ref := gava.Ref(field.HashCode()<<1) & 0xFFFFFFFF
	env.fields.Store(ref, field)
	return ref
}

func (env *jniEnv) deRefArgs(args []any) []any {
	for i := range args {
		if ref, ok := args[i].(gava.Ref); ok {
			args[i] = env.GetObject(ref)
		}
	}
	return args
}

func (env *jniEnv) extractVaList(method java.IMethod, args java.VaList) ([]any, error) {
	arr := make([]any, method.GetParameterCount())
	for i, typ := range method.GetParameterTypes() {
		var err error
		switch typ.DescriptorString().String() {
		case "Z":
			var v java.JBoolean
			err = args.Extract(&v)
			arr[i] = v
		case "C":
			var v java.JChar
			err = args.Extract(&v)
			arr[i] = v
		case "B":
			var v java.JByte
			err = args.Extract(&v)
			arr[i] = v
		case "S":
			var v java.JShort
			err = args.Extract(&v)
			arr[i] = v
		case "I":
			var v java.JInt
			err = args.Extract(&v)
			arr[i] = v
		case "J":
			var v java.JLong
			err = args.Extract(&v)
			arr[i] = v
		case "F":
			var v java.JFloat
			err = args.Extract(&v)
			arr[i] = v
		case "D":
			var v java.JDouble
			err = args.Extract(&v)
			arr[i] = v
		default:
			var v gava.Ref
			err = args.Extract(&v)
			arr[i] = env.GetObject(v)
		}
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func (env *jniEnv) extractPtr(method java.IMethod, args java.TypePtr[java.JValue]) ([]any, error) {
	v := make([]java.JValue, method.GetParameterCount())
	if _, err := args.ReadAt(v, 0); err != nil {
		return nil, err
	}
	arr := make([]any, len(v))
	for i, typ := range method.GetParameterTypes() {
		switch typ.DescriptorString().String() {
		case "Z":
			arr[i] = v[i].JBoolean()
		case "C":
			arr[i] = v[i].JChar()
		case "B":
			arr[i] = v[i].JByte()
		case "S":
			arr[i] = v[i].JShort()
		case "I":
			arr[i] = v[i].JInt()
		case "J":
			arr[i] = v[i].JLong()
		case "F":
			arr[i] = v[i].JFloat()
		case "D":
			arr[i] = v[i].JDouble()
		default:
			arr[i] = env.GetObject(v[i].JObject())
		}
	}
	return arr, nil
}

func getPtr(v any) unsafe.Pointer {
	return (*struct{ _, data unsafe.Pointer })(unsafe.Pointer(&v)).data
}
