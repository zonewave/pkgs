package curryutil

// FnVar is a function with Variadic type that can be curried.
type FnVar[T, T1 any] func(...T1) T

func (cf FnVar[T, T1]) New(val T1) func(...T1) T {
	return func(val2 ...T1) T {
		return cf(append([]T1{val}, val2...)...)
	}
}

// FnVar1 is a function with Variadic type that can be curried.
type FnVar1[T, T1, T2 any] func(T1, ...T2) T

// New returns a curried function
func (cf FnVar1[T, T1, T2]) New(val T1) func(...T2) T {
	return func(val2 ...T2) T {
		return cf(val, val2...)
	}
}

// FnVar2 is a function with Variadic type that can be curried.
type FnVar2[T, T1, T2, T3 any] func(T1, T2, ...T3) T

// New returns a curried function
func (cf FnVar2[T, T1, T2, T3]) New(val T1) func(T2) func(...T3) T {
	return func(val2 T2) func(...T3) T {
		return func(val3 ...T3) T {
			return cf(val, val2, val3...)
		}
	}
}

// FnVar3 is a function with Variadic type that can be curried.
type FnVar3[T, T1, T2, T3, T4 any] func(T1, T2, T3, ...T4) T

// New returns a curried function
func (cf FnVar3[T, T1, T2, T3, T4]) New(val T1) func(T2) func(T3) func(...T4) T {
	return func(val2 T2) func(T3) func(...T4) T {
		return func(val3 T3) func(...T4) T {
			return func(val4 ...T4) T {
				return cf(val, val2, val3, val4...)
			}
		}
	}
}

// FnVar4 is a function with Variadic type that can be curried.
type FnVar4[T, T1, T2, T3, T4, T5 any] func(T1, T2, T3, T4, ...T5) T

// New returns a curried function
func (cf FnVar4[T, T1, T2, T3, T4, T5]) New(val T1) func(T2) func(T3) func(T4) func(...T5) T {
	return func(val2 T2) func(T3) func(T4) func(...T5) T {
		return func(val3 T3) func(T4) func(...T5) T {
			return func(val4 T4) func(...T5) T {
				return func(val5 ...T5) T {
					return cf(val, val2, val3, val4, val5...)
				}
			}
		}
	}
}

// FnVar5 is a function with Variadic type that can be curried.
type FnVar5[T, T1, T2, T3, T4, T5, T6 any] func(T1, T2, T3, T4, T5, ...T6) T

// New returns a curried function
func (cf FnVar5[T, T1, T2, T3, T4, T5, T6]) New(val T1) func(T2) func(T3) func(T4) func(T5) func(...T6) T {
	return func(val2 T2) func(T3) func(T4) func(T5) func(...T6) T {
		return func(val3 T3) func(T4) func(T5) func(...T6) T {
			return func(val4 T4) func(T5) func(...T6) T {
				return func(val5 T5) func(...T6) T {
					return func(val6 ...T6) T {
						return cf(val, val2, val3, val4, val5, val6...)
					}
				}
			}
		}
	}
}
