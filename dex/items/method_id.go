package items

import (
	"bytes"
	"encoding/binary"
)

/*
alignment:4 bytes
*/
type MethodIdItem struct {
	/*
		index into the type_ids list for the definer of this method. This must be a class or array type, and not a primitive type.
	*/
	ClassIdx uint16
	/*
		index into the proto_ids list for the prototype of this method
	*/
	ProtoIdx uint16
	/*
		index into the string_ids list for the name of this method. The string must conform to the syntax for MemberName, defined above.
	*/
	NameIdx uint32
}

func ParseMethodIds(dexSource []byte, startPoint uint32, size uint32) (methodIds []MethodIdItem, err error) {
	sz := uint32(binary.Size(&FieldIdItem{}))
	methodIds = make([]MethodIdItem, size, size)
	err = binary.Read(bytes.NewBuffer(dexSource[startPoint:startPoint+sz*size]), binary.LittleEndian, &methodIds)
	return
}
