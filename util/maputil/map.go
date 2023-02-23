package maputil

func IterFn[K comparable, V any](data map[K]V, fn func(key K, value V) bool) {
	for k, v := range data {
		if ok := fn(k, v); !ok {
			break
		}
	}
	return
}
