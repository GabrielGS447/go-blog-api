package services

import (
	"context"

	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gabrielgaspar447/go-blog-api/database"
	"github.com/gabrielgaspar447/go-blog-api/errs"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(ctx context.Context, input *models.User) (string, error) {
	user, err := database.UserFindByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	if user.Id != 0 {
		return "", errs.ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return "", err
	}

	input.Password = string(hash)

	err = database.UserCreate(ctx, input)
	if err != nil {
		return "", err
	}

	input.Password = "" // Clear password just in case

	return auth.SignJWT(input.Id)
}

func UserLogin(ctx context.Context, input *models.LoginDTO) (string, error) {
	user, err := database.UserFindByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	if user.Id == 0 {
		return "", errs.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errs.ErrInvalidPassword
	}

	return auth.SignJWT(user.Id)
}

func UserList(ctx context.Context, includePosts bool) (*[]models.User, error) {
	users, err := database.UserList(ctx, includePosts)
	if err != nil {
		return nil, err
	}

	if includePosts {
		for i := range *users {
			for j := range (*users)[i].Posts {
				(*users)[i].Posts[j].UserId = 0
			}
		}
	}

	return users, nil
}

func UserGetById(ctx context.Context, id uint, includePosts bool) (*models.User, error) {
	user, err := database.UserGetById(ctx, id, includePosts)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, errs.ErrUserNotFound
	}

	if includePosts {
		for i := range user.Posts {
			user.Posts[i].UserId = 0
		}
	}

	return user, nil
}

func UserDeleteSelf(ctx context.Context, id uint) error {
	return database.UserDeleteById(ctx, id)
}
