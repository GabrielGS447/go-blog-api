package database

import (
	"context"
	"time"

	"github.com/gabrielgaspar447/go-blog-api/models"
)

func seedUsers() {
	var data = []models.User{
		{
			Id:          1,
			DisplayName: "John Doe",
			Email:       "johndoe@go.com",
			Password:    "$2a$08$I/wJJtinKh5jEjZjRGsVUes2Jfo.ZFe4n0D7amPHkmONzX4dGuEHy", // "123456"
		},
		{
			Id:          2,
			DisplayName: "Jane Doe",
			Email:       "janedoe@go.com",
			Password:    "$2a$08$I/wJJtinKh5jEjZjRGsVUes2Jfo.ZFe4n0D7amPHkmONzX4dGuEHy", // "123456"
		},
	}

	db.Create(&data)
}

func UserFindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := db.WithContext(timeoutCtx).Limit(1).Find(&user, "email = ?", email).Error
	return &user, err
}

func UserCreate(ctx context.Context, input *models.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return db.WithContext(timeoutCtx).Omit("Id").Create(input).Error
}

func UserList(ctx context.Context, includePosts bool) (*[]models.User, error) {
	users := make([]models.User, 0)
	var err error

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if includePosts {
		err = db.WithContext(timeoutCtx).Omit("Password").Preload("Posts").Find(&users).Error
	} else {
		err = db.WithContext(timeoutCtx).Omit("Password").Find(&users).Error
	}

	return &users, err
}

func UserGetById(ctx context.Context, id uint, includePosts bool) (*models.User, error) {
	user := &models.User{}
	var err error

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if includePosts {
		err = db.WithContext(timeoutCtx).Omit("Password").Preload("Posts").Find(user, id).Error
	} else {
		err = db.WithContext(timeoutCtx).Omit("Password").Find(user, id).Error
	}

	return user, err
}

func UserDeleteById(ctx context.Context, id uint) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return db.WithContext(timeoutCtx).Delete(&models.User{}, id).Error
}
