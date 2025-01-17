package res

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"io"
	"unicode/utf16"
	"unsafe"
)

type stringPool struct {
	utf8    bool
	start   uint32
	offsets []uint32
	buff    []byte
}

func (sp *stringPool) Len() int {
	return len(sp.offsets)
}

func (sp *stringPool) Get(index int) string {
	if index < 0 || index >= len(sp.offsets) {
		return ""
	}
	offset := sp.start + sp.offsets[index]
	r := bytes.NewReader(sp.buff[offset:])
	var lens [2]byte
	_, err := io.ReadFull(r, lens[:])
	if err != nil {
		return ""
	}
	if sp.utf8 {
		buf := make([]byte, lens[1])
		io.ReadFull(r, buf)
		return string(buf)
	} else {
		buf := make([]uint16, lens[0])
		binary.Read(r, binary.LittleEndian, buf)
		return string(utf16.Decode(buf))
	}
}

func (sp *stringPool) Range(yield func(int, string) bool) {
	for i := 0; i < sp.Len(); i++ {
		if !yield(i, sp.Get(i)) {
			break
		}
	}
}

func (sp *stringPool) parse(r io.Reader) error {
	var pool ResStringPoolHeader
	err := binary.Read(r, binary.LittleEndian, &pool)
	if err != nil {
		return err
	} else if pool.Header.Type != RES_STRING_POOL_TYPE || pool.Header.HeaderSize != uint16(unsafe.Sizeof(ResStringPoolHeader{})) {
		return xml.UnmarshalError("invalid format")
	}
	const (
		SORTED_FLAG = 1 << 0
		UTF8_FLAG   = 1 << 8
	)
	sp.utf8 = pool.Flags&UTF8_FLAG != 0
	sp.start = pool.StringsStart - uint32(pool.Header.HeaderSize) - (pool.StringCount+pool.StyleCount)*4
	sp.offsets = make([]uint32, pool.StringCount)
	err = binary.Read(r, binary.LittleEndian, sp.offsets)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, make([]uint32, pool.StyleCount))
	if err != nil {
		return err
	}
	size := pool.Header.Size - uint32(pool.Header.HeaderSize) - (pool.StringCount+pool.StyleCount)*4
	sp.buff = make([]byte, size)
	_, err = io.ReadFull(r, sp.buff)
	return err
}
