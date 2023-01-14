package database

import (
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

func UserFindByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.Limit(1).Find(&user, "email = ?", email).Error
	return &user, err
}

func UserCreate(input *models.User) error {
	return db.Omit("Id").Create(input).Error
}

func UserList(includePosts bool) (*[]models.User, error) {
	users := make([]models.User, 0)
	var err error

	if includePosts {
		err = db.Omit("Password").Preload("Posts").Find(&users).Error
	} else {
		err = db.Omit("Password").Find(&users).Error
	}

	return &users, err
}

func UserGetById(id uint, includePosts bool) (*models.User, error) {
	user := &models.User{}
	var err error

	if includePosts {
		err = db.Omit("Password").Preload("Posts").Find(user, id).Error
	} else {
		err = db.Omit("Password").Find(user, id).Error
	}

	return user, err
}

func UserDeleteById(id uint) error {
	return db.Delete(&models.User{}, id).Error
}
