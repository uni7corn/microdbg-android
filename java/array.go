package java

import (
	"errors"
	"reflect"
	"slices"
	"strconv"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
)

type FakeZArray = fakeArray[java.JBoolean]
type FakeBArray = fakeArray[java.JByte]
type FakeCArray = fakeArray[java.JChar]
type FakeSArray = fakeArray[java.JShort]
type FakeIArray = fakeArray[java.JInt]
type FakeJArray = fakeArray[java.JLong]
type FakeFArray = fakeArray[java.JFloat]
type FakeDArray = fakeArray[java.JDouble]

type fakeArray[V comparable] []V

type fakeObjectArray struct {
	fakeArray[java.IObject]
	cls java.IClass
}

func (arr fakeArray[V]) HashCode() java.JInt {
	ptr := (*struct{ data unsafe.Pointer })(unsafe.Pointer((&arr))).data
	return int32(uintptr(ptr))
}

func (arr fakeArray[V]) Equals(obj java.IObject) java.JBoolean {
	if other, ok := obj.(fakeArray[V]); ok {
		return java.JBoolean(slices.Equal(arr, other))
	}
	return false
}

func (arr fakeArray[V]) GetClass() java.IClass {
	switch any((*V)(nil)).(type) {
	case *java.JBoolean:
		return FakeZArrayClass
	case *java.JByte:
		return FakeBArrayClass
	case *java.JChar:
		return FakeCArrayClass
	case *java.JShort:
		return FakeSArrayClass
	case *java.JInt:
		return FakeIArrayClass
	case *java.JLong:
		return FakeJArrayClass
	case *java.JFloat:
		return FakeFArrayClass
	case *java.JDouble:
		return FakeDArrayClass
	default:
		panic(errors.New("[PrimitiveArray] type not supported"))
	}
}

func (arr fakeArray[V]) ToString() java.IString {
	return FakeString(arr.GetClass().GetName().String() + "@" + strconv.FormatInt(int64(arr.HashCode()), 16))
}

func (arr fakeArray[V]) Length() java.JInt {
	return java.JInt(len(arr))
}

func (arr fakeArray[V]) Get(index java.JInt) V {
	return arr[index]
}

func (arr fakeArray[V]) Set(index java.JInt, value V) {
	arr[index] = value
}

func (arr fakeArray[V]) Elements() []V {
	return arr
}

func (arr fakeObjectArray) Equals(obj java.IObject) java.JBoolean {
	if other, ok := obj.(fakeObjectArray); ok {
		return java.JBoolean(slices.Equal(arr.fakeArray, other.fakeArray))
	}
	return false
}

func (arr fakeObjectArray) GetClass() java.IClass {
	return arr.cls
}

func (arr fakeObjectArray) ToString() java.IString {
	return FakeString(arr.GetClass().GetName().String() + "@" + strconv.FormatInt(int64(arr.HashCode()), 16))
}

func BytesOf(arr []byte) java.IArray {
	return FakeBArray(unsafe.Slice((*java.JByte)(unsafe.Pointer(unsafe.SliceData(arr))), len(arr)))
}

func ArrayOf(cls FakeClass, arr any) java.IArray {
	if arr == nil {
		return nil
	}
	switch v := arr.(type) {
	case []java.JBoolean:
		return FakeZArray(v)
	case []java.JByte:
		return FakeBArray(v)
	case []java.JChar:
		return FakeCArray(v)
	case []java.JShort:
		return FakeSArray(v)
	case []java.JInt:
		return FakeIArray(v)
	case []java.JLong:
		return FakeJArray(v)
	case []java.JFloat:
		return FakeFArray(v)
	case []java.JDouble:
		return FakeDArray(v)
	case []java.IObject:
		return fakeObjectArray{fakeArray: v, cls: FakeObjectArrayClass}
	case []java.IClass:
		return fakeObjectArray{fakeArray: arrayConver(v), cls: FakeClassArrayClass}
	case []java.IString:
		return fakeObjectArray{fakeArray: arrayConver(v), cls: FakeStringArrayClass}
	case []byte:
		return FakeBArray(unsafe.Slice((*java.JByte)(unsafe.Pointer(unsafe.SliceData(v))), len(v)))
	case []string:
		arr := fakeObjectArray{fakeArray: make([]java.IObject, len(v)), cls: FakeStringArrayClass}
		for i := range v {
			arr.fakeArray[i] = FakeString(v[i])
		}
		return arr
	}
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		return nil
	}
	fake := fakeObjectArray{fakeArray: make([]java.IObject, v.Len()), cls: cls}
	for i := range fake.fakeArray {
		fake.fakeArray[i] = ToObject[java.IObject](v.Index(i).Interface())
	}
	return fake
}

func arrayConver[V java.IObject](arr []V) []java.IObject {
	r := make([]java.IObject, len(arr))
	for i := range arr {
		r[i] = arr[i]
	}
	return r
}
