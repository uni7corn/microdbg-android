package java

import (
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type FakeLong java.JLong

func (l FakeLong) GetClass() java.IClass {
	return FakeLongClass
}

func (l FakeLong) HashCode() java.JInt {
	v := uint64(l)
	return java.JInt(v ^ (v >> 32))
}

func (l FakeLong) Equals(obj java.IObject) java.JBoolean {
	return l == obj
}

func (l FakeLong) ToString() java.IString {
	return FakeString(strconv.FormatInt(int64(l), 10))
}

func init() {
	FakeLongClass.Set("TYPE", FakeLongTYPE)
	definePrimitiveMethod(FakeLongClass, "byteValue", "()B", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JByte(obj.(FakeLong))
	})
	definePrimitiveMethod(FakeLongClass, "shortValue", "()S", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JShort(obj.(FakeLong))
	})
	definePrimitiveMethod(FakeLongClass, "intValue", "()I", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JInt(obj.(FakeLong))
	})
	definePrimitiveMethod(FakeLongClass, "longValue", "()J", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JLong(obj.(FakeLong))
	})
	definePrimitiveMethod(FakeLongClass, "floatValue", "()F", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JFloat(obj.(FakeLong))
	})
	definePrimitiveMethod(FakeLongClass, "doubleValue", "()D", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JDouble(obj.(FakeLong))
	})
}
