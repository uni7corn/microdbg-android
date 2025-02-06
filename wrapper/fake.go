package wrapper

import (
	"fmt"
	"sync"
	"unicode/utf16"

	android "github.com/wnxd/microdbg-android"
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg/debugger"
)

type FakeDefineHandler interface {
	DefineClass(FakeContext, string) gava.FakeClass
	DefineMethod(FakeContext, gava.FakeClass, string, string) gava.FakeMethod
	DefineStaticMethod(FakeContext, gava.FakeClass, string, string) gava.FakeMethod
	DefineField(FakeContext, gava.FakeClass, string, string) gava.FakeField
	DefineStaticField(FakeContext, gava.FakeClass, string, string) gava.FakeField
}

type FakeClassHandler interface {
	CallMethod(android.JNIContext, gava.FakeObject, string, string, ...any) any
	CallStaticMethod(android.JNIContext, gava.FakeClass, string, string, ...any) any
	GetField(android.JNIContext, gava.FakeObject, string) any
	SetField(android.JNIContext, gava.FakeObject, string, any)
	GetStaticField(android.JNIContext, gava.FakeClass, string) any
	SetStaticField(android.JNIContext, gava.FakeClass, string, any)
}

type FakeContext interface {
	gava.ClassFactory
	Bind(gava.FakeClass, FakeClassHandler)
	BindClass(string, FakeClassHandler)
}

type FakeWrapper interface {
	android.JNIEnv
	ClassFactory() gava.ClassFactory
	Bind(gava.FakeClass, FakeClassHandler)
	BindClass(string, FakeClassHandler)
}

type fakeWrapper struct {
	handler FakeDefineHandler
	cf      gava.ClassFactory
	binds   sync.Map
}

type fakeContext struct {
	w *fakeWrapper
}

type fakeJNIContext struct {
	android.JNIContext
	cf gava.ClassFactory
}

func NewFake(handler FakeDefineHandler) FakeWrapper {
	w := &fakeWrapper{handler: handler}
	w.cf = gava.NewClassFactory(w.preDefineClass)
	return w
}

func (w *fakeWrapper) DefineClass(android.JNIContext, string, java.IObject, []java.JByte) (java.IClass, error) {
	panic(fmt.Errorf("[FakeWrapper.DefineClass] %w", debugger.ErrNotImplemented))
}

func (w *fakeWrapper) FindClass(ctx android.JNIContext, name string) (java.IClass, error) {
	return w.cf.GetClass(name), nil
}

func (w *fakeWrapper) ThrowNew(ctx android.JNIContext, clazz java.IClass, message string) (java.JInt, error) {
	fake := clazz.(gava.FakeClass)
	ex := fake.NewThrowable(message)
	ctx.Throw(ex)
	return java.JNI_OK, nil
}

func (w *fakeWrapper) GetMethod(ctx android.JNIContext, clazz java.IClass, name, sig string) (java.IMethod, error) {
	fake := clazz.(gava.FakeClass)
	var method gava.FakeMethod
	if w.handler != nil {
		method = w.handler.DefineMethod(fakeContext{w}, fake, name, sig)
	}
	if method == nil {
		method = w.getMethod(fake, name, sig)
	}
	if h, ok := w.getHandler(fake); ok {
		method.BindCall(func(obj java.IObject, args ...any) any {
			return h.CallMethod(fakeJNIContext{ctx, fakeContext{w}}, obj.(gava.FakeObject), name, sig, args...)
		})
	}
	return method, nil
}

func (w *fakeWrapper) GetField(ctx android.JNIContext, clazz java.IClass, name, sig string) (java.IField, error) {
	fake := clazz.(gava.FakeClass)
	var field gava.FakeField
	if w.handler != nil {
		field = w.handler.DefineField(fakeContext{w}, fake, name, sig)
	}
	if field == nil {
		field = w.getField(fake, name, sig)
	}
	if h, ok := w.getHandler(fake); ok {
		field.BindGet(func(obj java.IObject) any {
			return h.GetField(fakeJNIContext{ctx, fakeContext{w}}, obj.(gava.FakeObject), name)
		})
		field.BindSet(func(obj java.IObject, value any) {
			h.SetField(fakeJNIContext{ctx, fakeContext{w}}, obj.(gava.FakeObject), name, value)
		})
	}
	return field, nil
}

func (w *fakeWrapper) GetStaticMethod(ctx android.JNIContext, clazz java.IClass, name, sig string) (java.IMethod, error) {
	fake := clazz.(gava.FakeClass)
	var method gava.FakeMethod
	if w.handler != nil {
		method = w.handler.DefineStaticMethod(fakeContext{w}, fake, name, sig)
	}
	if method == nil {
		method = w.getStaticMethod(fake, name, sig)
	}
	if h, ok := w.getHandler(fake); ok {
		method.BindCall(func(obj java.IObject, args ...any) any {
			return h.CallStaticMethod(fakeJNIContext{ctx, fakeContext{w}}, fake, name, sig, args...)
		})
	}
	return method, nil
}

func (w *fakeWrapper) GetStaticField(ctx android.JNIContext, clazz java.IClass, name, sig string) (java.IField, error) {
	fake := clazz.(gava.FakeClass)
	var field gava.FakeField
	if w.handler != nil {
		field = w.handler.DefineStaticField(fakeContext{w}, fake, name, sig)
	}
	if field == nil {
		field = w.getStaticField(fake, name, sig)
	}
	if h, ok := w.getHandler(fake); ok {
		field.BindGet(func(obj java.IObject) any {
			return h.GetStaticField(fakeJNIContext{ctx, fakeContext{w}}, fake, name)
		})
		field.BindSet(func(obj java.IObject, value any) {
			h.SetStaticField(fakeJNIContext{ctx, fakeContext{w}}, fake, name, value)
		})
	}
	return field, nil
}

func (w *fakeWrapper) NewString(ctx android.JNIContext, chars []java.JChar) (java.IString, error) {
	return gava.FakeString(utf16.Decode(chars)), nil
}

func (w *fakeWrapper) NewStringUTF(ctx android.JNIContext, bytes string) (java.IString, error) {
	return gava.FakeString(bytes), nil
}

func (w *fakeWrapper) NewObjectArray(ctx android.JNIContext, length java.JSize, elementClass java.IClass, initialElement java.IObject) (java.IObjectArray, error) {
	arr := elementClass.(gava.FakeClass).NewArray(int(length)).(java.IObjectArray)
	if initialElement != nil {
		raw := arr.Elements()
		for i := range raw {
			raw[i] = initialElement
		}
	}
	return arr, nil
}

func (w *fakeWrapper) NewBooleanArray(ctx android.JNIContext, length java.JSize) (java.IBooleanArray, error) {
	return gava.FakeBooleanTYPE.NewArray(int(length)).(java.IBooleanArray), nil
}

func (w *fakeWrapper) NewByteArray(ctx android.JNIContext, length java.JSize) (java.IByteArray, error) {
	return gava.FakeByteTYPE.NewArray(int(length)).(java.IByteArray), nil
}

func (w *fakeWrapper) NewCharArray(ctx android.JNIContext, length java.JSize) (java.ICharArray, error) {
	return gava.FakeCharTYPE.NewArray(int(length)).(java.ICharArray), nil
}

func (w *fakeWrapper) NewShortArray(ctx android.JNIContext, length java.JSize) (java.IShortArray, error) {
	return gava.FakeShortTYPE.NewArray(int(length)).(java.IShortArray), nil
}

func (w *fakeWrapper) NewIntArray(ctx android.JNIContext, length java.JSize) (java.IIntArray, error) {
	return gava.FakeIntTYPE.NewArray(int(length)).(java.IIntArray), nil
}

func (w *fakeWrapper) NewLongArray(ctx android.JNIContext, length java.JSize) (java.ILongArray, error) {
	return gava.FakeLongTYPE.NewArray(int(length)).(java.ILongArray), nil
}

func (w *fakeWrapper) NewFloatArray(ctx android.JNIContext, length java.JSize) (java.IFloatArray, error) {
	return gava.FakeFloatTYPE.NewArray(int(length)).(java.IFloatArray), nil
}

func (w *fakeWrapper) NewDoubleArray(ctx android.JNIContext, length java.JSize) (java.IDoubleArray, error) {
	return gava.FakeDoubleTYPE.NewArray(int(length)).(java.IDoubleArray), nil
}

func (w *fakeWrapper) ClassFactory() gava.ClassFactory {
	return w.cf
}

func (w *fakeWrapper) Bind(clazz gava.FakeClass, handler FakeClassHandler) {
	if clazz == nil {
		return
	}
	w.binds.Store(clazz.HashCode(), handler)
}

func (w *fakeWrapper) BindClass(name string, handler FakeClassHandler) {
	w.Bind(w.cf.GetClass(name), handler)
}

func (w *fakeWrapper) getHandler(clazz gava.FakeClass) (FakeClassHandler, bool) {
	if val, ok := w.binds.Load(clazz.HashCode()); ok {
		h, ok := val.(FakeClassHandler)
		return h, ok
	}
	return nil, false
}

func (w *fakeWrapper) getMethod(cls gava.FakeClass, name, sig string) gava.FakeMethod {
	if method := cls.FindMethod(name, sig); method != nil {
		return method
	}
	return w.cf.DefineMethod(cls, name, sig, gava.Modifier_PUBLIC)
}

func (w *fakeWrapper) getStaticMethod(cls gava.FakeClass, name, sig string) gava.FakeMethod {
	if method := cls.FindMethod(name, sig); method != nil {
		return method
	}
	return w.cf.DefineMethod(cls, name, sig, gava.Modifier_PUBLIC|gava.Modifier_STATIC)
}

func (w *fakeWrapper) getField(cls gava.FakeClass, name, sig string) gava.FakeField {
	if field := cls.FindField(name, sig); field != nil {
		return field
	}
	return w.cf.DefineField(cls, name, sig, gava.Modifier_PUBLIC)
}

func (w *fakeWrapper) getStaticField(cls gava.FakeClass, name, sig string) gava.FakeField {
	if field := cls.FindField(name, sig); field != nil {
		return field
	}
	return w.cf.DefineField(cls, name, sig, gava.Modifier_PUBLIC|gava.Modifier_STATIC)
}

func (w *fakeWrapper) preDefineClass(cf gava.ClassFactory, name string) gava.FakeClass {
	if w.handler != nil {
		return w.handler.DefineClass(fakeContext{w}, name)
	}
	return nil
}

func (ctx fakeContext) WrapClass(cls java.IClass) gava.FakeClass {
	return ctx.w.cf.WrapClass(cls)
}

func (ctx fakeContext) FindClass(name string) (gava.FakeClass, bool) {
	return ctx.w.cf.FindClass(name)
}

func (ctx fakeContext) GetClass(name string) gava.FakeClass {
	return ctx.w.cf.GetClass(name)
}

func (ctx fakeContext) DefineClass(name string, extends ...java.IClass) gava.FakeClass {
	return ctx.w.cf.DefineClass(name, extends...)
}

func (ctx fakeContext) ArrayOf(elem java.IClass) gava.FakeClass {
	return ctx.w.cf.ArrayOf(elem)
}

func (ctx fakeContext) DefineMethod(clazz gava.FakeClass, name, sig string, mod gava.Modifier) gava.FakeMethod {
	return ctx.w.cf.DefineMethod(clazz, name, sig, mod)
}

func (ctx fakeContext) DefineField(clazz gava.FakeClass, name, sig string, mod gava.Modifier) gava.FakeField {
	return ctx.w.cf.DefineField(clazz, name, sig, mod)
}

func (ctx fakeContext) Bind(clazz gava.FakeClass, handler FakeClassHandler) {
	ctx.w.Bind(clazz, handler)
}

func (ctx fakeContext) BindClass(name string, handler FakeClassHandler) {
	ctx.w.BindClass(name, handler)
}

func (ctx fakeJNIContext) ClassFactory() gava.ClassFactory {
	return ctx.cf
}
