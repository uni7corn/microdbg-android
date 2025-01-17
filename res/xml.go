package res

import (
	"encoding/binary"
	"encoding/xml"
	"io"
	"runtime"
	"unsafe"
)

type xmlDecoder struct {
	r       io.Reader
	strings stringPool
	resIds  []uint32
}

func (d *xmlDecoder) Token() (xml.Token, error) {
	var node ResXMLTreeNode
	err := binary.Read(d.r, binary.LittleEndian, &node)
	if err != nil {
		return nil, err
	}
	switch node.Header.Type {
	case RES_XML_START_NAMESPACE_TYPE:
		var ns ResXMLTreeNamespaceExt
		binary.Read(d.r, binary.LittleEndian, &ns)
		return xml.Comment{}, nil
	case RES_XML_END_NAMESPACE_TYPE:
		runtime.Breakpoint()
	case RES_XML_START_ELEMENT_TYPE:
		return d.getStartElement()
	case RES_XML_END_ELEMENT_TYPE:
		return d.getEndElement()
	case RES_XML_CDATA_TYPE:
		runtime.Breakpoint()
	}
	return nil, nil
}

func (d *xmlDecoder) init() error {
	err := d.strings.parse(d.r)
	if err != nil {
		return err
	}
	return d.parseResMapType()
}

func (d *xmlDecoder) parseResMapType() error {
	var res ResChunkHeader
	err := binary.Read(d.r, binary.LittleEndian, &res)
	if err != nil {
		return err
	} else if res.Type != RES_XML_RESOURCE_MAP_TYPE || res.HeaderSize != uint16(unsafe.Sizeof(ResChunkHeader{})) {
		return xml.UnmarshalError("invalid format")
	}
	d.resIds = make([]uint32, (res.Size-uint32(res.HeaderSize))/4)
	return binary.Read(d.r, binary.LittleEndian, d.resIds)
}

func (d *xmlDecoder) getStartElement() (se xml.StartElement, err error) {
	var element ResXMLTreeAttrExt
	err = binary.Read(d.r, binary.LittleEndian, &element)
	if err != nil {
		return
	}
	se.Name = xml.Name{
		Space: d.strings.Get(int(element.Ns)),
		Local: d.strings.Get(int(element.Name)),
	}
	se.Attr = make([]xml.Attr, element.AttributeCount)
	for i := range se.Attr {
		var attr ResXMLTreeAttribute
		err = binary.Read(d.r, binary.LittleEndian, &attr)
		if err != nil {
			return
		}
		se.Attr[i] = xml.Attr{
			Name: xml.Name{
				Space: d.strings.Get(int(attr.Ns)),
				Local: d.strings.Get(int(attr.Name)),
			},
		}
		if attr.RawValue == -1 {
			se.Attr[i].Value = attr.TypedValue.Display(&d.strings)
		} else {
			se.Attr[i].Value = d.strings.Get(int(attr.RawValue))
		}
	}
	return
}

func (d *xmlDecoder) getEndElement() (ee xml.EndElement, err error) {
	var element ResXMLTreeEndElementExt
	err = binary.Read(d.r, binary.LittleEndian, &element)
	if err != nil {
		return
	}
	ee.Name = xml.Name{
		Space: d.strings.Get(int(element.Ns)),
		Local: d.strings.Get(int(element.Name)),
	}
	return
}
