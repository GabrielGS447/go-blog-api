package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func GetValidationErrors(err error) []string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, e := range ve {
			out[i] = getErrorMsg(e)
		}
		return out
	}
	return nil
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters"
	default:
		return "An unknown error occurred"
	}
}
