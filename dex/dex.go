package dex

import (
	. "github.com/smm-goddess/dexparser/dex/items"
)

type Dex struct {
	header         HeaderItem
	StringIds      []StringIdItem
	TypeIds        []TypeIdItem
	ProtoIds       []ProtoIdItem
	FieldIds       []FieldIdItem
	MethodIds      []MethodIdItem
	ClassDefs      []ClassDefItem
	CallSiteIds    []CallSiteIdItem
	MethodHandlers []MethodHandlerItem
	Data           []byte
	LinkData       []byte
}
