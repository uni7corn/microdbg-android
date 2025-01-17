package internal

import (
	"io"
	"iter"
	"unsafe"

	"golang.org/x/exp/constraints"
)

type Rel[I constraints.Integer] struct {
	Off  I
	Info I
}

type Rela[I constraints.Integer] struct {
	Rel[I]
	Addend I
}

func newRelocIter[I constraints.Integer](r io.Reader) iter.Seq[Rela[I]] {
	const (
		RELOCATION_GROUPED_BY_INFO_FLAG = 1 << iota
		RELOCATION_GROUPED_BY_OFFSET_DELTA_FLAG
		RELOCATION_GROUPED_BY_ADDEND_FLAG
		RELOCATION_GROUP_HAS_ADDEND_FLAG
	)

	var reloc Rela[I]
	count, err := leb128[I](r)
	if err != nil {
		return nil
	}
	reloc.Off, err = leb128[I](r)
	if err != nil {
		return nil
	}
	var index, groupIndex, groupSize, groupFlags, groupOffset I
	readGroup := func() (err error) {
		groupSize, err = leb128[I](r)
		if err != nil {
			return
		}
		groupFlags, err = leb128[I](r)
		if err != nil {
			return
		}
		if groupFlags&RELOCATION_GROUPED_BY_OFFSET_DELTA_FLAG != 0 {
			groupOffset, err = leb128[I](r)
			if err != nil {
				return
			}
		}
		if groupFlags&RELOCATION_GROUPED_BY_INFO_FLAG != 0 {
			reloc.Info, err = leb128[I](r)
			if err != nil {
				return
			}
		}
		switch groupFlags & (RELOCATION_GROUP_HAS_ADDEND_FLAG | RELOCATION_GROUPED_BY_ADDEND_FLAG) {
		case RELOCATION_GROUP_HAS_ADDEND_FLAG | RELOCATION_GROUPED_BY_ADDEND_FLAG:
			v, err := leb128[I](r)
			if err != nil {
				break
			}
			reloc.Addend += v
		case RELOCATION_GROUP_HAS_ADDEND_FLAG:
			reloc.Addend = 0
		}
		groupIndex = 0
		return nil
	}
	return func(yield func(Rela[I]) bool) {
		for index < count {
			if groupIndex == groupSize {
				if readGroup() != nil {
					break
				}
			}
			if groupFlags&RELOCATION_GROUPED_BY_OFFSET_DELTA_FLAG != 0 {
				reloc.Off += groupOffset
			} else {
				v, err := leb128[I](r)
				if err != nil {
					break
				}
				reloc.Off += v
			}
			if groupFlags&RELOCATION_GROUPED_BY_INFO_FLAG == 0 {
				v, err := leb128[I](r)
				if err != nil {
					break
				}
				reloc.Info = v
			}
			if groupFlags&(RELOCATION_GROUP_HAS_ADDEND_FLAG|RELOCATION_GROUPED_BY_ADDEND_FLAG) == RELOCATION_GROUP_HAS_ADDEND_FLAG {
				v, err := leb128[I](r)
				if err != nil {
					break
				}
				reloc.Addend += v
			}
			index++
			groupIndex++
			if !yield(reloc) {
				break
			}
		}
	}
}

func leb128[I constraints.Integer](r io.Reader) (value I, err error) {
	buf := [1]byte{128}
	shift := 0
	for ; buf[0]&128 != 0; shift += 7 {
		_, err = r.Read(buf[:])
		if err != nil {
			return
		}
		value |= (I(buf[0] & 127)) << shift
	}
	if shift < (int(unsafe.Sizeof(value))*8) && buf[0]&64 != 0 {
		value |= -(I(1) << shift)
	}
	return
}
