package reader

import (
	"bytes"
	"encoding/binary"
	"github.com/smm-goddess/dexparser/dex/items"
)

func ReadTypeList(dexBytes []byte, start uint32) items.TypeListItem {
	var size uint32
	_ = binary.Read(bytes.NewBuffer(dexBytes[start:start+4]), binary.LittleEndian, &size)
	list := make([]items.TypeItem, size, size)
	for i := uint32(0); i < size; i++ {
		var typeIdx items.TypeItem
		_ = binary.Read(bytes.NewBuffer(dexBytes[start+4+2*i:start+4+2*(i+1)]), binary.LittleEndian, &typeIdx)
		list[i] = typeIdx
	}
	return items.TypeListItem{
		Size: size,
		List: list,
	}
}
