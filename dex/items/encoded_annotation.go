package items

import (
	"github.com/smm-goddess/dexparser/dex/reader"
)

type EncodedAnnotation struct {
	/*
		type of the annotation. This must be a class (not array or primitive) type.
	*/
	TypeIdx uint32
	/*
		number of name-value mappings in this annotation
	*/
	Size uint32
	/*
		elements of the annotation, represented directly in-line (not as offsets). Elements must be sorted in increasing order by string_id index.
	*/
	Elements []AnnotationElement
}

func ReadEncodeAnnotation(dexBytes []byte, offSet uint32) (encodedAnnotation EncodedAnnotation, readCount uint32) {
	var typeIdx, size, cnt uint32
	typeIdx, cnt = reader.ReadUnsignedLeb128(dexBytes[offSet:])
	encodedAnnotation.TypeIdx = typeIdx
	readCount += cnt
	size, cnt = reader.ReadUnsignedLeb128(dexBytes[offSet+cnt:])
	encodedAnnotation.Size = size
	readCount += cnt
	elements := make([]AnnotationElement, size, size)
	for i := uint32(0); i < size; i++ {
		elements[i], cnt = ReadAnnotationElement(dexBytes, offSet+readCount)
		readCount += cnt
	}
	encodedAnnotation.Elements = elements
	return
}
