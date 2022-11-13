package repositories

import (
	"github.com/gabrielgaspar447/go-blog-api/db"
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func PostCreate(input *models.Post) error {
	return db.DB.Create(input).Error
}

func PostList(includeUser bool) ([]models.Post, error) {
	var posts []models.Post

	if includeUser {
		return posts, db.DB.Preload("User").Find(&posts).Error
	}

	return posts, db.DB.Find(&posts).Error
}
