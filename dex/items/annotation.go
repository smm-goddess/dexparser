package items

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

func ReadAnnotationItem(dexBytes []byte, offSet uint32) (annotationItem AnnotationItem) {
	var visibility = dexBytes[offSet]
	annotationItem.Visibility = visibility
	encodedAnnotation, _ := ReadEncodeAnnotation(dexBytes, offSet+1)
	annotationItem.Annotation = encodedAnnotation
	return
}

const (
	VISIBILITY_BUILD   = 0x00
	VISIBILITY_RUNTIME = 0x01
	VISIBILITY_SYSTEM  = 0x02
)
