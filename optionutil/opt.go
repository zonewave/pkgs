package optionutil

// ApplyOpt apply option for item
func ApplyOpt[T any](item T, opts ...func(T)) T {
	for _, o := range opts {
		o(item)
	}
	return item
}
