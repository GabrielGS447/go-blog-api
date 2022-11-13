package repositories

import (
	"github.com/gabrielgaspar447/go-blog-api/db"
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func UserFindByEmail(email string) (models.User, error) {
	var user models.User
	err := db.DB.Limit(1).Find(&user, "email = ?", email).Error
	return user, err
}

func UserCreate(input *models.User) error {
	return db.DB.Create(input).Error
}

func UserList(includePosts bool) ([]models.User, error) {
	var users []models.User

	if includePosts {
		return users, db.DB.Omit("Password").Preload("Posts").Find(&users).Error
	}

	return users, db.DB.Omit("Password").Find(&users).Error
}

func UserFindById(id uint, includePosts bool) (models.User, error) {
	var user models.User

	if includePosts {
		return user, db.DB.Omit("Password").Preload("Posts").Find(&user, id).Error
	}

	return user, db.DB.Omit("Password").Find(&user, id).Error
}

func UserDeleteById(id uint) error {
	return db.DB.Delete(&models.User{}, id).Error
}
