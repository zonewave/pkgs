package curryutil

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFnVar(t *testing.T) {
	addCurry1 := FnVar[int, int](mathutil.Sum[int]).New(1)
	require.Equal(t, 4, addCurry1(3))
	require.Equal(t, 12, addCurry1(5, 6))

	addCurry2 := FnVar[int, int](addCurry1).New(5)
	require.Equal(t, 12, addCurry2(6))
	require.Equal(t, 19, addCurry2(7, 6))

}

func TestFnVar1(t *testing.T) {
	add := func(a, b, c int, d ...int) int {
		return a + b + c
	}
	addCurry1 := FnVar3[int, int, int, int, int](add).New(1)
	addCurry2 := addCurry1(2)
	require.Equal(t, 6, addCurry2(3)())
	require.Equal(t, 8, addCurry2(5)())

}
