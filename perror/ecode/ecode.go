package ecode

import "errors"

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