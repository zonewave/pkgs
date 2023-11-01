package fputil

func Compose[T1, T2, T3 any](fn1 func(T1) T2, fn2 func(T2) T3) func(T1) T3 {
	return func(t T1) T3 {
		return fn2(fn1(t))
	}
}

func Compose3[T1, T2, T3, T4 any](fn1 func(T1) T2, fn2 func(T2) T3, fn3 func(T3) T4) func(T1) T4 {
	return func(t T1) T4 {
		return fn3(fn2(fn1(t)))
	}
}

func Compose4[T1, T2, T3, T4, T5 any](fn1 func(T1) T2, fn2 func(T2) T3, fn3 func(T3) T4, fn4 func(T4) T5) func(T1) T5 {
	return func(t T1) T5 {
		return fn4(fn3(fn2(fn1(t))))
	}
}
