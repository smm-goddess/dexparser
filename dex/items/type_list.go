package items

import (
	"bytes"
	"encoding/binary"
)

type TypeItem uint16 // index into the type_ids list

type TypeListItem struct {
	Size uint32     // size of the list, in entries
	List []TypeItem // elements of the list
}

func ReadTypeList(dexBytes []byte, start uint32) TypeListItem {
	var size uint32
	_ = binary.Read(bytes.NewBuffer(dexBytes[start:start+4]), binary.LittleEndian, &size)
	list := make([]TypeItem, size, size)
	_ = binary.Read(bytes.NewBuffer(dexBytes[start+4:start+4+2*size]), binary.LittleEndian, &list)
	return TypeListItem{
		Size: size,
		List: list,
	}
}
