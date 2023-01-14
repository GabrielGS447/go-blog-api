package errs

import "errors"

var (
	ErrPostNotFound = errors.New("post not found")
	ErrPostNotOwned = errors.New("post does not belong to user")
)
