package items

/*
 *	alignment 4 byte
 */
type FieldIdItem struct {
	ClassIdx uint16
	TypeIdx  uint16
	NameIdx  uint32
}
