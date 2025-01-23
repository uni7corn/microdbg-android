package java

import (
	"math"
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type FakeDouble java.JDouble

func (d FakeDouble) GetClass() java.IClass {
	return FakeDoubleClass
}

func (d FakeDouble) HashCode() java.JInt {
	bits := math.Float64bits(float64(d))
	return (java.JInt)(bits ^ (bits >> 32))
}

func (d FakeDouble) Equals(obj java.IObject) java.JBoolean {
	return d == obj
}

func (d FakeDouble) ToString() java.IString {
	return FakeString(strconv.FormatFloat(float64(d), 'f', -1, 32))
}

func init() {
	FakeDoubleClass.Set("TYPE", FakeDoubleTYPE)
	definePrimitiveMethod(FakeDoubleClass, "byteValue", "()B", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JByte(obj.(FakeDouble))
	})
	definePrimitiveMethod(FakeDoubleClass, "shortValue", "()S", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JShort(obj.(FakeDouble))
	})
	definePrimitiveMethod(FakeDoubleClass, "intValue", "()I", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JInt(obj.(FakeDouble))
	})
	definePrimitiveMethod(FakeDoubleClass, "longValue", "()J", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JLong(obj.(FakeDouble))
	})
	definePrimitiveMethod(FakeDoubleClass, "floatValue", "()F", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JFloat(obj.(FakeDouble))
	})
	definePrimitiveMethod(FakeDoubleClass, "doubleValue", "()D", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JDouble(obj.(FakeDouble))
	})
}
