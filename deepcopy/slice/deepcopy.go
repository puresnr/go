package slice

import (
	"github.com/puresnr/go/algo"
	"github.com/puresnr/go/deepcopy/constraint"
)

// DeepcopyBasic creates a deep copy of a slice of basic data types.
func DeepcopyBasic[T constraint.NonPointerBasic](s []T) []T {
	if algo.Empty_slice(s) {
		return nil
	}

	// Create a new slice and copy the elements from the old slice.
	newSlice := make([]T, len(s))
	copy(newSlice, s)
	return newSlice
}

// Deepcopy creates a deep copy of a slice whose elements implement the Deepcopyable interface.
func Deepcopy[T constraint.Deepcopyable[T]](s []T) []T {
	if algo.Empty_slice(s) {
		return nil
	}

	newSlice := make([]T, len(s))
	for i, v := range s {
		newSlice[i] = v.Deepcopy()
	}
	return newSlice
}
