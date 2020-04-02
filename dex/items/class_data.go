package items

import (
	"github.com/smm-goddess/dexparser/dex/reader"
)

/*
https://source.android.com/devices/tech/dalvik/dex-format#class-data-item
*/
type ClassDataItem struct {
	/*
		the number of static fields defined in this item
		uleb128 format
	*/
	StaticFieldSize uint32
	/*
		the number of instance fields defined in this item
		uleb128 format
	*/
	InstanceFieldsSize uint32
	/*
		the number of direct methods defined in this item
		uleb128 format
	*/
	DirectMethodsSize uint32
	/*
		the number of virtual methods defined in this item
		uleb128 format
	*/
	VirtualMethodsSize uint32
	/*
		the defined static fields, represented as a sequence of encoded elements.
		The fields must be sorted by field_idx in increasing order.
	*/
	StaticFields []EncodedField
	/*
		the defined instance fields, represented as a sequence of encoded elements.
		The fields must be sorted by field_idx in increasing order.
	*/
	InstanceFields []EncodedField
	/*
		the defined direct (any of static, private, or constructor) methods, represented as a sequence of encoded elements.
		The methods must be sorted by method_idx in increasing order.
	*/
	DirectMethods []EncodedMethod
	/*
		the defined virtual (none of static, private, or constructor) methods, represented as a sequence of encoded elements.
		This list should not include inherited methods unless overridden by the class that this item represents. The methods
		must be sorted by method_idx in increasing order. The method_idx of a virtual method must not be the same as any direct method.
	*/
	VirtualMethods []EncodedMethod
}

func ParseClassData(dexBytes []byte, offset uint32) (classDataItem ClassDataItem) {
	var readCnt, staticFieldSize, instanceFieldSize, directMethodSize, virtualMethodSize uint32
	staticFieldSize, readCnt = reader.ReadUnsignedLeb128(dexBytes[offset:])
	classDataItem.StaticFieldSize, offset = staticFieldSize, offset+readCnt
	instanceFieldSize, readCnt = reader.ReadUnsignedLeb128(dexBytes[offset:])
	classDataItem.InstanceFieldsSize, offset = instanceFieldSize, offset+readCnt
	directMethodSize, readCnt = reader.ReadUnsignedLeb128(dexBytes[offset:])
	classDataItem.DirectMethodsSize, offset = directMethodSize, offset+readCnt
	virtualMethodSize, readCnt = reader.ReadUnsignedLeb128(dexBytes[offset:])
	classDataItem.VirtualMethodsSize, offset = virtualMethodSize, offset+readCnt

	staticFields := make([]EncodedField, classDataItem.StaticFieldSize, classDataItem.StaticFieldSize)
	for index := uint32(0); index < classDataItem.StaticFieldSize; index++ {
		staticFields[index], readCnt = ReadEncodedField(dexBytes, offset)
		offset += readCnt
	}
	classDataItem.StaticFields = staticFields

	instanceFields := make([]EncodedField, classDataItem.InstanceFieldsSize, classDataItem.InstanceFieldsSize)
	for index := uint32(0); index < classDataItem.InstanceFieldsSize; index++ {
		instanceFields[index], readCnt = ReadEncodedField(dexBytes, offset)
		offset += readCnt
	}
	classDataItem.InstanceFields = instanceFields

	directMethods := make([]EncodedMethod, classDataItem.DirectMethodsSize, classDataItem.DirectMethodsSize)
	for index := uint32(0); index < classDataItem.DirectMethodsSize; index++ {
		directMethods[index], readCnt = ReadEncodedMethod(dexBytes, offset)
		offset += readCnt
	}
	classDataItem.DirectMethods = directMethods

	virtualMethods := make([]EncodedMethod, classDataItem.VirtualMethodsSize, classDataItem.VirtualMethodsSize)
	for index := uint32(0); index < classDataItem.VirtualMethodsSize; index++ {
		virtualMethods[index], readCnt = ReadEncodedMethod(dexBytes, offset)
		offset += readCnt
	}
	classDataItem.VirtualMethods = virtualMethods

	return
}

type EncodedField struct {
	/*
		index into the field_ids list for the identity of this field (includes the name and descriptor), represented as
		a difference from the index of previous element in the list. The index of the first element in a list is represented directly.
		format uleb128
	*/
	FieldIdxDiff uint32
	/*
		access flags for the field (public, final, etc.). See "access_flags Definitions" for details.
		format uleb128
	*/
	AccessFlags uint32
}

func ReadEncodedField(dexBytes []byte, offset uint32) (encodedField EncodedField, readCount uint32) {
	var fieldIdxDiff, accessFlags, readCnt uint32
	fieldIdxDiff, readCnt = reader.ReadUnsignedLeb128(dexBytes[offset:])
	encodedField.FieldIdxDiff = fieldIdxDiff
	readCount += readCnt
	accessFlags, readCnt = reader.ReadUnsignedLeb128(dexBytes[offset+readCount:])
	encodedField.AccessFlags = accessFlags
	readCount += readCnt
	return
}

type EncodedMethod struct {
	/*
		index into the method_ids list for the identity of this method (includes the name and descriptor), represented
		as a difference from the index of previous element in the list. The index of the first element in a list is represented directly.
		format uleb128
	*/
	MethodIdxDiff uint32
	/*
		access flags for the method (public, final, etc.). See "access_flags Definitions" for details.
		format uleb128
	*/
	AccessFlags uint32
	/*
		offset from the start of the file to the code structure for this method, or 0 if this method is either abstract
		or native. The offset should be to a location in the data section. The format of the data is specified by "code_item" below.
		format uleb128
	*/
	CodeOff uint32
}

func ReadEncodedMethod(dexBytes []byte, offset uint32) (encodedMethod EncodedMethod, readCount uint32) {
	var methodIdxDiff, accessFlags, codeOff, readCnt uint32
	methodIdxDiff, readCnt = reader.ReadUnsignedLeb128(dexBytes[offset:])
	encodedMethod.MethodIdxDiff = methodIdxDiff
	readCount += readCnt
	accessFlags, readCnt = reader.ReadUnsignedLeb128(dexBytes[offset+readCount:])
	encodedMethod.AccessFlags = accessFlags
	readCount += readCnt
	codeOff, readCnt = reader.ReadUnsignedLeb128(dexBytes[offset+readCount:])
	encodedMethod.CodeOff = codeOff
	readCount += readCnt
	return
}
