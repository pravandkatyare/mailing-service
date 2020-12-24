package errs

import "errors"

var (
	// ErrUnauthorized is an unauthorized error
	ErrUnauthorized = errors.New("Unauthorized")
	// ErrBadRequest is an bad request error
	ErrBadRequest = errors.New("Bad Request")
)
