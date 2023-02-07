package util

func BitMapGet(index int, data []byte) bool {
	if index >= (len(data) << 3) {
		return false
	}
	div := index >> 3
	mod := index & 0x07
	return ((data[div] >> mod) & 1) == 1
}

func BitMapAll(data []byte) []int {
	indexes := make([]int, 0, len(data)*2)
	for idx, b := range data {
		for i := 0; i < 8; i++ {
			if b&(1<<i) == (1 << i) {
				indexes = append(indexes, idx*8+i)
			}
		}

	}

	return indexes
}
