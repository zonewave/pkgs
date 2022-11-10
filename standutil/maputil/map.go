package maputil

// ForEach iterates over a map and calls the given function for each key/value pair.
func ForEach[K comparable, V any](data map[K]V, fn func(key K, value V)) {
	for k, v := range data {
		fn(k, v)
	}
	return
}
