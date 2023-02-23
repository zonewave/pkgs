package sliceutil

import "golang.org/x/exp/constraints"

func IterCollect[T any](s []T, fn func(item T) (cn bool, exit bool)) []T {
	ret := make([]T, 0, len(s))
	for _, item := range s {
		if ok, et := fn(item); ok {
			ret = append(ret, item)
		} else if et {
			break
		}
	}
	return ret
}

func IterFn[T any](s []T, fn func(item T) bool) []T {
	for _, item := range s {
		if ok := fn(item); !ok {
			break
		}
	}
	return s
}

// Contain jungle slice whether contains item
func Contain[T comparable](item T, list []T) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
}

func NotContain[T comparable](item T, list []T) bool {
	return !Contain(item, list)
}

func Sum[T comparable, K constraints.Ordered](list []T, value func(item T) K) K {
	ret := K(0)
	for _, item := range list {
		ret += value(item)
	}
	return ret
}

func Reduce[T any](arr []T, fn func(curItem, preCum T) T) []T {

	if len(arr) == 0 {
		return nil
	}
	cum := make([]T, len(arr), len(arr))
	cum[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		cum[i] = fn(arr[i], cum[i-1])
	}
	return cum
}

func GenerateSequences[T constraints.Integer](start, end, step int) []T {
	if start > end {
		return []T{}
	}
	if step <= 0 {
		return []T{}
	}
	ret := make([]T, 0, (end-start)/step)
	for i := start; i < end; i += step {
		ret = append(ret, T(i))
	}
	return ret
}
