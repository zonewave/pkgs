// Package hightbit This package is designed for a bitmap with offset counted from the high bits of 8-bit bytes.
package hightbit

// Exist check index exist
func Exist(index int, data []byte) bool {
	if index >= (len(data) << 3) {
		return false
	}
	div := index >> 3
	mod := index & 0x07
	return ((data[div] >> mod) & 1) == 1
}

// All get all index
// Notice: not use method when data is too large
func All(data []byte) []int {
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
