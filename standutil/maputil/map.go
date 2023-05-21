package maputil

import "golang.org/x/exp/constraints"

// ForEach iterates over a map and calls the given function for each key/value pair.
func ForEach[K comparable, V any](data map[K]V, fn func(key K, value V)) {
	for k, v := range data {
		fn(k, v)
	}
	return
}

// Map applies a function to each key/value pair of a map and returns a new map with the results.
func Map[K comparable, V1, V2 any](data map[K]V1, fn func(key K, value V1) V2) map[K]V2 {
	return Transform(data, func(key K, value V1) (K, V2) {
		return key, fn(key, value)
	})
}

// Transform applies a function to each key/value pair of a map and returns a new map with the results.
func Transform[K1, K2 comparable, V1, V2 any](data map[K1]V1, fn func(key K1, value V1) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(data))
	for k, v := range data {
		k2, v2 := fn(k, v)
		result[k2] = v2
	}
	return result
}

// Sum sums up all values of a map.
func Sum[K comparable, V constraints.Integer | constraints.Float](data map[K]V) V {
	ret := V(0)
	ForEach(data, func(key K, value V) {
		ret += value
	})
	return ret
}
