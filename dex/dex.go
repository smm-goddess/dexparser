package dex

import (
	"bytes"
	"fmt"
	"github.com/smm-goddess/dexparser/dex/consts"
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

func getTypeStringBasedOnTypeIdIndex(dexBytes []byte, stringIds []StringIdItem, typeIds []TypeIdItem, typeIdIndex uint32) string {
	return Descriptor2Class(ReadStringData(dexBytes, stringIds[typeIds[typeIdIndex].DescriptorIdx]).Data)
}

func getStringBasedOnStringId(dexBytes []byte, stringIds []StringIdItem, stringIndex uint32) string {
	return string(ReadStringData(dexBytes, stringIds[stringIndex]).Data)
}

func ParseDexFile(dexBytes []byte) {
	header := ParseHeader(dexBytes)
	stringIds := ParseStringIds(dexBytes, header.StringIdsOff, header.StringIdsSize)
	typeIds := ParseTypeIdIds(dexBytes, header.TypeIdsOff, header.TypeIdsSize)
	protoIds := ParseProtoIds(dexBytes, header.ProtoIdsOff, header.ProtoIdsSize)

	/**
	**parse protoId
	 **/
	fmt.Println("---------- Parse ProtoId -----------")
	protoId := protoIds[5]
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
	fmt.Println("---------- Parse ProtoId End -----------\n")
	/*
	**	parse field ids
	 **/
	fmt.Println("---------- Parse FieldId -----------")
	fieldIds, _ := ParseFieldIds(dexBytes, header.FieldIdsOff, header.FieldIdsSize)
	fieldId := fieldIds[0]
	fmt.Println(Descriptor2Class(ReadStringData(dexBytes, stringIds[typeIds[fieldId.ClassIdx].DescriptorIdx]).Data))
	fmt.Println(Descriptor2Class(ReadStringData(dexBytes, stringIds[typeIds[fieldId.TypeIdx].DescriptorIdx]).Data))
	fmt.Println(string(ReadStringData(dexBytes, stringIds[fieldId.NameIdx]).Data))
	fmt.Println("---------- Parse FieldId End-----------\n")

	/*
		parse method ids
	*/
	fmt.Println("---------- Parse MethodId -----------")
	methodIds, _ := ParseMethodIds(dexBytes, header.MethodIdsOff, header.MethodIdsSize)
	methodId := methodIds[5]
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
	fmt.Println("---------- Parse MethodId End-----------\n")

	/*
		class def
	*/
	fmt.Println("---------- Parse Class Def -----------")
	classDefs, _ := ParseClassDefs(dexBytes, header.ClassDefsOff, header.ClassDefsSize)
	fmt.Println(len(classDefs))
	classDef := classDefs[1]
	fmt.Println(getTypeStringBasedOnTypeIdIndex(dexBytes, stringIds, typeIds, classDef.ClassIdx))
	fmt.Println(consts.GetAccessFlagsString(classDef.AccessFlags))
	if classDef.SuperClassIdx == consts.NO_INDEX {
		// it's a root class such as java.lang.Object
	} else {
		fmt.Println("superClass:" + getTypeStringBasedOnTypeIdIndex(dexBytes, stringIds, typeIds, classDef.SuperClassIdx))
	}
	if classDef.InterfaceOff == 0 {
		// no interface implements
	} else {
		interfaceItems := ReadTypeList(dexBytes, classDef.InterfaceOff)
		for _, item := range interfaceItems.List {
			fmt.Println("interface:" + getTypeStringBasedOnTypeIdIndex(dexBytes, stringIds, typeIds, uint32(item)))
		}
	}
	if classDef.SourceFileIdx == consts.NO_INDEX {
		// no source file
	} else {
		fmt.Println("source file:" + getStringBasedOnStringId(dexBytes, stringIds, classDef.SourceFileIdx))
	}

	if classDef.AnnotationOff == 0 {
		// no annotations on this class
	} else {

	}

}
