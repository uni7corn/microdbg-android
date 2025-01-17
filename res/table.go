package res

import (
	"encoding/binary"
	"errors"
	"io"
	"strings"
	"unicode/utf16"
	"unsafe"
)

type Table map[int]TablePackage

type TablePackage struct {
	Name  string
	Types []TableType
}

type TableType struct {
	Name    string
	Entries []TableEntry
}

type TableEntry struct {
	Name  string
	Value TableValue
}

type TableValue interface {
	IsComplex() bool
}

type Value struct {
	DataType uint8
	Data     uint32
	String   string
}

type Complex []struct {
	Name  ResTableRef
	Value Value
}

func (t Table) Get(id int) (TableEntry, bool) {
	pkg, ok := t[(id>>24)&0xFF]
	if !ok {
		return TableEntry{}, false
	}
	tid := (id >> 16) & 0xFF
	if tid <= 0 || tid >= len(pkg.Types) {
		return TableEntry{}, false
	}
	entries := pkg.Types[tid].Entries
	id &= 0xFFFF
	if id < 0 || id >= len(entries) {
		return TableEntry{}, false
	}
	return entries[id], true
}

func (t Table) parse(r io.Reader) error {
	var header ResTableHeader
	err := binary.Read(r, binary.LittleEndian, &header)
	if err != nil {
		return err
	} else if header.Header.Type != RES_TABLE_TYPE || header.Header.HeaderSize != uint16(unsafe.Sizeof(ResTableHeader{})) {
		return errors.New("[ResTableHeader] invalid format")
	}
	var pool stringPool
	err = pool.parse(r)
	if err != nil {
		return err
	}
	for i := 0; i < int(header.PackageCount); i++ {
		var pkg ResTablePackage
		err = binary.Read(r, binary.LittleEndian, &pkg)
		if err != nil {
			return err
		} else if pkg.Header.Type != RES_TABLE_PACKAGE_TYPE || pkg.Header.HeaderSize != uint16(unsafe.Sizeof(ResTablePackage{})) {
			return errors.New("[ResTablePackage] invalid format")
		}
		r := io.LimitReader(r, int64(pkg.Header.Size)-int64(pkg.Header.HeaderSize))
		var types, keys stringPool
		if err = types.parse(r); err != nil {
			return err
		} else if err = keys.parse(r); err != nil {
			return err
		}
		tp := TablePackage{
			Name:  strings.TrimRight(string(utf16.Decode(pkg.Name[:])), "\x00"),
			Types: make([]TableType, types.Len()+1),
		}
		for i, v := range types.Range {
			tp.Types[i+1] = TableType{Name: v}
		}
		var chunk ResChunkHeader
		for binary.Read(r, binary.LittleEndian, &chunk) == nil {
			r := io.LimitReader(r, int64(chunk.Size)-int64(unsafe.Sizeof(ResChunkHeader{})))
			if chunk.Type == RES_TABLE_TYPE_TYPE {
				var typ ResTableType
				_, err := io.ReadFull(r, unsafe.Slice((*byte)(unsafe.Pointer(&typ)), unsafe.Sizeof(ResTableType{}))[unsafe.Sizeof(ResChunkHeader{}):])
				if err != nil {
					return err
				}
				io.CopyN(io.Discard, r, int64(typ.EntriesStart)-int64(chunk.HeaderSize))
				entries := make([]TableEntry, typ.EntryCount)
				err = t.parseEntry(r, &pool, &keys, entries)
				id := int(typ.Id)
				tp.Types[id].Entries = append(tp.Types[id].Entries, entries...)
				if err != nil && err != io.EOF {
					return err
				}
			}
			io.Copy(io.Discard, r)
		}
		t[int(pkg.Id)] = tp
	}
	return nil
}

func (t Table) parseEntry(r io.Reader, strings, keys *stringPool, entries []TableEntry) error {
	for i := range entries {
		var entry ResTableEntry
		err := binary.Read(r, binary.LittleEndian, &entry)
		if err != nil {
			return err
		} else if entry.Size != uint16(unsafe.Sizeof(ResTableEntry{})) && entry.Size != uint16(unsafe.Sizeof(ResTableEntry{})+unsafe.Sizeof(ResTableMapEntry{})) {
			return errors.New("[ResTableEntry] invalid format")
		}
		te := TableEntry{Name: keys.Get(int(entry.Key))}
		const (
			FLAG_COMPLEX = 1 << iota
			FLAG_PUBLIC
			FLAG_WEAK
		)
		if entry.Flags&FLAG_COMPLEX == 0 {
			var value ResValue
			err = binary.Read(r, binary.LittleEndian, &value)
			if err != nil {
				return err
			}
			te.Value = Value{
				DataType: value.DataType,
				Data:     value.Data,
				String:   value.Display(strings),
			}
		} else {
			var mapEntry ResTableMapEntry
			err = binary.Read(r, binary.LittleEndian, &mapEntry)
			if err != nil {
				return err
			}
			ec := make(Complex, mapEntry.Count)
			te.Value = ec
			for i := range ec {
				var mp ResTableMap
				err := binary.Read(r, binary.LittleEndian, &mp)
				if err != nil {
					return err
				}
				ec[i].Name = mp.Name
				ec[i].Value = Value{
					DataType: mp.Value.DataType,
					Data:     mp.Value.Data,
					String:   mp.Value.Display(strings),
				}
			}
		}
		entries[i] = te
	}
	return nil
}

func (v Value) IsComplex() bool {
	return false
}

func (c Complex) IsComplex() bool {
	return true
}
