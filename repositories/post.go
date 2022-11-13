package repositories

import (
	"github.com/gabrielgaspar447/go-blog-api/db"
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func PostCreate(input *models.Post) error {
	return db.DB.Create(input).Error
}

func PostList(posts *[]models.Post, includeUser bool) error {
	if includeUser {
		return db.DB.Preload("User").Find(&posts).Error
	}

	return db.DB.Find(&posts).Error
}

func PostGetById(post *models.Post, id uint, includeUser bool) error {

	if includeUser {
		return db.DB.Preload("User").Find(&post, id).Error
	}

	return db.DB.First(&post, id).Error
}
