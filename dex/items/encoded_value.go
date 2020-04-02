package items

import "github.com/smm-goddess/dexparser/dex/reader"

/*
An encoded_value is an encoded piece of (nearly) arbitrary hierarchically structured data. The encoding is meant to be
both compact and straightforward to parse.
*/
type EncodedValue struct {
	ValueType uint8
	Value     []byte
}

const (
	/*
		Value formats
	*/
	VALUE_BYTE          = 0x00
	VALUE_SHORT         = 0x02
	VALUE_CHAR          = 0x03
	VALUE_INT           = 0x04
	VALUE_LONG          = 0x06
	VALUE_FLOAT         = 0x10
	VALUE_DOUBLE        = 0x11
	VALUE_METHOD_TYPE   = 0x15
	VALUE_METHOD_HANDLE = 0x16
	VALUE_STRING        = 0x17
	VALUE_TYPE          = 0x18
	VALUE_FIELD         = 0x19
	VALUE_METHOD        = 0x1a
	VALUE_ENUM          = 0x1b
	VALUE_ARRAY         = 0x1c
	VALUE_ANNOTATION    = 0x1d
	VALUE_NULL          = 0x1e
	VALUE_BOOLEAN       = 0x1f
)

func ReadEncodedValue(dexBytes []byte, offSet uint32) (EncodedValue, uint32) {
	value := EncodedValue{}
	length := uint32(1)
	valueType := dexBytes[offSet]
	argSize := (valueType >> 5) + 1
	value.ValueType = valueType
	t := valueType & 0b11111
	switch t {
	case VALUE_BYTE: // VALUE_BYTE
		v := make([]byte, 1)
		length += 1
		copy(v, dexBytes[offSet+1:offSet+2])
	case VALUE_SHORT, VALUE_CHAR, VALUE_INT, VALUE_LONG, VALUE_FLOAT, VALUE_DOUBLE, VALUE_METHOD_TYPE,
		VALUE_METHOD_HANDLE, VALUE_STRING, VALUE_TYPE, VALUE_FIELD, VALUE_METHOD, VALUE_ENUM:
		v := make([]byte, 1)
		length += uint32(argSize)
		copy(v, dexBytes[offSet+1:offSet+1+uint32(argSize)])
	case VALUE_ARRAY:
		size, count := reader.ReadUnsignedLeb128(dexBytes[offSet+1:])
		length += uint32(count)
		for i := uint32(0); i < size; i++ {
			_, cnt := ReadEncodedValue(dexBytes, offSet+length)
			length += cnt
		}
	case VALUE_ANNOTATION:
	case VALUE_NULL:
		value.Value = []byte{}
	case VALUE_BOOLEAN: // VALUE_BOOLEAN\
		var v []byte
		length += 1
		if argSize == 1 {
			v = []byte{0x01}
		} else {
			v = []byte{0x00}
		}
		value.Value = v
	}
	return value, length
}
