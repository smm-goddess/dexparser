package items

import (
	"bytes"
	"encoding/binary"
)

type ProtoIdItem struct {
	/*
	 * index into the string_ids list for the short-form descriptor string of this prototype. The string must conform to
	 * the syntax for ShortyDescriptor, defined above, and must correspond to the return type and parameters of this item.
	 */
	ShortyIdx uint32
	/*
	 * index into the type_ids list for the return type of this prototype
	 */
	ReturnTypeIdx uint32
	/*
	 * offset from the start of the file to the list of parameter types for this prototype, or 0 if this prototype has no
	 * parameters. This offset, if non-zero, should be in the data section, and the data there should be in the format
	 * specified by "type_list" below. Additionally, there should be no reference to the type void in the list.
	 */
	ParametersOff uint32
}

func ParseProtoIds(dexSource []byte, startPoint uint32, size uint32) (protoIds []ProtoIdItem) {
	sz := uint32(binary.Size(&ProtoIdItem{}))
	protoIds = make([]ProtoIdItem, size, size)
	_ = binary.Read(bytes.NewBuffer(dexSource[startPoint:startPoint+sz*(size)]), binary.LittleEndian, &protoIds)
	return
}
