package constraint

// Basic is a constraint that permits any of Go's basic types, excluding pointer types like uintptr.
// This is used when strict non-pointer types are required.
type Basic interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64 |
	~string | ~bool | ~complex64 | ~complex128
}

// Deepcopyable is an interface for types that can create a deep copy of themselves.
type Deepcopyable[T any] interface {
	Deepcopy() T
}
