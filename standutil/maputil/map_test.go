package maputil

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func mapInts(total int) map[int]int {
	m := make(map[int]int)
	for j := 1; j <= total; j++ {
		m[j] = j
	}
	return m
}

func TestForEach(t *testing.T) {
	data := mapInts(2)
	total := 0
	ForEach(data, func(key int, value int) {
		if key == 2 {
			return
		}
		total += value
	})
	require.Equal(t, 1, total)
}

func TestMap(t *testing.T) {
	data := mapInts(2)

	ret := Map(data, func(key, value int) string {
		return strconv.Itoa(value)
	})
	require.Equal(t, map[int]string{
		1: "1",
		2: "2",
	}, ret)

}

func TestTransform(t *testing.T) {
	data := mapInts(2)
	ret := Transform(data, func(key, value int) (string, string) {
		return strconv.Itoa(key), strconv.Itoa(value)
	})
	require.Equal(t, map[string]string{
		"1": "1",
		"2": "2",
	}, ret)
}
func TestSum(t *testing.T) {
	data := mapInts(6)
	require.Equal(t, 21, Sum(data))
}
