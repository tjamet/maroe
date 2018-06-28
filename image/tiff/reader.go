package tiff

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

var (
	tagIDLut   = map[string]uint16{}
	tagNameLut = map[uint16]string{}
	subIFDs    = map[uint16]string{
		subIFDTag: "SubIFD",
	}
)

func (t *tag) decode(b binary.ByteOrder, raw []byte) (interface{}, error) {
	switch t.dataType {
	case dtByte:
		u := make([]uint8, t.count)
		for i := uint32(0); i < t.count; i++ {
			u[i] = uint8(raw[i])
		}
		return u, nil
	case dtShort:
		u := make([]uint16, t.count)
		for i := uint32(0); i < t.count; i++ {
			u[i] = uint16(b.Uint16(raw[2*i : 2*(i+1)]))
		}
		return u, nil
	case dtLong, dtIFD:
		u := make([]uint32, t.count)
		for i := uint32(0); i < t.count; i++ {
			u[i] = uint32(b.Uint32(raw[4*i : 4*(i+1)]))
		}
		return u, nil
	case dtASCII:
		last := bytes.Index(raw, []byte{0})
		if last == -1 {
			last = len(raw)
		}
		return string(raw[:last]), nil
	default:
		return nil, fmt.Errorf("unsupported data type %d", t.dataType)
	}
}

// bytes returns the byte array for the tag value
func (t *tag) bytes(reader io.ReaderAt, b binary.ByteOrder) ([]byte, error) {
	raw := t.ptr
	dataSize, ok := lengths[t.dataType]
	if !ok {
		return nil, fmt.Errorf("invalid data type %d", t.dataType)
	}
	if t.count > math.MaxInt32/dataSize {
		return nil, fmt.Errorf("IFD data too large")
	}
	datalen := dataSize * t.count
	if datalen > 4 {
		raw = make([]byte, datalen)
		reader.ReadAt(raw, int64(b.Uint32(t.ptr)))
	}
	return raw, nil
}

func (t *Tiff) readTag(p []byte) (*tag, error) {
	if len(p) != tTagLen {
		return nil, fmt.Errorf("invalid buffer length")
	}
	tag := &tag{}
	tag.code = t.ByteOrder.Uint16(p[0:2])
	tag.dataType = t.ByteOrder.Uint16(p[2:4])
	tag.count = t.ByteOrder.Uint32(p[4:8])
	tag.ptr = p[8:tTagLen]

	if name, isSubIFD := subIFDs[tag.code]; isSubIFD {
		// https://web.archive.org/web/20060114005938/http://partners.adobe.com/public/developer/en/tiff/TIFFPM6.pdf
		// Each value is an offset (from the beginning of the TIFF file, as always) to a child
		// IFD.	Child images provide extra information for the parent image—such as a
		// subsampled version of the parent image.
		// TIFF data type 13, “IFD,” is otherwise identical to LONG, but is only used to
		// point to other valid IFDs
		bytes, err := tag.bytes(t.reader, t.ByteOrder)
		if err != nil {
			return nil, err
		}
		value, err := tag.decode(t.ByteOrder, bytes)
		if err != nil {
			return nil, err
		}
		for _, value := range value.([]uint32) {
			ifds, err := t.readIFDs(name, t.reader, value)
			if err != nil {
				return nil, err
			}
			tag.subIFDs = append(tag.subIFDs, ifds)
		}
	}
	return tag, nil
}

func (t *Tiff) readIFDs(name string, reader io.ReaderAt, offset uint32) ([]*tIFD, error) {
	ifd := &tIFD{
		name:   name,
		offset: offset,
	}
	p := make([]byte, 2)
	if _, err := reader.ReadAt(p[0:2], int64(offset)); err != nil {
		return nil, err
	}
	nTags := int(t.ByteOrder.Uint16(p[0:2]))
	tagsLen := tTagLen * nTags
	// All IFD entries are read in one chunk.
	p = make([]byte, tagsLen+tNextIFDLen)
	if _, err := reader.ReadAt(p, int64(offset+2)); err != nil {
		return nil, err
	}

	for i := 0; i < tagsLen; i += tTagLen {
		tag, err := t.readTag(p[i : i+tTagLen])
		if err != nil {
			return nil, err
		}
		ifd.tag = append(ifd.tag, tag)
	}

	ifds := []*tIFD{ifd}
	nextOffset := t.ByteOrder.Uint32(p[tagsLen : tagsLen+tNextIFDLen])
	if nextOffset != 0 {
		nextIFDs, err := t.readIFDs(name, reader, nextOffset)
		if err != nil {
			return nil, err
		}
		ifds = append(ifds, nextIFDs...)
	}
	return ifds, nil
}

// Read parses a tiff file to the Tiff structure
// Read does not read the whole file but
func Read(reader io.ReaderAt) (*Tiff, error) {
	t := &Tiff{
		reader: reader,
	}
	p := make([]byte, 8)
	if _, err := reader.ReadAt(p, 0); err != nil {
		return nil, err
	}
	switch string(p[0:2]) {
	case HeaderLittleEndian:
		t.ByteOrder = binary.LittleEndian
	case HeaderBigEndian:
		t.ByteOrder = binary.BigEndian
	default:
		return nil, fmt.Errorf("unsupported endianness %s", string(p[0:2]))
	}

	if t.ByteOrder.Uint16(p[2:4]) != tTIFFMagicNumber {
		return nil, fmt.Errorf("invalid TIFF magic number. Should be 42")
	}

	ifds, err := t.readIFDs("root", reader, t.ByteOrder.Uint32(p[4:8]))
	if err != nil {
		return nil, err
	}
	t.ifds = ifds
	return t, nil
}

// DecodeFirst returns the value of the first tag matching code encountered
func (t *Tiff) DecodeFirst(i interface{}) (interface{}, error) {
	tagName, ok := i.(string)
	var code uint16
	if ok {
		code = tagIDLut[tagName]
	} else {
		code = i.(uint16)
	}
	if t == nil {
		return nil, fmt.Errorf("no tiff data")
	}
	for _, ifd := range t.ifds {
		out, err := ifd.decodeFirst(t.reader, t.ByteOrder, code)
		if err == nil {
			return out, nil
		}
	}
	return nil, &ErrTagNotFound{code: code}
}

// RegisterTag registers a tag to be able to use names instead in decoder functions
func RegisterTag(code uint16, name string) {
	tagIDLut[name] = code
	tagNameLut[code] = name
}

func RegisterSubIFD(code uint16, name string) {
	subIFDs[code] = name
}

func DescribeTag(prefix string, code uint16) string {
	name, ok := tagNameLut[code]
	if ok {
		return fmt.Sprintf("%s%s (0x%x)", prefix, name, code)
	}
	return fmt.Sprintf("%sunknown tag %d (0x%x)", prefix, code, code)
}

func (t *Tiff) DescribeTagAndValues() {
	if t == nil {
		fmt.Println("no tiff data")
		return
	}
	for i, ifd := range t.ifds {
		ifd.DescribeTagAndValues(t.reader, t.ByteOrder, fmt.Sprintf("[%d]", i))
	}
}

func (t *Tiff) PrintCodes() {
	for i, ifd := range t.ifds {
		ifd.PrintCodes(fmt.Sprintf("[%d]", i))
	}
}

func (ifd *tIFD) DescribeTagAndValues(reader io.ReaderAt, byteOrder binary.ByteOrder, prefix string) {
	for _, tag := range ifd.tag {
		bytes, err := tag.bytes(reader, byteOrder)
		if err == nil {
			value, err := tag.decode(byteOrder, bytes)
			if err == nil {
				fmt.Println(DescribeTag(prefix, tag.code), value)
			}
		}
		for i, subIFDs := range tag.subIFDs {
			for _, subIFD := range subIFDs {
				subIFD.DescribeTagAndValues(reader, byteOrder, fmt.Sprintf("%s%s[%d]", prefix, subIFD.name, i))
			}
		}
	}
}

func (ifd *tIFD) PrintCodes(prefix string) {
	for _, tag := range ifd.tag {
		fmt.Println(DescribeTag(prefix, tag.code), tag.count, tag.dataType)
		for i, subIFDs := range tag.subIFDs {
			for _, subIFD := range subIFDs {
				subIFD.PrintCodes(fmt.Sprintf("%s%s[%d]", prefix, ifd.name, i))
			}
		}
	}
}

func (ifd *tIFD) decodeFirst(reader io.ReaderAt, byteOrder binary.ByteOrder, code uint16) (interface{}, error) {
	for _, tag := range ifd.tag {
		if tag.code == code {
			b, err := tag.bytes(reader, byteOrder)
			if err != nil {
				return nil, err
			}
			return tag.decode(byteOrder, b)
		}
	}
	for _, tag := range ifd.tag {
		out, err := tag.decodeFirst(reader, byteOrder, code)
		if err == nil {
			return out, nil
		}
	}
	return nil, &ErrTagNotFound{code: code}
}

func (tag *tag) decodeFirst(reader io.ReaderAt, byteOrder binary.ByteOrder, code uint16) (interface{}, error) {
	for _, subIFDs := range tag.subIFDs {
		for _, subIFD := range subIFDs {
			out, err := subIFD.decodeFirst(reader, byteOrder, code)
			if err == nil {
				return out, nil
			}
		}
	}
	return nil, &ErrTagNotFound{code: code}
}

type ErrTagNotFound struct {
	code uint16
}

func (e *ErrTagNotFound) Error() string {
	return fmt.Sprintf("no such tag 0x%x", e.code)
}

func (e *ErrTagNotFound) Is(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ErrTagNotFound)
	return ok
}
