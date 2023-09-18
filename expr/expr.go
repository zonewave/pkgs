package expr

import "golang.org/x/exp/constraints"

// Add two numbers
func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

// FirstOrDefault return first value or default value
func FirstOrDefault[T any](params []T, defaultValue T) T {
	if len(params) > 0 {
		return params[0]
	}
	return defaultValue
}
