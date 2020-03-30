package dex

import (
	"bytes"
	"fmt"
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

func ParseDexFile(dexBytes []byte) {
	header := ParseHeader(dexBytes)
	stringIds := ParseStringIds(dexBytes, header.StringIdsOff, header.StringIdsSize)
	typeIds := ParseTypeIdIds(dexBytes, header.TypeIdsOff, header.TypeIdsSize)
	protoIds := ParseProtoIds(dexBytes, header.ProtoIdsOff, header.ProtoIdsSize)

	/**
	**parse protoId
	 **/
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
		for _, i := range ReadTypeList(dexBytes, protoId.ParametersOff).List {
			item := ReadStringData(dexBytes, stringIds[typeIds[i].DescriptorIdx])
			paramBuffer.WriteString(Descriptor2Class(item.Data))
			paramBuffer.WriteByte(',')
		}
		paramBuffer.Truncate(paramBuffer.Len() - 1)
	}
	paramBuffer.WriteByte(')')
	fmt.Printf("%s%s\n", Descriptor2Class(returnTypeString.Data), paramBuffer.String())

	/*
	**	parse field ids
	 **/
	fieldIds, _ := ParseFieldIds(dexBytes, header.FieldIdsOff, header.FieldIdsSize)
	fieldId := fieldIds[13060]
	fmt.Println(Descriptor2Class(ReadStringData(dexBytes, stringIds[typeIds[fieldId.ClassIdx].DescriptorIdx]).Data))
	fmt.Println(Descriptor2Class(ReadStringData(dexBytes, stringIds[typeIds[fieldId.TypeIdx].DescriptorIdx]).Data))
	fmt.Println(string(ReadStringData(dexBytes, stringIds[fieldId.NameIdx]).Data))
	/*
		parse method ids
	*/
	fmt.Println("---------- Parse MethodId -----------")
	methodIds, _ := ParseMethodIds(dexBytes, header.MethodIdsOff, header.MethodIdsSize)
	methodId := methodIds[20]
	// ClassIndex
	class := Descriptor2Class(ReadStringData(dexBytes, stringIds[typeIds[methodId.ClassIdx].DescriptorIdx]).Data)
	// ProtoIndex
	protoId = protoIds[methodId.ProtoIdx]
	returnTypeIdx = stringIds[typeIds[protoId.ReturnTypeIdx].DescriptorIdx]
	returnTypeString = ReadStringData(dexBytes, returnTypeIdx)
	shortyIdx = stringIds[protoId.ShortyIdx]
	shortyStr = ReadStringData(dexBytes, shortyIdx)
	paramBuffer = bytes.Buffer{}
	paramBuffer.WriteByte('(')
	if protoId.ParametersOff > 0 {
		for _, i := range ReadTypeList(dexBytes, protoId.ParametersOff).List {
			item := ReadStringData(dexBytes, stringIds[typeIds[i].DescriptorIdx])
			paramBuffer.WriteString(Descriptor2Class(item.Data))
			paramBuffer.WriteByte(',')
		}
		paramBuffer.Truncate(paramBuffer.Len() - 1)
	}
	paramBuffer.WriteByte(')')
	//
	name := string(ReadStringData(dexBytes, stringIds[methodId.NameIdx]).Data)
	fmt.Printf("%s %s.%s %s\n", Descriptor2Class(returnTypeString.Data), class, name, paramBuffer.String())

	/*
		class def
	*/
	fmt.Println("---------- Parse Class Def -----------")


}
