package expr

import "golang.org/x/exp/constraints"

// CondExpr return a or b by cond
func CondExpr[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

// Max return max value
func Max[T constraints.Ordered](a, b T) T {
	return CondExpr(a > b, a, b)
}

// Min return min value
func Min[T constraints.Ordered](a, b T) T {
	return CondExpr(a < b, a, b)
}

// Add two numbers
func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

// CallList is a function list for s
func CallList[T any](s T, fn ...func(item T) T) T {
	for _, f := range fn {
		s = f(s)
	}
	return s
}

// FirstOrDefault return first value or default value
func FirstOrDefault[T any](params []T, defaultValue T) T {
	if len(params) > 0 {
		return params[0]
	}
	return defaultValue
}
