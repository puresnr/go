package mmap

import (
	"github.com/puresnr/go/algo"
	"github.com/puresnr/go/deepcopy/constraint"
)

// DeepcopyBasic creates a deep copy of a map with comparable keys and basic value types.
// Keys must be comparable. Values must be basic types (e.g., int, string, bool).
func DeepcopyBasic[K interface{ comparable; constraint.NonPointerBasic }, V constraint.NonPointerBasic](m map[K]V) map[K]V {
	// Consistent with slice deepcopy, return nil for nil or empty maps.
	if m == nil || algo.Empty_map(m) { // Assuming algo.Empty_map handles nil and empty maps
		return nil
	}

	newMap := make(map[K]V, len(m))
	for k, v := range m {
		newMap[k] = v // Basic types are copied by value
	}
	return newMap
}

// Deepcopy creates a deep copy of a map with comparable keys and values that implement the Deepcopyable interface.
// Keys must be comparable. Values must implement the constraint.Deepcopyable interface.
func Deepcopy[K interface{ comparable; constraint.NonPointerBasic }, V constraint.Deepcopyable[V]](m map[K]V) map[K]V {
	// Consistent with slice deepcopy, return nil for nil or empty maps.
	if m == nil || algo.Empty_map(m) { // Assuming algo.Empty_map handles nil and empty maps
		return nil
	}

	newMap := make(map[K]V, len(m))
	for k, v := range m {
		newMap[k] = v.Deepcopy() // Call Deepcopy method on each value
	}
	return newMap
}
