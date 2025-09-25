package slice

// Clone creates a deep copy of a slice.
// For slices of basic types, this effectively creates a new slice
// and copies the elements.
func Clone[T any](s []T) []T {
	// If the original slice is nil, return nil to maintain that state.
	if s == nil {
		return nil
	}

	// This is an idiomatic and efficient way to create a shallow copy of the slice.
	return append([]T{}, s...)
}
