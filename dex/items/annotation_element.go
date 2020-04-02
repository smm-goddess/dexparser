package items

import (
	"bytes"
	"encoding/binary"
)

type AnnotationElement struct {
	/*
		element name, represented as an index into the string_ids section. The string must conform to the syntax for MemberName, defined above.
	*/
	NameIdx uint32
	/*
		element value
	*/
	Value EncodedValue
}

func ReadAnnotationElement(dexBytes []byte, offSet uint32) (annotationElement AnnotationElement, readCount uint32) {
	nameIdxBytes := make([]byte, 4, 4)
	copy(nameIdxBytes, dexBytes[offSet:])
	var nameIdx uint32
	_ = binary.Read(bytes.NewBuffer(nameIdxBytes), binary.LittleEndian, &nameIdx)
	annotationElement.NameIdx = nameIdx
	readCount += 4
	value, cnt := ReadEncodedValue(dexBytes, offSet+4)
	annotationElement.Value = value
	readCount += cnt
	return
}
