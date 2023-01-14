package services

import (
	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gabrielgaspar447/go-blog-api/database"
	"github.com/gabrielgaspar447/go-blog-api/errs"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(input *models.User) (string, error) {
	user, err := database.UserFindByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user.ID != 0 {
		return "", errs.ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return "", err
	}

	input.Password = string(hash)

	err = database.UserCreate(input)
	if err != nil {
		return "", err
	}

	input.Password = "" // Clear password just in case

	return auth.SignJWT(input.ID)
}

func UserLogin(input *models.LoginDTO) (string, error) {
	user, err := database.UserFindByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", errs.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errs.ErrInvalidPassword
	}

	return auth.SignJWT(user.ID)
}

func UserList(users *[]models.User, includePosts bool) error {
	err := database.UserList(users, includePosts)
	if err != nil {
		return err
	}

	if includePosts {
		for i := range *users {
			for j := range (*users)[i].Posts {
				(*users)[i].Posts[j].UserID = 0
			}
		}
	}

	return nil
}

func UserGetById(user *models.User, id uint, includePosts bool) error {
	err := database.UserGetById(user, id, includePosts)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errs.ErrUserNotFound
	}

	if includePosts {
		for i := range user.Posts {
			user.Posts[i].UserID = 0
		}
	}

	return nil
}

func UserDeleteSelf(id uint) error {
	return database.UserDeleteById(id)
}
