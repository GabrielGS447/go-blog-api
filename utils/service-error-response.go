package utils

import "github.com/gabrielgaspar447/go-blog-api/constants"

func GetServiceErrorResponse(err error) (int, string) {
	switch errMsg := err.Error(); errMsg {
	case constants.UserAlreadyExists:
		return constants.Conflict, constants.UserAlreadyExists
	case constants.UserNotFound:
		return constants.NotFound, constants.UserNotFound
	case constants.InvalidPassword:
		return constants.Unauthorized, constants.InvalidPassword
	default:
		return constants.InternalServerError, constants.SomethingWentWrong
	}
}
