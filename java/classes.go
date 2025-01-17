package java

import java "github.com/wnxd/microdbg-java"

type defaultFactory map[string]FakeClass

var (
	FakeBooleanTYPE = &fakeClass{
		name: "boolean",
		mod:  Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeByteTYPE = &fakeClass{
		name: "byte",
		mod:  Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeCharTYPE = &fakeClass{
		name: "char",
		mod:  Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeShortTYPE = &fakeClass{
		name: "short",
		mod:  Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeIntTYPE = &fakeClass{
		name: "int",
		mod:  Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeLongTYPE = &fakeClass{
		name: "long",
		mod:  Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeFloatTYPE = &fakeClass{
		name: "float",
		mod:  Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeDoubleTYPE = &fakeClass{
		name: "double",
		mod:  Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeVoidTYPE = &fakeClass{
		name: "void",
		mod:  Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeObjectClass = &fakeClass{
		name: "java.lang.Object",
		mod:  Modifier_PUBLIC,
	}
	FakeClassClass = &fakeClass{
		name:  "java.lang.Class",
		super: FakeObjectClass,
		iface: []java.IClass{FakeSerializableClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeStringClass = &fakeClass{
		name:  "java.lang.String",
		super: FakeObjectClass,
		iface: []java.IClass{FakeSerializableClass, FakeComparableClass, FakeCharSequenceClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL | Modifier_ABSTRACT,
	}
	FakeBooleanClass = &fakeClass{
		name:  "java.lang.Boolean",
		super: FakeObjectClass,
		iface: []java.IClass{FakeSerializableClass, FakeComparableClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeByteClass = &fakeClass{
		name:  "java.lang.Byte",
		super: FakeNumberClass,
		iface: []java.IClass{FakeComparableClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeCharacterClass = &fakeClass{
		name:  "java.lang.Character",
		super: FakeObjectClass,
		iface: []java.IClass{FakeSerializableClass, FakeComparableClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeShortClass = &fakeClass{
		name:  "java.lang.Short",
		super: FakeNumberClass,
		iface: []java.IClass{FakeComparableClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeIntegerClass = &fakeClass{
		name:  "java.lang.Integer",
		super: FakeNumberClass,
		iface: []java.IClass{FakeComparableClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeLongClass = &fakeClass{
		name:  "java.lang.Long",
		super: FakeNumberClass,
		iface: []java.IClass{FakeComparableClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeFloatClass = &fakeClass{
		name:  "java.lang.Float",
		super: FakeNumberClass,
		iface: []java.IClass{FakeComparableClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeDoubleClass = &fakeClass{
		name:  "java.lang.Double",
		super: FakeNumberClass,
		iface: []java.IClass{FakeComparableClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeNumberClass = &fakeClass{
		name:  "java.lang.Number",
		super: FakeObjectClass,
		iface: []java.IClass{FakeSerializableClass},
		mod:   Modifier_PUBLIC | Modifier_ABSTRACT,
	}
	FakeIterableClass = &fakeClass{
		name: "java.lang.Iterable",
		mod:  Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeMethodClass = &fakeClass{
		name:  "java.lang.reflect.Method",
		super: FakeExecutableClass,
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeConstructorClass = &fakeClass{
		name:  "java.lang.reflect.Constructor",
		super: FakeExecutableClass,
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeExecutableClass = &fakeClass{
		name:  "java.lang.reflect.Executable",
		super: FakeAccessibleObjectClass,
		iface: []java.IClass{FakeMemberClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeFieldClass = &fakeClass{
		name:  "java.lang.reflect.Field",
		super: FakeAccessibleObjectClass,
		iface: []java.IClass{FakeMemberClass},
		mod:   Modifier_PUBLIC | Modifier_FINAL,
	}
	FakeAccessibleObjectClass = &fakeClass{
		name:  "java.lang.reflect.AccessibleObject",
		super: FakeObjectClass,
		mod:   Modifier_PUBLIC,
	}
	FakeMemberClass = &fakeClass{
		name:  "java.lang.reflect.Member",
		super: FakeObjectClass,
		mod:   Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeCloneableClass = &fakeClass{
		name: "java.lang.Cloneable",
		mod:  Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeSerializableClass = &fakeClass{
		name: "java.io.Serializable",
		mod:  Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeComparableClass = &fakeClass{
		name: "java.lang.Comparable",
		mod:  Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeCharSequenceClass = &fakeClass{
		name: "java.lang.CharSequence",
		mod:  Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}

	FakeZArrayClass      = arrayOf(nil, FakeBooleanTYPE)
	FakeBArrayClass      = arrayOf(nil, FakeByteTYPE)
	FakeCArrayClass      = arrayOf(nil, FakeCharTYPE)
	FakeSArrayClass      = arrayOf(nil, FakeShortTYPE)
	FakeIArrayClass      = arrayOf(nil, FakeIntTYPE)
	FakeJArrayClass      = arrayOf(nil, FakeLongTYPE)
	FakeFArrayClass      = arrayOf(nil, FakeFloatTYPE)
	FakeDArrayClass      = arrayOf(nil, FakeDoubleTYPE)
	FakeObjectArrayClass = arrayOf(nil, FakeObjectClass)
	FakeClassArrayClass  = arrayOf(nil, FakeClassClass)
	FakeStringArrayClass = arrayOf(nil, FakeStringClass)

	FakeIteratorClass = &fakeClass{
		name: "java.util.Iterator",
		mod:  Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeCollectionClass = &fakeClass{
		name:  "java.util.Collection",
		iface: []java.IClass{FakeIterableClass},
		mod:   Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeSetClass = &fakeClass{
		name:  "java.util.Set",
		iface: []java.IClass{FakeCollectionClass},
		mod:   Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeMapClass = &fakeClass{
		name: "java.util.Map",
		mod:  Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeMapEntryClass = &fakeClass{
		name: "java.util.Map$Entry",
		mod:  Modifier_PUBLIC | Modifier_INTERFACE | Modifier_ABSTRACT,
	}
	FakeHashMapClass = &fakeClass{
		name:  "java.util.HashMap",
		super: FakeObjectClass,
		iface: []java.IClass{FakeMapClass, FakeCloneableClass, FakeSerializableClass},
		mod:   Modifier_PUBLIC,
	}

	defaultClasses = make(defaultFactory)
)

func (defaultFactory) WrapClass(java.IClass) FakeClass {
	return nil
}

func (defaultFactory) FindClass(name string) (FakeClass, bool) {
	cls, ok := defaultClasses[name]
	return cls, ok
}

func (defaultFactory) GetClass(name string) FakeClass {
	return defaultClasses[name]
}

func (defaultFactory) DefineClass(name string, extends ...java.IClass) FakeClass {
	cls := &fakeClass{
		name: name,
		mod:  Modifier_PUBLIC,
	}
	if len(extends) > 0 {
		cls.super = extends[0]
		cls.iface = extends[1:]
	}
	defaultClasses[name] = cls
	return cls
}

func (cf defaultFactory) ArrayOf(elem java.IClass) FakeClass {
	cls := arrayOf(cf, elem)
	cls.cf = nil
	return cls
}

func (cf defaultFactory) DefineMethod(cls FakeClass, name, sig string, mod Modifier) FakeMethod {
	method := &fakeMethod{cls: cls, name: name, mod: mod}
	method.parseDescriptor(cf, sig)
	return method
}

func (cf defaultFactory) DefineField(cls FakeClass, name, sig string, mod Modifier) FakeField {
	field := &fakeField{cls: cls, name: name, mod: mod}
	field.parseDescriptor(cf, sig)
	return field
}

func init() {
	list := []FakeClass{FakeBooleanTYPE, FakeByteTYPE, FakeCharTYPE, FakeShortTYPE, FakeIntTYPE, FakeLongTYPE, FakeFloatTYPE, FakeDoubleTYPE, FakeVoidTYPE, FakeObjectClass, FakeClassClass, FakeStringClass, FakeBooleanClass, FakeByteClass, FakeCharacterClass, FakeShortClass, FakeIntegerClass, FakeLongClass, FakeFloatClass, FakeDoubleClass, FakeNumberClass, FakeIterableClass, FakeMethodClass, FakeConstructorClass, FakeExecutableClass, FakeFieldClass, FakeAccessibleObjectClass, FakeMemberClass, FakeCloneableClass, FakeSerializableClass, FakeComparableClass, FakeCharSequenceClass, FakeZArrayClass, FakeBArrayClass, FakeCArrayClass, FakeSArrayClass, FakeIArrayClass, FakeJArrayClass, FakeFArrayClass, FakeDArrayClass, FakeObjectArrayClass, FakeClassArrayClass, FakeStringArrayClass, FakeIteratorClass, FakeCollectionClass, FakeSetClass, FakeMapClass, FakeMapEntryClass, FakeHashMapClass}

	for _, cls := range list {
		defaultClasses[cls.GetName().String()] = cls
	}
}
