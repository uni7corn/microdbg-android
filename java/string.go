package java

import (
	"strconv"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
)

type FakeString string

func (str FakeString) GetClass() java.IClass {
	return nil
}

func (str FakeString) HashCode() java.JInt {
	ptr := (*struct{ rtype, data unsafe.Pointer })(unsafe.Pointer((&str))).data
	return int32(uintptr(ptr))
}

func (str FakeString) Equals(obj java.IObject) java.JBoolean {
	return str == obj
}

func (str FakeString) ToString() java.IString {
	return FakeString(str.GetClass().GetName().String() + "@" + strconv.FormatInt(int64(str.HashCode()), 16))
}

func (str FakeString) Length() java.JInt {
	return java.JInt(len(str))
}

func (str FakeString) String() string {
	return string(str)
}
