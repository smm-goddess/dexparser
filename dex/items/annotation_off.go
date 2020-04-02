package items

type AnnotationOffItem struct {
	/*
		offset from the start of the file to an annotation. The offset should be to a location in the data section,
		and the format of the data at that location is specified by "annotation_item" below.
	*/
	AnnotationOff uint32
}
