package java

import (
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type FakeByte java.JByte

func (b FakeByte) GetClass() java.IClass {
	return FakeByteClass
}

func (b FakeByte) HashCode() java.JInt {
	return java.JInt(b)
}

func (b FakeByte) Equals(obj java.IObject) java.JBoolean {
	return b == obj
}

func (b FakeByte) ToString() java.IString {
	return FakeString(strconv.FormatInt(int64(b), 10))
}

func init() {
	FakeByteClass.Set("TYPE", FakeByteTYPE)
	definePrimitiveMethod(FakeByteClass, "byteValue", "()B", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JByte(obj.(FakeByte))
	})
	definePrimitiveMethod(FakeByteClass, "shortValue", "()S", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JShort(obj.(FakeByte))
	})
	definePrimitiveMethod(FakeByteClass, "intValue", "()I", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JInt(obj.(FakeByte))
	})
	definePrimitiveMethod(FakeByteClass, "longValue", "()J", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JLong(obj.(FakeByte))
	})
	definePrimitiveMethod(FakeByteClass, "floatValue", "()F", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JFloat(obj.(FakeByte))
	})
	definePrimitiveMethod(FakeByteClass, "doubleValue", "()D", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JDouble(obj.(FakeByte))
	})
}
