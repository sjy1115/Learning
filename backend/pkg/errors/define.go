package errs

import "errors"

var (
	NotFound = errors.New("not found")
	Empty    = errors.New("empty")
	NotEqual = errors.New("not equal")
	NotValid = errors.New("not valid")
)
