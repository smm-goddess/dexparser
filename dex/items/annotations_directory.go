package items

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

type AnnotationOffItem struct {
	/*
		offset from the start of the file to an annotation. The offset should be to a location in the data section,
		and the format of the data at that location is specified by "annotation_item" below.
	*/
	AnnotationOff uint32
}

type AnnotationItem struct {
	/*
		intended visibility of this annotation (see below)
	*/
	Visibility uint8
	/*
		encoded annotation contents, in the format described by "encoded_annotation format" under "encoded_value encoding" above.
	*/
	Annotation EncodedAnnotation
}
