package byteutil

import "encoding/binary"

// Uint64BigEndianToByte converts uint64 to byte array in big endian.
func Uint64BigEndianToByte(src uint64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, src)
	return bs
}
