package java

import (
	"strconv"

	java "github.com/wnxd/microdbg-java"
)

type FakeBoolean java.JBoolean

func (b FakeBoolean) GetClass() java.IClass {
	return FakeBooleanClass
}

func (b FakeBoolean) HashCode() java.JInt {
	if b {
		return 1231
	}
	return 1237
}

func (b FakeBoolean) Equals(obj java.IObject) java.JBoolean {
	return b == obj
}

func (b FakeBoolean) ToString() java.IString {
	return FakeString(strconv.FormatBool(bool(b)))
}

func init() {
	FakeBooleanClass.Set("TYPE", FakeBooleanTYPE)
	definePrimitiveMethod(FakeBooleanClass, "booleanValue", "()Z", Modifier_PUBLIC).BindCall(func(obj java.IObject, _ ...any) any {
		return java.JBoolean(obj.(FakeBoolean))
	})
}
