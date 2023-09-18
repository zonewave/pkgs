package hightbit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitMapAll(t *testing.T) {
	bs := []byte{uint8(18)}
	require.True(t, Exist(4, bs))
	require.False(t, Exist(5, bs))
	require.Equal(t, []int{1, 4}, All(bs))
	require.False(t, Exist(1024, bs))
}
