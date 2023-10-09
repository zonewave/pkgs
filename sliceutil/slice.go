package sliceutil

import (
	"github.com/zonewave/pkgs/optionutil"
	"golang.org/x/exp/constraints"
)

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

// GenerateSliceOption generate slice option
type GenerateSliceOption[T any] struct {
	ApplySliceItem func(int, T) T
}

// NewGenerateSliceOption new generate slice option
func NewGenerateSliceOption[T any](opts ...func(*GenerateSliceOption[T])) *GenerateSliceOption[T] {
	option := &GenerateSliceOption[T]{}
	return optionutil.ApplyOpt(option, opts...)
}

// GenerateSlice generate slice
func GenerateSlice[T any](total int, generate func() T, opts ...func(*GenerateSliceOption[T])) []T {
	ret := make([]T, 0, total)
	option := NewGenerateSliceOption[T](opts...)
	for i := 0; i < total; i++ {
		item := generate()
		if option.ApplySliceItem != nil {
			item = option.ApplySliceItem(i, item)
		}
		ret = append(ret, item)
	}
	return ret
}
