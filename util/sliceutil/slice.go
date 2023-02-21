package sliceutil

func IterCollect[T any](s []T, fn func(item T) (bool, bool)) []T {
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
