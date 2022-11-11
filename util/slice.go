package util

func ItemInSlice[T comparable](item T, list []T) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
}
