package sliceutil

import "golang.org/x/exp/constraints"

// IterFn iter slice with func
func IterFn[T any](s []T, fn func(index int, item T) bool) []T {
	for i, item := range s {
		if ok := fn(i, item); !ok {
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

// NotContain jungle slice whether not contains item
func NotContain[T comparable](item T, list []T) bool {
	return !Contain(item, list)
}

// Sum sum slice
func Sum[T comparable, K constraints.Integer | constraints.Float](list []T, value func(item T) K) K {
	ret := K(0)
	for _, item := range list {
		ret += value(item)
	}
	return ret
}

// Map applies a  function to each element of slice.
func Map[T, K any](arr []T, fn func(item T) K) []K {
	ret := make([]K, len(arr), len(arr))
	for i, item := range arr {
		ret[i] = fn(item)
	}
	return ret
}

// Filter filter slice
func Filter[T any](arr []T, fn func(item T) bool) []T {
	ret := make([]T, 0, len(arr))
	for _, item := range arr {
		if fn(item) {
			ret = append(ret, item)
		}
	}
	return ret
}

// Reduce reduce slice
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

// GenerateSequences generate sequence
func GenerateSequences[T constraints.Integer](start, end, step T) []T {

	if step == 0 {
		return []T{}
	}

	if step > 0 {
		if end < start {
			return []T{}
		}
		ret := make([]T, 0, int((end-start)/step)+1)
		for i := start; i < end; i += step {
			ret = append(ret, T(i))
		}
		return ret
	}

	if end > start {
		return []T{}
	}
	ret := make([]T, 0, int((end-start)/step)+1)
	for i := start; i > end; i += step {
		ret = append(ret, T(i))
	}
	return ret

}

// GroupBy slice group by to map
func GroupBy[T any, K comparable](slice []T, id func(T) K) map[K][]T {
	ret := make(map[K][]T)
	for _, item := range slice {
		k := id(item)
		ret[k] = append(ret[k], item)
	}
	return ret
}
