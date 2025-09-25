package basic

import "github.com/puresnr/go/algo"

// Basic is a constraint that permits any of Go's basic types.
// This is used to ensure that Clone is not used with reference types
// where it would only perform a shallow copy.
type Basic interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string | ~bool | ~complex64 | ~complex128
}

// Deepcopy creates a deep copy of a slice of basic data types.
func Deepcopy[T Basic](s []T) []T {
	if algo.Empty_slice(s) {
		return nil
	}

	// Create a new slice and copy the elements from the old slice.
	newSlice := make([]T, len(s))
	copy(newSlice, s)
	return newSlice
}
