package java

import (
	"fmt"
	"strings"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg/debugger"
)

const ConstructorMethodName = "<init>"

type FakeMethod interface {
	java.IMethod
	GetReturnType() java.IClass
	Descriptor() string
	IsConstructor() bool
	BindCall(f func(obj java.IObject, args ...any) any) FakeMethod
}

type fakeMethod struct {
	cls    java.IClass
	name   string
	sig    string
	params []java.IClass
	ret    java.IClass
	mod    Modifier
	f      func(java.IObject, ...any) any
}

func (method *fakeMethod) GetClass() java.IClass {
	if method.IsConstructor() {
		return FakeConstructorClass
	}
	return FakeMethodClass
}

func (method *fakeMethod) HashCode() java.JInt {
	return java.JInt(uintptr(unsafe.Pointer(method)))
}

func (method *fakeMethod) Equals(obj java.IObject) java.JBoolean {
	if other, ok := obj.(*fakeMethod); ok {
		return method.cls.Equals(other.cls) && method.name == other.name
	}
	return false
}

func (method *fakeMethod) ToString() java.IString {
	var sb strings.Builder
	if method.mod != 0 {
		sb.WriteString(method.mod.String())
		sb.WriteByte(' ')
	}
	if method.IsConstructor() {
		sb.WriteString(method.cls.GetTypeName().String())
	} else {
		sb.WriteString(method.ret.GetTypeName().String())
		sb.WriteByte(' ')
		sb.WriteString(method.cls.GetTypeName().String())
		sb.WriteByte('.')
		sb.WriteString(method.name)
	}
	sb.WriteByte('(')
	for i, v := range method.params {
		if i != 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(v.GetTypeName().String())
	}
	sb.WriteByte(')')
	return FakeString(sb.String())
}

func (method *fakeMethod) GetName() java.IString {
	if method.IsConstructor() {
		return method.cls.GetName()
	}
	return FakeString(method.name)
}

func (method *fakeMethod) GetModifiers() java.JInt {
	return java.JInt(method.mod)
}

func (method *fakeMethod) GetParameterTypes() []java.IClass {
	return method.params
}

func (method *fakeMethod) GetParameterCount() java.JInt {
	return java.JInt(len(method.params))
}

func (method *fakeMethod) Call(obj java.IObject, args ...any) java.IObject {
	return ToObject[java.IObject](method.CallPrimitive(obj, args...))
}

func (method *fakeMethod) CallPrimitive(obj java.IObject, args ...any) any {
	if method.f == nil {
		panic(fmt.Errorf("%s %w", method.ToString(), debugger.ErrNotImplemented))
	}
	return method.f(obj, args...)
}

func (method *fakeMethod) GetReturnType() java.IClass {
	if method.IsConstructor() {
		return method.cls
	}
	return method.ret
}

func (method *fakeMethod) Descriptor() string {
	return method.sig
}

func (method *fakeMethod) IsConstructor() bool {
	return method.name == ConstructorMethodName
}

func (method *fakeMethod) BindCall(f func(java.IObject, ...any) any) FakeMethod {
	method.f = f
	return method
}

func (method *fakeMethod) parseDescriptor(cf ClassFactory, sig string) {
	method.sig = sig
	method.ret = FakeVoidTYPE
	var dim int
	for i := 1; i < len(sig); i++ {
		var cls FakeClass
		switch sig[i] {
		case 'Z':
			cls = FakeBooleanTYPE
		case 'B':
			cls = FakeByteTYPE
		case 'C':
			cls = FakeCharTYPE
		case 'S':
			cls = FakeShortTYPE
		case 'I':
			cls = FakeIntTYPE
		case 'J':
			cls = FakeLongTYPE
		case 'F':
			cls = FakeFloatTYPE
		case 'D':
			cls = FakeDoubleTYPE
		case 'V':
			cls = FakeVoidTYPE
		case 'L':
			n := strings.IndexByte(sig[i:], ';') + i
			cls = cf.GetClass(sig[i+1 : n])
			i = n
		case '[':
			dim++
			continue
		case ')':
			method.ret = cf.GetClass(sig[i+1:])
			return
		}
		for n := 0; n < dim; n++ {
			cls = cf.ArrayOf(cls)
		}
		dim = 0
		method.params = append(method.params, cls)
	}
}

func GetMethodDescriptor(method java.IMethod) (string, string) {
	if i, ok := method.(interface{ Descriptor() string }); ok {
		return method.GetName().String(), i.Descriptor()
	}
	var sb strings.Builder
	sb.WriteByte('(')
	for _, typ := range method.GetParameterTypes() {
		sb.WriteString(typ.DescriptorString().String())
	}
	sb.WriteByte(')')
	if m, ok := method.(interface{ GetReturnType() java.IClass }); ok {
		sb.WriteString(m.GetReturnType().DescriptorString().String())
	} else {
		sb.WriteByte('V')
	}
	return method.GetName().String(), sb.String()
}
