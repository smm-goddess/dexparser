package items

/*
https://source.android.com/devices/tech/dalvik/dex-format#encoded-array-item
*/
type EncodedArrayItem struct {
	Value EncodedArray
}

type EncodedArray struct {
	Size   uint32
	Values []EncodedValue
}
