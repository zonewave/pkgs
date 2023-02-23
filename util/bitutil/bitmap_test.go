package bitutil

import (
	"fmt"
	"testing"
)

func TestBitMapAll(t *testing.T) {
	bs := []byte{uint8(1)}
	fmt.Println(BitMapAll(bs))
	fmt.Println(BitMapGet(0, bs))
}
