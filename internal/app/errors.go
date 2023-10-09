package app

import "errors"

var (
	ErrNotFoundOperations = errors.New("not found operations for type")
	ErrNotFoundAccount    = errors.New("not found account")
	ErrInactiveOperation  = errors.New("cannot save this transaction because invalid operation type or inactive")
)
