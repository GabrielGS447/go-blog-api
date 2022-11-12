package repositories

import (
	"github.com/gabrielgaspar447/go-blog-api/constants"
	"github.com/gabrielgaspar447/go-blog-api/db"
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func UserFindByEmail(email string) (models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil && err.Error() == constants.GormNotFound {
		return user, nil
	}
	return user, err
}

func UserCreate(input *models.User) error {
	return db.DB.Create(input).Error
}
