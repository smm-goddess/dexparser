package items

import (
	"bytes"
	"encoding/binary"
)

/*
An encoded_value is an encoded piece of (nearly) arbitrary hierarchically structured data. The encoding is meant to be
both compact and straightforward to parse.
*/
type EncodeValue struct {
	ValueType uint8
	Value     interface{}
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

func ReadEncodedValue(dexBytes []byte, offSet uint32) interface{} {
	valueType := dexBytes[offSet]
	argSize := (valueType >> 5) + 1
	t := valueType & 0b11111
	switch t {
	case VALUE_BYTE: // VALUE_BYTE
		bs := []byte{dexBytes[offSet+1]}
		var value uint8
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_SHORT: // VALUE_SHORT sign-extended
		var bs []byte
		if argSize == 1 {
			sign := dexBytes[offSet+1]>>3&1 == 1
			if sign {
				bs = []byte{dexBytes[offSet+1], 1}
			} else {
				bs = []byte{dexBytes[offSet+1], 0}
			}
		} else {
			bs = []byte{dexBytes[offSet+1], dexBytes[offSet+2]}
		}
		var value int16
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_CHAR: // VALUE_CHAR  zero-extendedI[
		var bs []byte
		if argSize == 1 {
			bs = []byte{dexBytes[offSet+1], 0}
		} else {
			bs = []byte{dexBytes[offSet+1], dexBytes[offSet+2]}
		}
		var value uint16
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_INT: // VALUE_INT 4 bytes sign-extended
		bs := make([]byte, 4, 4)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		sign := dexBytes[offSet+1]>>3&1 == 1
		for i := argSize; i < 4; i++ {
			if sign {
				bs[i] = 1
			} else {
				bs[i] = 0
			}
		}
		var value int32
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_LONG: // VALUE_LONG  8 bytes sign-extended
		bs := make([]byte, 8, 8)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		sign := dexBytes[offSet+1]>>3&1 == 1
		for i := argSize; i < 8; i++ {
			if sign {
				bs[i] = 1
			} else {
				bs[i] = 0
			}
		}
		var value int64
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_FLOAT: // VALUE_FLOAT
		bs := make([]byte, 4, 4)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		for i := argSize; i < 4; i++ {
			bs[i] = 0
		}
		var value float32
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_DOUBLE: // VALUE_DOUBLE
		bs := make([]byte, 8, 8)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		for i := argSize; i < 8; i++ {
			bs[i] = 0
		}
		var value float64
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_METHOD_TYPE: // VALUE_METHOD_TYPE
		bs := make([]byte, 4, 4)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		for i := argSize; i < 4; i++ {
			bs[i] = 0
		}
		var value int32
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_METHOD_HANDLE: // VALUE_METHOD_HANDLE
		bs := make([]byte, 4, 4)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		for i := argSize; i < 4; i++ {
			bs[i] = 0
		}
		var value int32
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_STRING: // VALUE_STRING TODO string_ids
		bs := make([]byte, 4, 4)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		for i := argSize; i < 4; i++ {
			bs[i] = 0
		}
		var value int32
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_TYPE: // VALUE_TYPE TODO type_ids
		bs := make([]byte, 4, 4)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		for i := argSize; i < 4; i++ {
			bs[i] = 0
		}
		var value int32
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_FIELD: // VALUE_FIELD TODO field_ids
		bs := make([]byte, 4, 4)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		for i := argSize; i < 4; i++ {
			bs[i] = 0
		}
		var value int32
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_METHOD: // VALUE_METHOD //TODO method_ids
		bs := make([]byte, 4, 4)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		for i := argSize; i < 4; i++ {
			bs[i] = 0
		}
		var value int32
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_ENUM: // VALUE_ENUM //TODO field_ids
		bs := make([]byte, 4, 4)
		for i := uint32(0); i < uint32(argSize); i++ {
			bs[i] = dexBytes[offSet+1+i]
		}
		for i := argSize; i < 4; i++ {
			bs[i] = 0
		}
		var value int32
		_ = binary.Read(bytes.NewReader(bs), binary.LittleEndian, &value)
		return value
	case VALUE_ARRAY: // VALUE_ARRAY
	// TODO array value
	case VALUE_ANNOTATION: // VALUE_ANNOTATION
	// TODO annotation value
	case VALUE_NULL: // VALUE_NULL
		return nil
	case VALUE_BOOLEAN: // VALUE_BOOLEAN
		if argSize == 1 {
			return true
		} else {
			return false
		}
	}
	return nil
}
