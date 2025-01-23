package java

import (
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type FakeInteger java.JInt

func (i FakeInteger) GetClass() java.IClass {
	return FakeIntegerClass
}

func (i FakeInteger) HashCode() java.JInt {
	return java.JInt(i)
}

func (i FakeInteger) Equals(obj java.IObject) java.JBoolean {
	return i == obj
}

func (i FakeInteger) ToString() java.IString {
	return FakeString(strconv.FormatInt(int64(i), 10))
}

func init() {
	FakeIntegerClass.Set("TYPE", FakeIntTYPE)
	definePrimitiveMethod(FakeIntegerClass, "byteValue", "()B", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JByte(obj.(FakeInteger))
	})
	definePrimitiveMethod(FakeIntegerClass, "shortValue", "()S", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JShort(obj.(FakeInteger))
	})
	definePrimitiveMethod(FakeIntegerClass, "intValue", "()I", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JInt(obj.(FakeInteger))
	})
	definePrimitiveMethod(FakeIntegerClass, "longValue", "()J", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JLong(obj.(FakeInteger))
	})
	definePrimitiveMethod(FakeIntegerClass, "floatValue", "()F", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JFloat(obj.(FakeInteger))
	})
	definePrimitiveMethod(FakeIntegerClass, "doubleValue", "()D", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JDouble(obj.(FakeInteger))
	})
}
