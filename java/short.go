package java

import (
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type FakeShort java.JShort

func (s FakeShort) GetClass() java.IClass {
	return FakeShortClass
}

func (s FakeShort) HashCode() java.JInt {
	return java.JInt(s)
}

func (s FakeShort) Equals(obj java.IObject) java.JBoolean {
	return s == obj
}

func (s FakeShort) ToString() java.IString {
	return FakeString(strconv.FormatInt(int64(s), 10))
}

func init() {
	FakeShortClass.Set("TYPE", FakeShortTYPE)
	definePrimitiveMethod(FakeShortClass, "byteValue", "()B", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JByte(obj.(FakeShort))
	})
	definePrimitiveMethod(FakeShortClass, "shortValue", "()S", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JShort(obj.(FakeShort))
	})
	definePrimitiveMethod(FakeShortClass, "intValue", "()I", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JInt(obj.(FakeShort))
	})
	definePrimitiveMethod(FakeShortClass, "longValue", "()J", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JLong(obj.(FakeShort))
	})
	definePrimitiveMethod(FakeShortClass, "floatValue", "()F", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JFloat(obj.(FakeShort))
	})
	definePrimitiveMethod(FakeShortClass, "doubleValue", "()D", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JDouble(obj.(FakeShort))
	})
}
