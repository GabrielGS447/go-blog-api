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
