package java

import (
	"errors"
	"unsafe"

	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg/emulator"
)

type uniPtr struct {
	p emulator.Pointer
}

type ptr[V any] struct {
	uniPtr
}

type jvaluePtr struct {
	ptr[JValue]
}

type JValue uint64

func newPtr(p emulator.Pointer) java.Ptr {
	return uniPtr{p}
}

func newTypePtr[V any](p emulator.Pointer) java.TypePtr[V] {
	return ptr[V]{uniPtr{p}}
}

func newJValuePtr(p emulator.Pointer) java.TypePtr[java.JValue] {
	return jvaluePtr{ptr[JValue]{uniPtr{p}}}
}

func (ptr uniPtr) Address() uintptr {
	return uintptr(ptr.p.Address())
}

func (ptr ptr[V]) Get(index int) (V, error) {
	var value V
	size := uint64(unsafe.Sizeof(value))
	err := ptr.p.Add(uint64(index)*size).MemReadPtr(size, unsafe.Pointer(&value))
	return value, err
}

func (ptr ptr[V]) Set(index int, value V) error {
	size := uint64(unsafe.Sizeof(value))
	return ptr.p.Add(uint64(index)*size).MemWritePtr(size, unsafe.Pointer(&value))
}

func (ptr ptr[V]) ReadAt(b []V, off int64) (int, error) {
	count := len(b)
	if count == 0 {
		return 0, nil
	}
	size := count * int(unsafe.Sizeof(b[0]))
	err := ptr.p.Add(uint64(off)).MemReadPtr(uint64(size), unsafe.Pointer(unsafe.SliceData(b)))
	return size, err
}

func (ptr ptr[V]) WriteAt(b []V, off int64) (int, error) {
	count := len(b)
	if count == 0 {
		return 0, nil
	}
	size := count * int(unsafe.Sizeof(b[0]))
	err := ptr.p.Add(uint64(off)).MemWritePtr(uint64(size), unsafe.Pointer(unsafe.SliceData(b)))
	return size, err
}

func (ptr jvaluePtr) Get(index int) (java.JValue, error) {
	return ptr.ptr.Get(index)
}

func (ptr jvaluePtr) Set(index int, value java.JValue) error {
	return errors.ErrUnsupported
}

func (ptr jvaluePtr) ReadAt(b []java.JValue, off int64) (int, error) {
	count := len(b)
	if count == 0 {
		return 0, nil
	}
	arr := make([]JValue, count)
	n, err := ptr.ptr.ReadAt(arr, off)
	if err != nil {
		return 0, err
	}
	for i, v := range arr {
		b[i] = v
	}
	return n, nil
}

func (ptr jvaluePtr) WriteAt(b []java.JValue, off int64) (int, error) {
	return 0, errors.ErrUnsupported
}

func (v JValue) JBoolean() java.JBoolean {
	return *(*java.JBoolean)(unsafe.Pointer(&v))
}

func (v JValue) JByte() java.JByte {
	return *(*java.JByte)(unsafe.Pointer(&v))
}

func (v JValue) JChar() java.JChar {
	return *(*java.JChar)(unsafe.Pointer(&v))
}

func (v JValue) JShort() java.JShort {
	return *(*java.JShort)(unsafe.Pointer(&v))
}

func (v JValue) JInt() java.JInt {
	return *(*java.JInt)(unsafe.Pointer(&v))
}

func (v JValue) JLong() java.JLong {
	return *(*java.JLong)(unsafe.Pointer(&v))
}

func (v JValue) JFloat() java.JFloat {
	return *(*java.JFloat)(unsafe.Pointer(&v))
}

func (v JValue) JDouble() java.JDouble {
	return *(*java.JDouble)(unsafe.Pointer(&v))
}

func (v JValue) JObject() java.JObject {
	return *(*Ref)(unsafe.Pointer(&v))
}
