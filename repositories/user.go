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

func UserList(users *[]models.User, includePosts bool) error {

	if includePosts {
		return db.DB.Omit("Password").Preload("Posts").Find(users).Error
	}

	return db.DB.Omit("Password").Find(users).Error
}

func UserFindById(user *models.User, id uint, includePosts bool) error {
	if includePosts {
		return db.DB.Omit("Password").Preload("Posts").Find(user, id).Error
	}

	return db.DB.Omit("Password").Find(user, id).Error
}

func UserDeleteById(id uint) error {
	return db.DB.Delete(&models.User{}, id).Error
}
