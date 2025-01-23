package java

import (
	"fmt"
	"strconv"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
)

type FakeObject interface {
	java.IObject
	Value() any
	FakeProperty
	CallMethod(method java.IMethod, args ...any) any
}

type fakeObject struct {
	cls FakeClass
	val any
	fakeProperty
}

func (obj *fakeObject) GetClass() java.IClass {
	return obj.cls
}

func (obj *fakeObject) HashCode() java.JInt {
	return java.JInt(uintptr(unsafe.Pointer(obj)))
}

func (obj *fakeObject) Equals(other java.IObject) java.JBoolean {
	return obj == other
}

func (obj *fakeObject) ToString() java.IString {
	return FakeString(obj.GetClass().GetName().String() + "@" + strconv.FormatInt(int64(obj.HashCode()), 16))
}

func (obj *fakeObject) GetMessage() java.IString {
	if msg, ok := obj.Get("detailMessage"); ok {
		return msg.(java.IString)
	}
	return nil
}

func (obj *fakeObject) Value() any {
	return obj.val
}

func (obj *fakeObject) CallMethod(method java.IMethod, args ...any) any {
	fake := obj.cls.FindMethod(method.GetName().String(), GetMethodDescriptor(method))
	if fake == nil {
		return method.CallPrimitive(obj, args...)
	}
	return fake.CallPrimitive(obj, args...)
}

func ToObject[O java.IObject](v any) O {
	r, _ := v.(O)
	return r
}

func init() {
	definePrimitiveMethod(FakeObjectClass, "getClass", "()Ljava/lang/Class;", Modifier_PUBLIC|Modifier_FINAL|Modifier_NATIVE).BindCall(func(obj java.IObject, _ ...any) any {
		return obj.GetClass()
	})
	definePrimitiveMethod(FakeObjectClass, "hashCode", "()I", Modifier_PUBLIC|Modifier_NATIVE).BindCall(func(obj java.IObject, _ ...any) any {
		return obj.HashCode()
	})
	definePrimitiveMethod(FakeObjectClass, "equals", "(Ljava/lang/Object;)Z", Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		return obj.Equals(ToObject[java.IObject](args[0]))
	})
	definePrimitiveMethod(FakeObjectClass, "clone", "()Ljava/lang/Object;", Modifier_PROTECTED|Modifier_NATIVE).BindCall(func(obj java.IObject, args ...any) any {
		panic(fmt.Errorf("[Object.clone] %s not implemented", obj.GetClass().GetName()))
	})
	definePrimitiveMethod(FakeObjectClass, "toString", "()Ljava/lang/String;", Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		return obj.ToString()
	})
}
