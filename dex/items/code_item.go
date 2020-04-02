package items

import (
	"bytes"
	"encoding/binary"
)

/*
https://source.android.com/devices/tech/dalvik/dex-format#code-item
*/
type CodeItem struct {
	/*
		the number of registers used by this code
	*/
	RegistersSize uint16
	/*
		the number of words of incoming arguments to the method that this code is for
	*/
	InsSize uint16
	/*
		the number of words of outgoing argument space required by this code for method invocation
	*/
	OutsSize uint16
	/*
		the number of try_items for this instance. If non-zero,
		then these appear as the tries array just after the insns in this instance.
	*/
	TriesSize uint16
	/*
		offset from the start of the file to the debug info (line numbers + local variable info) sequence for this code
		, or 0 if there simply is no information. The offset, if non-zero, should be to a location in the data section.
		The format of the data is specified by "debug_info_item" below.
	*/
	DebugInfoOff uint32
	/*
		size of the instructions list, in 16-bit code units
	*/
	InsnsSize uint32
	/*
		actual array of bytecode. The format of code in an insns array is specified by the companion document
		Dalvik bytecode. Note that though this is defined as an array of ushort, there are some internal structures
		that prefer four-byte alignment. Also, if this happens to be in an endian-swapped file, then the swapping is
		only done on individual ushorts and not on the larger internal structures.
	*/
	Insns []uint16
	/*
		two bytes of padding to make tries four-byte aligned. This element is only present if tries_size is non-zero and insns_size is odd.
	*/
	Padding uint16
	/*
		array indicating where in the code exceptions are caught and how to handle them. Elements of the array must be
		non-overlapping in range and in order from low to high address. This element is only present if tries_size is non-zero.
	*/
	Tries []TryItem
	/*
		bytes representing a list of lists of catch types and associated handler addresses. Each try_item has a byte-wise
		offset into this structure. This element is only present if tries_size is non-zero.
	*/
	Handlers []EncodedCatchHandlerList
}

func ParseCodeItem(dexBytes []byte, offset uint32) (codeItem CodeItem) {
	var registerSize, insSize, outsSize, triesSize uint16
	var insnsSize uint32
	_ = binary.Read(bytes.NewReader(dexBytes[offset:]), binary.LittleEndian, &registerSize)
	codeItem.RegistersSize = registerSize
	offset += 2
	_ = binary.Read(bytes.NewReader(dexBytes[offset:]), binary.LittleEndian, &insSize)
	codeItem.InsSize = insSize
	offset += 2
	_ = binary.Read(bytes.NewReader(dexBytes[offset:]), binary.LittleEndian, &outsSize)
	codeItem.OutsSize = outsSize
	offset += 2
	_ = binary.Read(bytes.NewReader(dexBytes[offset:]), binary.LittleEndian, &triesSize)
	codeItem.TriesSize = triesSize
	offset += 2
	_ = binary.Read(bytes.NewReader(dexBytes[offset:]), binary.LittleEndian, &insnsSize)
	codeItem.InsnsSize = insnsSize
	offset += 4
	return
}

type TryItem struct {
}

type EncodedCatchHandlerList struct {
}
