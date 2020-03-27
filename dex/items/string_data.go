package items

import (
	"fmt"
	"github.com/smm-goddess/dexparser/dex/reader"
)

/**
 * alignment: none
 */
type StringDataItem struct {
	/*
	 * size of this string, in UTF-16 code units (which is the "string length" in many systems). That is, this is
	 * the decoded length of the string. (The encoded length is implied by the position of the 0 byte.)
	 */
	Utf16Size uint32
	/*
	 * a series of MUTF-8 code units (a.k.a. octets, a.k.a. bytes) followed by a byte of value 0. See "MUTF-8
	 * (Modified UTF-8) Encoding" above for details and discussion about the data format.
	 */
	Data []byte
}

func (item StringDataItem) String() string {
	return fmt.Sprintf("length:%d,data:%s", item.Utf16Size, string(item.Data))
}

func ReadStringData(dexBytes []byte, off IOff) StringDataItem {
	length, readCount := reader.ReadUnsignedLeb128(dexBytes[off.GetOffset():])
	data := make([]byte, 0)
	for index := off.GetOffset() + uint32(readCount); dexBytes[index] != 0x0; index++ {
		data = append(data, dexBytes[index])
	}
	return StringDataItem{length, data}
}
