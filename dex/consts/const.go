package consts

import (
	"strings"
)

const (
	HEADER_SIZE                    = 0X70
	ENDIAN_CONSTANT         uint32 = 0x12345678
	REVERSE_ENDIAN_CONSTANT uint32 = 0x78563412
	NO_INDEX                uint32 = 0xffffffff // == -1 if treated as a signed int
)

const (
	/*
		判断可以用异或，相同为0，不同为1
	*/
	ACC_PUBLIC                = 0x1
	ACC_PRIVATE               = 0x2
	ACC_PROTECTED             = 0x4
	ACC_STATIC                = 0x8
	ACC_FINAL                 = 0x10
	ACC_SYNCHRONIZED          = 0x20
	ACC_VOLATILE              = 0x40
	ACC_BRIDGE                = 0x40
	ACC_TRANSIENT             = 0x80
	ACC_VARARGS               = 0x80
	ACC_NATIVE                = 0x100
	ACC_INTERFACE             = 0x200
	ACC_ABSTRACT              = 0x400
	ACC_STRICT                = 0x800
	ACC_SYNTHETIC             = 0x1000
	ACC_ANNOTATION            = 0x2000
	ACC_ENUM                  = 0x4000
	UNUSED                    = 0x8000
	ACC_CONSTRUCTOR           = 0x10000
	ACC_DECLARED_SYNCHRONIZED = 0x20000

	ACC_CLASS_MASK       = ACC_PUBLIC | ACC_FINAL | ACC_INTERFACE | ACC_ABSTRACT | ACC_SYNTHETIC | ACC_ANNOTATION | ACC_ENUM
	ACC_INNER_CLASS_MASK = ACC_CLASS_MASK | ACC_PRIVATE | ACC_PROTECTED | ACC_STATIC
	ACC_FIELD_MASK       = ACC_PUBLIC | ACC_PRIVATE | ACC_PROTECTED | ACC_STATIC | ACC_FINAL | ACC_VOLATILE | ACC_TRANSIENT | ACC_SYNTHETIC | ACC_ENUM
	ACC_METHOD_MASK      = ACC_PUBLIC | ACC_PRIVATE | ACC_PROTECTED | ACC_STATIC | ACC_FINAL | ACC_SYNCHRONIZED | ACC_BRIDGE | ACC_VARARGS | ACC_NATIVE | ACC_ABSTRACT | ACC_STRICT | ACC_SYNTHETIC | ACC_CONSTRUCTOR | ACC_DECLARED_SYNCHRONIZED
)

var accFlagMap = map[int]string{
	0x1:     "ACC_PUBLIC",
	0x2:     "ACC_PRIVATE",
	0x4:     "ACC_PROTECTED",
	0x8:     "ACC_STATIC",
	0x10:    "ACC_FINAL",
	0x20:    "ACC_SYNCHRONIZED",
	0x40:    "ACC_VOLATILE",  //TODO multiple choice ACC_BRIDGE
	0x80:    "ACC_TRANSIENT", //TODO multiple choice ACC_VARARGS
	0x100:   "ACC_NATIVE",
	0x200:   "ACC_INTERFACE",
	0x400:   "ACC_ABSTRACT",
	0x800:   "ACC_STRICT",
	0x1000:  "ACC_SYNTHETIC",
	0x2000:  "ACC_ANNOTATION",
	0x4000:  "ACC_ENUM",
	0x10000: "ACC_CONSTRUCTOR",
	0x20000: "ACC_DECLARED_SYNCHRONIZED",
}

func isLocBit1(bit uint32, bitLoc uint8) bool {
	return (bit>>(bitLoc-1))&1 == 1
}

func GetAccessFlagsString(flag uint32) string {
	flags := make([]string, 0)
	for i := uint8(0); i < 19; i++ {
		if isLocBit1(flag, i) {
			flags = append(flags, accFlagMap[1<<i])
		}
	}
	if len(flags) > 0 {
		return strings.Join(flags, "|")
	}
	return "UNKNOWN"
}
