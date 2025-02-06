package java

import (
	"hash/fnv"
	"strings"
	"sync"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
)

type ClassFactory interface {
	WrapClass(clazz java.IClass) FakeClass
	FindClass(name string) (FakeClass, bool)
	GetClass(name string) FakeClass
	DefineClass(name string, extends ...java.IClass) FakeClass
	ArrayOf(clazz java.IClass) FakeClass
	DefineMethod(clazz FakeClass, name, sig string, mod Modifier) FakeMethod
	DefineField(clazz FakeClass, name, sig string, mod Modifier) FakeField
}

type classFactory struct {
	f       func(ClassFactory, string) FakeClass
	classes sync.Map
}

func NewClassFactory(f func(ClassFactory, string) FakeClass) ClassFactory {
	return &classFactory{f: f}
}

func (cf *classFactory) WrapClass(cls java.IClass) FakeClass {
	if isNil(cls) {
		return nil
	} else if fake, ok := cls.(FakeClass); ok {
		return fake
	}
	return &wrapClass{
		fakeClass: fakeClass{
			cf:    cf,
			name:  cls.GetName().String(),
			super: cls.GetSuperclass(),
			iface: cls.GetInterfaces(),
		},
		cls: cls,
	}
}

func (cf *classFactory) FindClass(name string) (FakeClass, bool) {
	name, n := nameFormat(name)
	var fake FakeClass
	if cls, ok := defaultClasses[name]; ok {
		fake = cls
	} else if val, ok := cf.classes.Load(HashCode(name)); ok {
		fake = val.(FakeClass)
	} else {
		return nil, false
	}
	for i := 0; i < n; i++ {
		fake = cf.ArrayOf(fake)
	}
	return fake, true
}

func (cf *classFactory) GetClass(name string) FakeClass {
	name, n := nameFormat(name)
	if cls, ok := defaultClasses[name]; ok {
		return cf.arrayOfN(cls, n)
	}
	h := HashCode(name)
	if val, ok := cf.classes.Load(h); ok {
		return cf.arrayOfN(val.(FakeClass), n)
	}
	if cf.f != nil {
		if cls := cf.f(cf, name); cls != nil {
			return cf.arrayOfN(cls, n)
		} else if val, ok := cf.classes.Load(h); ok {
			return cf.arrayOfN(val.(FakeClass), n)
		}
	}
	cls := &fakeClass{name: name, mod: Modifier_PUBLIC}
	cls.cf = cf
	cf.classes.Store(h, cls)
	return cf.arrayOfN(cls, n)
}

func (cf *classFactory) DefineClass(name string, extends ...java.IClass) (fake FakeClass) {
	name, n := nameFormat(name)
	fake, ok := defaultClasses[name]
	if !ok {
		h := HashCode(name)
		var cls *fakeClass
		if val, ok := cf.classes.Load(h); ok {
			cls = val.(*fakeClass)
		} else {
			cls = &fakeClass{cf: cf, name: name, mod: Modifier_PUBLIC}
			cf.classes.Store(h, cls)
		}
		if len(extends) > 0 {
			cls.super = extends[0]
			cls.iface = extends[1:]
		}
		fake = cls
	}
	for i := 0; i < n; i++ {
		fake = cf.ArrayOf(fake)
	}
	return fake
}

func (cf *classFactory) ArrayOf(elem java.IClass) FakeClass {
	return arrayOf(cf, elem)
}

func (cf *classFactory) DefineMethod(cls FakeClass, name, sig string, mod Modifier) FakeMethod {
	method := &fakeMethod{cls: cls, name: name, mod: mod}
	method.parseDescriptor(cf, sig)
	return method
}

func (cf *classFactory) DefineField(cls FakeClass, name, sig string, mod Modifier) FakeField {
	field := &fakeField{cls: cls, name: name, mod: mod}
	field.parseDescriptor(cf, sig)
	return field
}

func (cf *classFactory) arrayOfN(elem FakeClass, n int) FakeClass {
	for i := 0; i < n; i++ {
		elem = cf.ArrayOf(elem)
	}
	return elem
}

func nameFormat(name string) (string, int) {
	for i := 0; ; i++ {
		switch name[i] {
		case 'Z':
			name = "boolean"
		case 'B':
			name = "byte"
		case 'C':
			name = "char"
		case 'S':
			name = "short"
		case 'I':
			name = "int"
		case 'J':
			name = "long"
		case 'F':
			name = "float"
		case 'D':
			name = "double"
		case 'V':
			name = "void"
		case '[':
			continue
		case 'L':
			name = name[i+1 : len(name)-1]
			fallthrough
		default:
			name = strings.ReplaceAll(name, "/", ".")
		}
		return name, i
	}
}

func HashCode(str string) java.JInt {
	h := fnv.New32a()
	h.Write(unsafe.Slice(unsafe.StringData(str), len(str)))
	return java.JInt(h.Sum32())
}

func definePrimitiveMethod(cls *fakeClass, name, sig string, mod Modifier) FakeMethod {
	h := HashCode(name) ^ HashCode(sig)
	method := defaultClasses.DefineMethod(cls, name, sig, mod)
	cls.methods.Store(h, method)
	return method
}

func arrayOf(cf ClassFactory, elem java.IClass) *fakeArrayClass {
	var name string
	if elem.IsPrimitive() {
		name = "[" + elem.DescriptorString().String()
	} else if elem.IsArray() {
		name = "[" + elem.GetName().String()
	} else {
		name = "[L" + elem.GetName().String() + ";"
	}
	return &fakeArrayClass{
		fakeClass: fakeClass{
			cf:    cf,
			name:  name,
			super: FakeObjectClass,
			iface: []java.IClass{FakeCloneableClass, FakeSerializableClass},
			mod:   Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
		},
		elem: elem,
	}
}

func isNil(v any) bool {
	p := (*struct{ rtype, data unsafe.Pointer })(unsafe.Pointer(&v))
	return p.rtype == nil || p.data == nil
}
