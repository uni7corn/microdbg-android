package java

import (
	"fmt"
	"strings"
	"sync"

	java "github.com/wnxd/microdbg-java"
)

type FakeClass interface {
	java.IClass
	GetModifiers() java.JInt
	SetModifiers(mod Modifier)
	FakeProperty
	NewObject(v any) FakeObject
	NewThrowable(msg string) java.IThrowable
	NewArray(length int) java.IArray
	FindMethod(name, sig string) FakeMethod
	GetMethod(name, sig string) FakeMethod
	GetStaticMethod(name, sig string) FakeMethod
	DefineMethod(name, sig string, mod Modifier) FakeMethod
	ClearNativeMethods()
	FindField(name, sig string) FakeField
	GetField(name, sig string) FakeField
	GetStaticField(name, sig string) FakeField
	DefineField(name, sig string, mod Modifier) FakeField
}

type wrapClass struct {
	fakeClass
	cls java.IClass
}

type fakeClass struct {
	cf    ClassFactory
	name  string
	super java.IClass
	iface []java.IClass
	mod   Modifier
	fakeProperty
	methods sync.Map
	fields  sync.Map
}

type fakeArrayClass struct {
	fakeClass
	elem java.IClass
}

func (cls *wrapClass) IsInterface() java.JBoolean {
	return cls.cls.IsInterface()
}

func (cls *wrapClass) ComponentType() java.IClass {
	if arr, ok := cls.cls.(interface{ ComponentType() java.IClass }); ok {
		return arr.ComponentType()
	}
	return nil
}

func (cls *fakeClass) GetClass() java.IClass {
	return FakeClassClass
}

func (cls *fakeClass) HashCode() java.JInt {
	return HashCode(cls.name)
}

func (cls *fakeClass) Equals(obj java.IObject) java.JBoolean {
	if other, ok := obj.(interface {
		IsArray() java.JBoolean
		ComponentType() java.IClass
	}); ok && cls.IsArray() && other.IsArray() {
		return cls.ComponentType().Equals(other.ComponentType())
	} else if other, ok := obj.(*fakeClass); ok {
		return cls.name == other.name
	} else if other, ok := obj.(java.IClass); ok {
		return cls.name == other.GetName().String()
	}
	return false
}

func (cls *fakeClass) ToString() java.IString {
	var prefix string
	if cls.IsInterface() {
		prefix = "interface "
	} else if cls.IsPrimitive() {
		prefix = ""
	} else {
		prefix = "class "
	}
	prefix += cls.name
	return FakeString(prefix)
}

func (cls *fakeClass) NewInstance() java.IObject {
	return cls.NewObject(nil)
}

func (cls *fakeClass) GetName() java.IString {
	return FakeString(cls.name)
}

func (cls *fakeClass) GetSimpleName() java.IString {
	return FakeString(cls.name[strings.LastIndexByte(cls.name, '.')+1:])
}

func (cls *fakeClass) GetTypeName() java.IString {
	return cls.GetName()
}

func (cls *fakeClass) DescriptorString() java.IString {
	if cls.IsPrimitive() {
		switch cls.name {
		case "boolean":
			return FakeString("Z")
		case "byte":
			return FakeString("B")
		case "char":
			return FakeString("C")
		case "short":
			return FakeString("S")
		case "int":
			return FakeString("I")
		case "long":
			return FakeString("J")
		case "float":
			return FakeString("F")
		case "double":
			return FakeString("D")
		}
	}
	return FakeString("L" + strings.ReplaceAll(cls.name, ".", "/") + ";")
}

func (cls *fakeClass) GetSuperclass() java.IClass {
	if cls.super != nil {
		return cls.super
	} else if cls.IsInterface() {
		return nil
	}
	return FakeObjectClass
}

func (cls *fakeClass) GetInterfaces() []java.IClass {
	return cls.iface
}

func (cls *fakeClass) IsInterface() java.JBoolean {
	return cls.mod&Modifier_INTERFACE != 0
}

func (cls *fakeClass) IsAssignableFrom(clazz java.IClass) java.JBoolean {
	for ; clazz != nil; clazz = clazz.GetSuperclass() {
		if cls.IsInterface() {
			for _, iface := range clazz.GetInterfaces() {
				if cls.Equals(iface) {
					return true
				}
			}
		} else if cls.Equals(clazz) {
			return true
		}
	}
	return false
}

func (cls *fakeClass) IsPrimitive() java.JBoolean {
	return strings.IndexByte(cls.name, '.') == -1
}

func (cls *fakeClass) IsArray() java.JBoolean {
	return strings.HasPrefix(cls.name, "[")
}

func (cls *fakeClass) IsInstance(obj java.IObject) java.JBoolean {
	if obj == nil {
		return false
	}
	return cls.IsAssignableFrom(obj.GetClass())
}

func (cls *fakeClass) Cast(obj java.IObject) java.IObject {
	if fake, ok := obj.(*fakeObject); ok {
		return &fakeObject{cls: cls, val: fake.val}
	}
	return nil
}

func (cls *fakeClass) ComponentType() java.IClass {
	return nil
}

func (cls *fakeClass) GetModifiers() java.JInt {
	return java.JInt(cls.mod)
}

func (cls *fakeClass) SetModifiers(mod Modifier) {
	cls.mod = mod
}

func (cls *fakeClass) NewObject(val any) FakeObject {
	return &fakeObject{cls: cls, val: val}
}

func (cls *fakeClass) NewThrowable(msg string) java.IThrowable {
	obj := &fakeObject{cls: cls}
	obj.Set("detailMessage", FakeString(msg))
	return obj
}

func (cls *fakeClass) NewArray(length int) java.IArray {
	switch cls.name {
	case "boolean":
		return make(FakeZArray, length)
	case "byte":
		return make(FakeBArray, length)
	case "char":
		return make(FakeCArray, length)
	case "short":
		return make(FakeSArray, length)
	case "int":
		return make(FakeIArray, length)
	case "long":
		return make(FakeJArray, length)
	case "float":
		return make(FakeFArray, length)
	case "double":
		return make(FakeDArray, length)
	case "void":
		return nil
	default:
		return fakeObjectArray{fakeArray: make(fakeArray[java.IObject], length), cls: arrayOf(cls.cf, cls)}
	}
}

func (cls *fakeClass) FindMethod(name, sig string) FakeMethod {
	h := HashCode(name) ^ HashCode(sig)
	if method, ok := cls.methods.Load(h); ok {
		return method.(FakeMethod)
	}
	if super, ok := cls.super.(FakeClass); ok {
		return super.FindMethod(name, sig)
	}
	return nil
}

func (cls *fakeClass) GetMethod(name, sig string) FakeMethod {
	if method := cls.FindMethod(name, sig); method != nil {
		return method
	}
	return cls.DefineMethod(name, sig, Modifier_PUBLIC)
}

func (cls *fakeClass) GetStaticMethod(name, sig string) FakeMethod {
	if method := cls.FindMethod(name, sig); method != nil {
		return method
	}
	return cls.DefineMethod(name, sig, Modifier_PUBLIC|Modifier_STATIC)
}

func (cls *fakeClass) DefineMethod(name, sig string, mod Modifier) FakeMethod {
	if cls.cf == nil {
		panic(fmt.Errorf("[DefineMethod] %s: %s %s not allowed", cls.name, name, sig))
	}
	h := HashCode(name) ^ HashCode(sig)
	method := cls.cf.DefineMethod(cls, name, sig, mod)
	cls.methods.Store(h, method)
	return method
}

func (cls *fakeClass) ClearNativeMethods() {
	for h, method := range cls.methods.Range {
		if IsNative(method.(FakeMethod)) {
			cls.methods.Delete(h)
		}
	}
}

func (cls *fakeClass) FindField(name, sig string) FakeField {
	h := HashCode(name) ^ HashCode(sig)
	if field, ok := cls.fields.Load(h); ok {
		return field.(FakeField)
	}
	if super, ok := cls.super.(FakeClass); ok {
		return super.FindField(name, sig)
	}
	return nil
}

func (cls *fakeClass) GetField(name, sig string) FakeField {
	if field := cls.FindField(name, sig); field != nil {
		return field
	}
	return cls.DefineField(name, sig, Modifier_PUBLIC)
}

func (cls *fakeClass) GetStaticField(name, sig string) FakeField {
	if field := cls.FindField(name, sig); field != nil {
		return field
	}
	return cls.DefineField(name, sig, Modifier_PUBLIC|Modifier_STATIC)
}

func (cls *fakeClass) DefineField(name, sig string, mod Modifier) FakeField {
	if cls.cf == nil {
		panic(fmt.Errorf("[DefineField] %s: %s %s not allowed", cls.name, name, sig))
	}
	h := HashCode(name) ^ HashCode(sig)
	field := cls.cf.DefineField(cls, name, sig, mod)
	cls.fields.Store(h, field)
	return field
}

func (cls *fakeArrayClass) GetSimpleName() java.IString {
	return FakeString(cls.elem.GetSimpleName().String() + "[]")
}

func (cls *fakeArrayClass) GetTypeName() java.IString {
	return FakeString(cls.elem.GetTypeName().String() + "[]")
}

func (cls *fakeArrayClass) DescriptorString() java.IString {
	return FakeString("[" + cls.elem.DescriptorString().String())
}

func (cls *fakeArrayClass) ComponentType() java.IClass {
	return cls.elem
}

func ObjectOf(cls FakeClass, v any) java.IObject {
	if cls.IsArray() {
		return ArrayOf(cls, v)
	}
	return &fakeObject{
		cls: cls,
		val: v,
	}
}

func init() {
	definePrimitiveMethod(FakeClassClass, "getName", "()Ljava/lang/String;", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return obj.(java.IClass).GetName()
	})
	definePrimitiveMethod(FakeClassClass, "getSimpleName", "()Ljava/lang/String;", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return obj.(java.IClass).GetSimpleName()
	})
	definePrimitiveMethod(FakeClassClass, "getTypeName", "()Ljava/lang/String;", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return obj.(java.IClass).GetTypeName()
	})
	definePrimitiveMethod(FakeClassClass, "descriptorString", "()Ljava/lang/String;", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return obj.(java.IClass).DescriptorString()
	})
	definePrimitiveMethod(FakeClassClass, "getSuperclass", "()Ljava/lang/Class;", Modifier_PUBLIC|Modifier_NATIVE).BindCall(func(obj java.IObject, _ ...any) any {
		return obj.(java.IClass).GetSuperclass()
	})
	definePrimitiveMethod(FakeClassClass, "getInterfaces", "()[Ljava/lang/Class;", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return ArrayOf(FakeClassArrayClass, obj.(java.IClass).GetInterfaces())
	})
}
