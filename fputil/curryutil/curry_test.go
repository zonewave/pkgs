package curryutil

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCurryFn(t *testing.T) {

	addCurry1 := Fn[int, int, int](func(a, b int) int {
		return a + b
	}).New(1)
	require.Equal(t, 4, addCurry1(3))
	require.Equal(t, 6, addCurry1(5))

}

func TestCurryFn3(t *testing.T) {
	add := func(a, b, c int) int {
		return a + b + c
	}
	addCurry1 := Fn3[int, int, int, int](add).New(1)
	addCurry2 := addCurry1(2)
	require.Equal(t, 6, addCurry2(3))
	require.Equal(t, 8, addCurry2(5))

}
