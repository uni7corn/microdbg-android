package res

import (
	"encoding/binary"
	"encoding/xml"
	"io"
	"math"
	"strconv"
	"unsafe"
)

const (
	RES_NULL_TYPE        = 0x0000
	RES_STRING_POOL_TYPE = 0x0001
	RES_TABLE_TYPE       = 0x0002
	RES_XML_TYPE         = 0x0003

	RES_XML_FIRST_CHUNK_TYPE     = 0x0100
	RES_XML_START_NAMESPACE_TYPE = 0x0100
	RES_XML_END_NAMESPACE_TYPE   = 0x0101
	RES_XML_START_ELEMENT_TYPE   = 0x0102
	RES_XML_END_ELEMENT_TYPE     = 0x0103
	RES_XML_CDATA_TYPE           = 0x0104
	RES_XML_LAST_CHUNK_TYPE      = 0x017f

	RES_XML_RESOURCE_MAP_TYPE = 0x0180

	RES_TABLE_PACKAGE_TYPE   = 0x0200
	RES_TABLE_TYPE_TYPE      = 0x0201
	RES_TABLE_TYPE_SPEC_TYPE = 0x0202
)

type ResChunkHeader struct {
	Type       uint16
	HeaderSize uint16
	Size       uint32
}

type ResStringPoolHeader struct {
	Header       ResChunkHeader
	StringCount  uint32
	StyleCount   uint32
	Flags        uint32
	StringsStart uint32
	StylesStart  uint32
}

type ResStringPoolRef int32
type ResTableRef uint32

type ResValue struct {
	Size     uint16
	Res0     uint8
	DataType uint8
	Data     uint32
}

type ResXMLTreeHeader struct {
	Header ResChunkHeader
}

type ResXMLTreeNode struct {
	Header     ResChunkHeader
	LineNumber uint32
	Comment    ResStringPoolRef
}

type ResXMLTreeCdataExt struct {
	Data      ResStringPoolRef
	TypedData ResValue
}

type ResXMLTreeNamespaceExt struct {
	Prefix ResStringPoolRef
	Uri    ResStringPoolRef
}

type ResXMLTreeEndElementExt struct {
	Ns   ResStringPoolRef
	Name ResStringPoolRef
}

type ResXMLTreeAttrExt struct {
	Ns             ResStringPoolRef
	Name           ResStringPoolRef
	AttributeStart uint16
	AttributeSize  uint16
	AttributeCount uint16
	IdIndex        uint16
	ClassIndex     uint16
	StyleIndex     uint16
}

type ResXMLTreeAttribute struct {
	Ns         ResStringPoolRef
	Name       ResStringPoolRef
	RawValue   ResStringPoolRef
	TypedValue ResValue
}

type ResTableHeader struct {
	Header       ResChunkHeader
	PackageCount uint32
}

type ResTablePackage struct {
	Header         ResChunkHeader
	Id             uint32
	Name           [128]uint16
	TypeStrings    uint32
	LastPublicType uint32
	KeyStrings     uint32
	LastPublicKey  uint32
	TypeIdOffset   uint32
}

type ResTableTypeSpec struct {
	Header     ResChunkHeader
	Id         uint8
	Res0       uint8
	Res1       uint16
	EntryCount uint32
}

type ResTableType struct {
	Header       ResChunkHeader
	Id           uint8
	Flags        uint8
	Reserved     uint16
	EntryCount   uint32
	EntriesStart uint32
	Config       ResTableConfig
}

type ResTableConfig struct {
	Size                    uint32
	Imsi                    uint32
	Locale                  uint32
	ScreenType              uint32
	Input                   uint32
	ScreenSize              uint32
	Version                 uint32
	ScreenConfig            uint32
	ScreenSizeDp            uint32
	LocaleScript            [4]byte
	LocaleVariant           [8]byte
	ScreenConfig2           uint32
	LocaleScriptWasComputed bool
	_                       [3]byte
	LocaleNumberingSystem   [8]byte
}

type ResTableEntry struct {
	Size  uint16
	Flags uint16
	Key   ResStringPoolRef
}

type ResTableMapEntry struct {
	Parent ResTableRef
	Count  uint32
}

type ResTableMap struct {
	Name  ResTableRef
	Value ResValue
}

func NewXMLDecoder(r io.Reader) (*xml.Decoder, error) {
	var header ResXMLTreeHeader
	err := binary.Read(r, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	} else if header.Header.Type != RES_XML_TYPE || header.Header.HeaderSize != uint16(unsafe.Sizeof(ResXMLTreeHeader{})) {
		return nil, xml.UnmarshalError("invalid format")
	}
	d := &xmlDecoder{r: r}
	if err = d.init(); err != nil {
		return nil, err
	}
	return xml.NewTokenDecoder(d), nil
}

func ParseTable(r io.Reader) (Table, error) {
	t := make(Table)
	err := t.parse(r)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (v ResValue) Display(pool *stringPool) string {
	const (
		TYPE_NULL      = 0x00
		TYPE_REFERENCE = 0x01
		TYPE_ATTRIBUTE = 0x02
		TYPE_STRING    = 0x03
		TYPE_FLOAT     = 0x04
		TYPE_DIMENSION = 0x05
		TYPE_FRACTION  = 0x06

		TYPE_FIRST_INT = 0x10

		TYPE_INT_DEC     = 0x10
		TYPE_INT_HEX     = 0x11
		TYPE_INT_BOOLEAN = 0x12

		TYPE_FIRST_COLOR_INT = 0x1c

		TYPE_INT_COLOR_ARGB8 = 0x1c
		TYPE_INT_COLOR_RGB8  = 0x1d
		TYPE_INT_COLOR_ARGB4 = 0x1e
		TYPE_INT_COLOR_RGB4  = 0x1f

		TYPE_LAST_COLOR_INT = 0x1f

		TYPE_LAST_INT = 0x1f
	)
	switch v.DataType {
	case TYPE_STRING:
		return pool.Get(int(v.Data))
	case TYPE_FLOAT:
		return strconv.FormatFloat(float64(math.Float32frombits(v.Data)), 'f', -1, 32)
	case TYPE_INT_DEC:
		return strconv.FormatInt(int64(v.Data), 10)
	case TYPE_REFERENCE, TYPE_INT_HEX:
		return "0x" + strconv.FormatInt(int64(v.Data), 16)
	case TYPE_INT_BOOLEAN:
		return strconv.FormatBool(v.Data != 0)
	}
	return ""
}
