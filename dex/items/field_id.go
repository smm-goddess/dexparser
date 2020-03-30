package items

import (
	"bytes"
	"encoding/binary"
)

/*
 *	alignment 4 byte
 */
type FieldIdItem struct {
	/*
		index into the type_ids list for the definer of this field. This must be a class type,
		and not an array or primitive type.
	*/
	ClassIdx uint16
	/*
		index into the type_ids list for the type of this field
	*/
	TypeIdx uint16
	/*
		index into the string_ids list for the name of this field. The string must conform to the syntax for MemberName,
		defined above.
	*/
	NameIdx uint32
}

func ParseFieldIds(dexSource []byte, startPoint uint32, size uint32) (fieldIds []FieldIdItem, err error) {
	sz := uint32(binary.Size(&FieldIdItem{}))
	fieldIds = make([]FieldIdItem, size, size)
	err = binary.Read(bytes.NewBuffer(dexSource[startPoint:startPoint+sz*size]), binary.LittleEndian, &fieldIds)
	return
}
