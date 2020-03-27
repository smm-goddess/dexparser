package reader

func ParseTypeBasedOnStartAndSize(dexSource []byte, startPoint uint32, size uint32, t interface{}) (typeIds []interface{}) {
	//sz := uint32(binary.Size(t))
	//typeIds = make([]interface{}, size, size)
	//
	//for i := uint32(0); i < size; i++ {
	//	var item t
	//	_ = binary.Read(bytes.NewBuffer(dexSource[startPoint+sz*i:startPoint+sz*(i+1)]), binary.LittleEndian, &item)
	//	typeIds[i] = item
	//}
	return
}
