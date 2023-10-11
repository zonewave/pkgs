package curryutil

// Fn is a function type that can be curried.
type Fn[T, T1, T2 any] func(T1, T2) T

// New returns a curried function
func (cf Fn[T, T1, T2]) New(val T1) func(T2) T {
	return func(val2 T2) T {
		return cf(val, val2)
	}
}

// Fn3 is a function type that can be curried.
type Fn3[T, T1, T2, T3 any] func(T1, T2, T3) T

// New returns a curried function
func (cf Fn3[T, T1, T2, T3]) New(val T1) func(T2) func(T3) T {
	return func(val2 T2) func(T3) T {
		return func(val3 T3) T {
			return cf(val, val2, val3)
		}
	}
}

// Fn4 is a function type that can be curried.
type Fn4[T, T1, T2, T3, T4 any] func(T1, T2, T3, T4) T

// New returns a curried function
func (cf Fn4[T, T1, T2, T3, T4]) New(val T1) func(T2) func(T3) func(T4) T {
	return func(val2 T2) func(T3) func(T4) T {
		return func(val3 T3) func(T4) T {
			return func(val4 T4) T {
				return cf(val, val2, val3, val4)
			}
		}
	}
}

// Fn5 is a function type that can be curried.
type Fn5[T, T1, T2, T3, T4, T5 any] func(T1, T2, T3, T4, T5) T

// New returns a curried function
func (cf Fn5[T, T1, T2, T3, T4, T5]) New(val T1) func(T2) func(T3) func(T4) func(T5) T {
	return func(val2 T2) func(T3) func(T4) func(T5) T {
		return func(val3 T3) func(T4) func(T5) T {
			return func(val4 T4) func(T5) T {
				return func(val5 T5) T {
					return cf(val, val2, val3, val4, val5)
				}
			}
		}
	}
}

// Fn6 is a function type that can be curried.
type Fn6[T, T1, T2, T3, T4, T5, T6 any] func(T1, T2, T3, T4, T5, T6) T

// New returns a curried function
func (cf Fn6[T, T1, T2, T3, T4, T5, T6]) New(val T1) func(T2) func(T3) func(T4) func(T5) func(T6) T {
	return func(val2 T2) func(T3) func(T4) func(T5) func(T6) T {
		return func(val3 T3) func(T4) func(T5) func(T6) T {
			return func(val4 T4) func(T5) func(T6) T {
				return func(val5 T5) func(T6) T {
					return func(val6 T6) T {
						return cf(val, val2, val3, val4, val5, val6)
					}
				}
			}
		}
	}
}
