package sliceutil

import "golang.org/x/exp/constraints"

// Sequences generate sequence
func Sequences[T constraints.Integer](start, end, step T) []T {
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

// GenerateSlice generate slice
func GenerateSlice[T any](total int, generate func(int) T) []T {
	ret := make([]T, 0, total)
	for i := 0; i < total; i++ {
		ret = append(ret, generate(i))
	}
	return ret
}
