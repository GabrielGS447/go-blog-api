package utils

import "github.com/gabrielgaspar447/go-blog-api/constants"

func GetServiceErrorResponse(err error) (int, string) {
	switch errMsg := err.Error(); errMsg {
	case constants.UserAlreadyExists:
		return constants.HTTP_Conflict, constants.UserAlreadyExists
	case constants.UserNotFound:
		return constants.HTTP_NotFound, constants.UserNotFound
	case constants.InvalidPassword:
		return constants.HTTP_Unauthorized, constants.InvalidPassword
	case constants.PostNotFound:
		return constants.HTTP_NotFound, constants.PostNotFound
	default:
		return constants.HTTP_InternalServerError, constants.SomethingWentWrong
	}
}
