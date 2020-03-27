package items

type TypeItem uint16 // index into the type_ids list

type TypeListItem struct {
	Size uint32     // size of the list, in entries
	List []TypeItem // elements of the list
}
