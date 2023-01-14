package errs

import "errors"

var (
	ErrUnknown   = errors.New("something went wrong, please try again later")
	ErrInvalidId = errors.New("invalid id")
)
