package items

import (
	"bytes"
	"encoding/binary"
)

type AnnotationsDirectoryItem struct {
	/*
		offset from the start of the file to the annotations made directly on the class, or 0 if the class has
		no direct annotations. The offset, if non-zero, should be to a location in the data section. The format
		of the data is specified by "annotation_set_item" below.
	*/
	ClassAnnotationsOff uint32
	/*
		count of fields annotated by this item
	*/
	FieldsSize uint32
	/*
		count of methods annotated by this item
	*/
	AnnotatedMethodsSize uint32
	/*
		count of method parameter lists annotated by this item
	*/
	AnnotatedParametersSize uint32
	FieldAnnotations        []FieldAnnotation
	MethodAnnotations       []MethodAnnotation
	ParameterAnnotation     []ParameterAnnotation
}

func ReadAnnotationDirectory(dexBytes []byte, offSet uint32) AnnotationsDirectoryItem {
	var classAnnotationOff, fieldsSize, annotationMethodSize, annotatedParametersSize uint32
	_ = binary.Read(bytes.NewBuffer(dexBytes[offSet:]), binary.LittleEndian, &classAnnotationOff)
	_ = binary.Read(bytes.NewBuffer(dexBytes[offSet+4:]), binary.LittleEndian, &fieldsSize)
	_ = binary.Read(bytes.NewBuffer(dexBytes[offSet+8:]), binary.LittleEndian, &annotationMethodSize)
	_ = binary.Read(bytes.NewBuffer(dexBytes[offSet+12:]), binary.LittleEndian, &annotatedParametersSize)
	return AnnotationsDirectoryItem{
		ClassAnnotationsOff:     classAnnotationOff,
		FieldsSize:              fieldsSize,
		AnnotatedMethodsSize:    annotationMethodSize,
		AnnotatedParametersSize: annotatedParametersSize,
		FieldAnnotations:        nil,
		MethodAnnotations:       nil,
		ParameterAnnotation:     nil,
	}
}

type FieldAnnotation struct {
	/*
		index into the field_ids list for the identity of the field being annotated
	*/
	FieldIdx uint32
	/*
		offset from the start of the file to the list of annotations for the field. The offset should be to a location
		in the data section. The format of the data is specified by "annotation_set_item" below.
	*/
	AnnotationsOff uint32
}

type MethodAnnotation struct {
	/*
		index into the method_ids list for the identity of the method being annotated
	*/
	MethodIdx uint32
	/*
		offset from the start of the file to the list of annotations for the field. The offset should be to a location
		in the data section. The format of the data is specified by "annotation_set_item" below.
	*/
	AnnotationsOff uint32
}
type ParameterAnnotation struct {
	/*
		index into the method_ids list for the identity of the method whose parameters are being annotated
	*/
	ParameterIdx uint32
	/*
		offset from the start of the file to the list of annotations for the field. The offset should be to a location
		in the data section. The format of the data is specified by "annotation_set_item" below.
	*/
	AnnotationsOff uint32
}

type AnnotationSetRefList struct {
	/*
		size of the list, in entries
	*/
	Size uint32
	/*
		elements of the list
	*/
	List []AnnotationSetRefItem
}

type AnnotationSetRefItem struct {
	/*
		offset from the start of the file to the referenced annotation set or 0 if there are no annotations for this element.
		The offset, if non-zero, should be to a location in the data section. The format of the data is specified by "annotation_set_item" below.
	*/
	AnnotationsOff uint32
}
