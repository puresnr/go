package ecode

import "errors"

// Ecode represents an error type identified by a string code.
// It is designed to work with the standard errors.Is function.
// Two Ecode instances are considered equal by errors.Is if they are both of type Ecode
// and were created with the same code string.
// Ecode is intended to be a source of errors, not a wrapper around other errors,
// so it does not implement the Unwrap method.
type Ecode struct{ error }

func New(ecode string) *Ecode { return &Ecode{error: errors.New(ecode)} }

func (e *Ecode) Is(target error) bool {
	if target == nil {return e == nil}
	
	et, ok := target.(*Ecode)
	if !ok {
		return false
	}

	return e.Error() == et.Error()
}