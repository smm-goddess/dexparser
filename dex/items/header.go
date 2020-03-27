package items

import (
	"fmt"
	"github.com/smm-goddess/dexparser/dex/consts"
	"unsafe"
)

type HeaderItem struct {
	Magic         [8]byte // { 0x64 0x65 0x78 0x0a 0x30 0x33 0x39 0x00 } = "dex\n039\0"
	Checksum      uint32
	Signature     [20]byte
	FileSize      uint32
	HeaderSize    uint32
	EndianTag     uint32
	LinkSize      uint32
	LinkOff       uint32
	MapOff        uint32
	StringIdsSize uint32
	StringIdsOff  uint32
	TypeIdsSize   uint32
	TypeIdsOff    uint32
	ProtoIdsSize  uint32
	ProtoIdsOff   uint32
	FieldIdsSize  uint32
	FieldIdsOff   uint32
	MethodIdsSize uint32
	MethodIdsOff  uint32
	ClassDefsSize uint32
	ClassDefsOff  uint32
	DataSize      uint32
	DataOff       uint32
}

func ParseHeader(dexSource []byte) (header *HeaderItem) {
	source := dexSource[:consts.HEADER_SIZE]
	header = *(**HeaderItem)(unsafe.Pointer(&source))
	return
}

func (header *HeaderItem) String() string {
	return fmt.Sprintf(
		`Magic:%c
CheckSum:%xh
Signature:%xh
FileSize:%d
HeaderSize:%d
EndianTag:%xh
LinkSize:%d
LinkOff:%d
MapOff:%d
StringIdsSize:%d
StringIdsOff:%d
TypeIdsSize:%d
TypeIdsOff:%d
ProtoIdsSize:%d
ProtoIdsOff:%d
FieldIdsSize:%d
FieldIdsOff:%d
MethodIdsSize:%d
MethodIdsOff:%d
ClassDefsSize:%d
ClassDefsOff:%d
DataSize:%d
DataOff:%d
`, header.Magic, header.Checksum, header.Signature, header.FileSize, header.HeaderSize, header.EndianTag, header.LinkSize, header.LinkOff,
		header.MapOff, header.StringIdsSize, header.StringIdsOff, header.TypeIdsSize, header.TypeIdsOff, header.ProtoIdsSize, header.ProtoIdsOff,
		header.FieldIdsSize, header.FieldIdsOff, header.MethodIdsSize, header.MethodIdsOff, header.ClassDefsSize, header.ClassDefsOff,
		header.DataSize, header.DataOff)
}
