package items

import (
	"bytes"
	"encoding/binary"
)

/*
 *	alignment 4 byte
 */
type StringIdItem struct {
	/*
	 *offset from the start of the file to the string data for this item. The offset should be to a location in the data section,
	 *and the data should be in the format specified by "string_data_item" below. There is no alignment requirement for the offset.
	 */
	StringDataOff uint32
}

func (item StringIdItem) GetOffset() uint32 {
	return item.StringDataOff
}

func ParseStringIds(dexSource []byte, startPoint uint32, size uint32) (stringIds []StringIdItem) {
	sz := uint32(binary.Size(&StringIdItem{}))
	stringIds = make([]StringIdItem, size, size)
	_ = binary.Read(bytes.NewBuffer(dexSource[startPoint:startPoint+sz*size]), binary.LittleEndian, &stringIds)
	return
}
