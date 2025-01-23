package java

import (
	"math"
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type FakeFloat java.JFloat

func (f FakeFloat) GetClass() java.IClass {
	return FakeFloatClass
}

func (f FakeFloat) HashCode() java.JInt {
	return java.JInt(math.Float32bits(float32(f)))
}

func (f FakeFloat) Equals(obj java.IObject) java.JBoolean {
	return f == obj
}

func (f FakeFloat) ToString() java.IString {
	return FakeString(strconv.FormatFloat(float64(f), 'f', -1, 32))
}

func init() {
	FakeFloatClass.Set("TYPE", FakeFloatTYPE)
	definePrimitiveMethod(FakeFloatClass, "byteValue", "()B", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JByte(obj.(FakeFloat))
	})
	definePrimitiveMethod(FakeFloatClass, "shortValue", "()S", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JShort(obj.(FakeFloat))
	})
	definePrimitiveMethod(FakeFloatClass, "intValue", "()I", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JInt(obj.(FakeFloat))
	})
	definePrimitiveMethod(FakeFloatClass, "longValue ", "()J", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JLong(obj.(FakeFloat))
	})
	definePrimitiveMethod(FakeFloatClass, "floatValue", "()F", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JFloat(obj.(FakeFloat))
	})
	definePrimitiveMethod(FakeFloatClass, "doubleValue", "()D", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JDouble(obj.(FakeFloat))
	})
}
