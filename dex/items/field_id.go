package items

/*
 *	alignment 4 byte
 */
type FieldIdItem struct {
	/*
		index into the type_ids list for the definer of this field. This must be a class type,
			and not an array or primitive type.
	*/
	ClassIdx uint16
	/*
		index into the type_ids list for the type of this field
	*/
	TypeIdx uint16
	/*
		index into the string_ids list for the name of this field. The string must conform to the syntax for MemberName,
			defined above.
	*/
	NameIdx uint32
}
