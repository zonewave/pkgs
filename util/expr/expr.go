package expr

import "golang.org/x/exp/constraints"

func condExpr[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	return condExpr(a > b, a, b)
}
func Min[T constraints.Ordered](a, b T) T {
	return condExpr(a < b, a, b)
}

func FnList[T any](s T, fn ...func(item T) T) T {
	for _, f := range fn {
		s = f(s)
	}
	return s
}
