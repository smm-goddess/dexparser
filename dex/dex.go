package dex

import (
	"bytes"
	"fmt"
	. "github.com/smm-goddess/dexparser/dex/items"
	"github.com/smm-goddess/dexparser/dex/reader"
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

func ParseDexFile(dexBytes []byte) {
	header := ParseHeader(dexBytes)
	stringIds := ParseStringIds(dexBytes, header.StringIdsOff, header.StringIdsSize)
	typeIds := ParseTypeIdIds(dexBytes, header.TypeIdsOff, header.TypeIdsSize)
	protoIds := ParseProtoIds(dexBytes, header.ProtoIdsOff, header.ProtoIdsSize)

	/**
	parse protoId
	*/
	protoId := protoIds[692]
	returnTypeIdx := stringIds[typeIds[protoId.ReturnTypeIdx].DescriptorIdx]
	returnTypeString := ReadStringData(dexBytes, returnTypeIdx)
	fmt.Println(Descriptor2Class(returnTypeString.Data))

	shortyIdx := stringIds[protoId.ShortyIdx]
	shortyStr := ReadStringData(dexBytes, shortyIdx)
	fmt.Println(string(shortyStr.Data))

	paramBuffer := bytes.Buffer{}
	paramBuffer.WriteByte('(')
	if protoId.ParametersOff > 0 {
		for _, i := range reader.ReadTypeList(dexBytes, protoId.ParametersOff).List {
			item := ReadStringData(dexBytes, stringIds[typeIds[i].DescriptorIdx])
			paramBuffer.WriteString(Descriptor2Class(item.Data))
			paramBuffer.WriteByte(',')
		}
		paramBuffer.Truncate(paramBuffer.Len() - 1)
	}
	paramBuffer.WriteByte(')')
	fmt.Printf("%s%s\n", Descriptor2Class(returnTypeString.Data), paramBuffer.String())
}
