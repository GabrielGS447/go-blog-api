package database

import (
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func seedUsers() {
	var data = []models.User{
		{
			ID:          1,
			DisplayName: "John Doe",
			Email:       "johndoe@go.com",
			Password:    "$2a$08$I/wJJtinKh5jEjZjRGsVUes2Jfo.ZFe4n0D7amPHkmONzX4dGuEHy", // "123456"
		},
		{
			ID:          2,
			DisplayName: "Jane Doe",
			Email:       "janedoe@go.com",
			Password:    "$2a$08$I/wJJtinKh5jEjZjRGsVUes2Jfo.ZFe4n0D7amPHkmONzX4dGuEHy", // "123456"
		},
	}

	db.Create(&data)
}

func UserFindByEmail(email string) (models.User, error) {
	var user models.User
	err := db.Limit(1).Find(&user, "email = ?", email).Error
	return user, err
}

func UserCreate(input *models.User) error {
	return db.Create(input).Error
}

func UserList(users *[]models.User, includePosts bool) error {

	if includePosts {
		return db.Omit("Password").Preload("Posts").Find(users).Error
	}

	return db.Omit("Password").Find(users).Error
}

func UserGetById(user *models.User, id uint, includePosts bool) error {
	if includePosts {
		return db.Omit("Password").Preload("Posts").Find(user, id).Error
	}

	return db.Omit("Password").Find(user, id).Error
}

func UserDeleteById(id uint) error {
	return db.Delete(&models.User{}, id).Error
}
