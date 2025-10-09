package ecode

import "errors"

type Ecode struct{ error }

func New(ecode string) *Ecode { return &Ecode{error: errors.New(ecode)} }