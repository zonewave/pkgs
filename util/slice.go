package util

// SliceContain jungle slice whether contains item
func SliceContain[T comparable](item T, list []T) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
}
