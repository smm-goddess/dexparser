package reader

func ReadUnsignedLeb128(source []byte) (result uint32, readCount int) {
	result, readCount = uint32(source[0]), 1
	if result > 0x7f {
		cur := uint32(source[1])
		result = (result & 0x7f) | ((cur & 0x7f) << 7)
		readCount = 2
		if cur > 0x7f {
			cur = uint32(source[2])
			result |= (cur & 0x7f) << 14
			readCount = 3
			if cur > 0x7f {
				cur = uint32(source[3])
				result |= (cur & 0x7f) << 21
				readCount = 4
				if cur > 0x7f {
					cur = uint32(source[4])
					result |= cur << 28
					readCount = 5
				}
			}
		}
	}
	return
}
