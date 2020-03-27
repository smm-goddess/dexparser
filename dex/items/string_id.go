package items

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

type StringIdItem struct {
	StringDataOff uint32 // offset from the start of the file to the string data for this item. The offset should be to a location in the data section, and the data should be in the format specified by "string_data_item" below. There is no alignment requirement for the offset.
}

func ParseStringIds(dexSource []byte, startPoint uint32, size uint32) (stringIds []StringIdItem) {
	sz := uint32(binary.Size(&StringIdItem{}))
	stringIds = make([]StringIdItem, size, size)
	for i := uint32(0); i < size; i++ {
		if i == 26 {
			fmt.Printf("%x\n", dexSource[startPoint+sz*i:startPoint+sz*(i+1)])
		}
		var item StringIdItem
		err := binary.Read(bytes.NewBuffer(dexSource[startPoint+sz*i:startPoint+sz*(i+1)]), binary.LittleEndian, &item)
		if err != nil {
			log.Fatal("read string ids error")
		}
		stringIds[i] = item
	}
	return
}
