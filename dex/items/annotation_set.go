package items

import (
	"bytes"
	"encoding/binary"
)

type AnnotationSetItem struct {
	/*
		size of the set, in entries
	*/
	Size uint32
	/*
		elements of the set. The elements must be sorted in increasing order, by type_idx.
	*/
	Entries []AnnotationOffItem
}

func ReadAnnotationSetItem(dexBytes []byte, offSet uint32) (annotationSetItem AnnotationSetItem) {
	var size uint32
	_ = binary.Read(bytes.NewBuffer(dexBytes[offSet:]), binary.LittleEndian, &size)
	annotationSetItem.Size = size
	entries := make([]AnnotationOffItem, annotationSetItem.Size, annotationSetItem.Size)
	_ = binary.Read(bytes.NewReader(dexBytes[offSet+4:]), binary.LittleEndian, &entries)
	annotationSetItem.Entries = entries
	return
}
