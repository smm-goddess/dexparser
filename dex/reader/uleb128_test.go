package reader

import (
	"fmt"
	"testing"
)

func TestReadUnsignedLeb128(t *testing.T) {
	source := []byte{0x01}
	fmt.Println(ReadUnsignedLeb128(source))
	source = []byte{0x24}
	fmt.Println(ReadUnsignedLeb128(source))
	source = []byte{0xf7, 0x03, 0x7b}
	fmt.Println(ReadUnsignedLeb128(source))
}
