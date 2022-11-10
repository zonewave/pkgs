package maputil

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMap(t *testing.T) {

	data := map[int]int{
		1: 1,
		2: 2,
	}
	total := 0
	fn := func(key int, value int) {
		total += value
	}
	ForEach(data, fn)
	require.Equal(t, 3, total)

	total = 0
	ForEach(data, func(key int, value int) {
		if key == 2 {
			return
		}
		total += value
	})
	require.Equal(t, 1, total)
}
