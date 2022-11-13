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

func PostGetById(id uint, includeUser bool) (models.Post, error) {
	var post models.Post

	if includeUser {
		return post, db.DB.Preload("User").First(&post, id).Error
	}

	return post, db.DB.First(&post, id).Error
}
