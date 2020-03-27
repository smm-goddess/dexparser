package items

import (
	"bytes"
	"encoding/binary"
)

/*
 *	alignment 4 byte
 */
type TypeIdItem struct {
	/*
	 * index into the string_ids list for the descriptor string of this type. The string must conform to the syntax for
	 * TypeDescriptor, defined above.
	 */
	DescriptorIdx uint32
}

func ParseTypeIdIds(dexSource []byte, startPoint uint32, size uint32) (typeIds []TypeIdItem) {
	sz := uint32(binary.Size(&StringIdItem{}))
	typeIds = make([]TypeIdItem, size, size)
	for i := uint32(0); i < size; i++ {
		var item TypeIdItem
		_ = binary.Read(bytes.NewBuffer(dexSource[startPoint+sz*i:startPoint+sz*(i+1)]), binary.LittleEndian, &item)
		typeIds[i] = item
	}
	return
}

func Descriptor2Class(desc []byte) string {
	if desc[0] == 'V' {
		return "void"
	}
	arrayLength := 0
	for desc[arrayLength] == '[' {
		arrayLength++
	}
	buffer := bytes.Buffer{}
	switch desc[arrayLength] {
	case 'Z':
		buffer.WriteString("boolean")
	case 'B':
		buffer.WriteString("byte")
	case 'S':
		buffer.WriteString("short")
	case 'C':
		buffer.WriteString("char")
	case 'I':
		buffer.WriteString("int")
	case 'J':
		buffer.WriteString("long")
	case 'F':
		buffer.WriteString("float")
	case 'D':
		buffer.WriteString("double")
	case 'L':
		buffer.Write(desc[arrayLength+1 : len(desc)-1])
	}
	for arrayLength > 0 {
		buffer.WriteString("[]")
		arrayLength--
	}
	return buffer.String()
}
