package items

/*
https://source.android.com/devices/tech/dalvik/dex-format#call-site-id-item
alignment: 4 bytes
*/
type CallSiteIdItem struct {
	/*
		offset from the start of the file to call site definition. The offset should be in the data section, and the
		data there should be in the format specified by "call_site_item" below.
	*/
	CallSiteOff uint32
}

/*
https://source.android.com/devices/tech/dalvik/dex-format#call-site-item
alignment: none (byte aligned)
encoded_array_item
*/
type CallSiteItem struct {
}
