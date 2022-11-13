package repositories

import (
	"github.com/gabrielgaspar447/go-blog-api/db"
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func PostCreate(input *models.Post) error {
	return db.DB.Create(input).Error
}
