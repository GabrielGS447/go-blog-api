package user

import (
	"errors"

	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gabrielgaspar447/go-blog-api/constants"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/gabrielgaspar447/go-blog-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

func signupService(input *models.User) (string, error) {
	user, err := repositories.UserFindByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user.ID != 0 {
		return "", errors.New(constants.UserAlreadyExists)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return "", err
	}

	input.Password = string(hash)

	err = repositories.UserCreate(input)
	if err != nil {
		return "", err
	}

	input.Password = "" // Clear password just in case

	return auth.SignJWT(input)
}

func loginService(input *models.LoginDTO) (string, error) {
	user, err := repositories.UserFindByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", errors.New(constants.UserNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New(constants.InvalidPassword)
	}

	return auth.SignJWT(&user)
}
