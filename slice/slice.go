package slice

type Slices[T any] []T

func (s Slices[T]) IterCollect(fn func(item T) (bool, bool)) Slices[T] {
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

func (s Slices[T]) IterFn(fn func(item T) bool) Slices[T] {
	for _, item := range s {
		if ok := fn(item); !ok {
			break
		}
	}
	return s
}

func Mapper[T1 any, T2 any](slice Slices[T1], mapper func(T1) T2) Slices[T2] {
	mapped := make([]T2, 0, len(slice))
	for _, item := range slice {
		mapped = append(mapped, mapper(item))
	}
	return mapped
}
