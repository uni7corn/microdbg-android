package java

import (
	"fmt"
	"strings"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg/debugger"
)

type FakeField interface {
	java.IField
	BindGet(getter func(obj java.IObject) any) FakeField
	BindSet(setter func(obj java.IObject, value any)) FakeField
}

type fakeField struct {
	cls  FakeClass
	name string
	typ  java.IClass
	mod  Modifier
	get  func(java.IObject) any
	set  func(java.IObject, any)
}

func (field *fakeField) GetClass() java.IClass {
	return FakeFieldClass
}

func (field *fakeField) HashCode() java.JInt {
	return java.JInt(uintptr(unsafe.Pointer(field)))
}

func (field *fakeField) Equals(obj java.IObject) java.JBoolean {
	if other, ok := obj.(*fakeField); ok {
		return field.cls.Equals(other.cls) && field.name == other.name && field.typ.Equals(other.typ)
	}
	return false
}

func (field *fakeField) ToString() java.IString {
	var sb strings.Builder
	if field.mod != 0 {
		sb.WriteString(field.mod.String())
		sb.WriteByte(' ')
	}
	sb.WriteString(field.typ.GetTypeName().String())
	sb.WriteByte(' ')
	sb.WriteString(field.cls.GetTypeName().String())
	sb.WriteByte('.')
	sb.WriteString(field.name)
	return FakeString(sb.String())
}

func (field *fakeField) GetName() java.IString {
	return FakeString(field.name)
}

func (field *fakeField) GetModifiers() java.JInt {
	return java.JInt(field.mod)
}

func (field *fakeField) GetType() java.IClass {
	return field.typ
}

func (field *fakeField) Get(obj java.IObject) java.IObject {
	return ToObject[java.IObject](field.GetPrimitive(obj))
}

func (field *fakeField) GetPrimitive(obj java.IObject) any {
	if field.get != nil {
		return field.get(obj)
	} else if fake, ok := obj.(FakeProperty); ok {
		if val, ok := fake.Get(field.name); ok {
			return val
		}
	} else if IsStatic(field) {
		if val, ok := field.cls.Get(field.name); ok {
			return val
		}
	}
	panic(fmt.Errorf("%s %w", field.ToString(), debugger.ErrNotImplemented))
}

func (field *fakeField) Set(obj java.IObject, value java.IObject) {
	field.SetPrimitive(obj, value)
}

func (field *fakeField) SetPrimitive(obj java.IObject, value any) {
	if field.set != nil {
		field.set(obj, value)
	} else if fake, ok := obj.(FakeProperty); ok {
		fake.Set(field.name, value)
	} else if IsStatic(field) {
		field.cls.Set(field.name, value)
	}
}

func (field *fakeField) BindGet(get func(java.IObject) any) FakeField {
	field.get = get
	return field
}

func (field *fakeField) BindSet(set func(java.IObject, any)) FakeField {
	field.set = set
	return field
}

func (field *fakeField) parseDescriptor(cf ClassFactory, sig string) {
	field.typ = cf.GetClass(sig)
}
