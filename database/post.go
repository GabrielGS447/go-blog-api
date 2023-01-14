package database

import (
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func seedPosts() {
	var data = []models.Post{
		{
			ID:      1,
			Title:   "First Post",
			Content: "This is the first post",
			UserID:  1,
		},
		{
			ID:      2,
			Title:   "Second Post",
			Content: "This is the second post",
			UserID:  2,
		},
	}

	db.Create(&data)
}

func PostCreate(input *models.Post) error {
	return db.Create(input).Error
}

func PostList(posts *[]models.Post, includeUser bool) error {
	if includeUser {
		return db.Preload("User").Find(posts).Error
	}

	return db.Find(posts).Error
}

func PostGetById(post *models.Post, id uint, includeUser bool) error {
	if includeUser {
		return db.Preload("User").Find(post, id).Error
	}

	return db.Find(post, id).Error
}

func PostSearch(posts *[]models.Post, query string, includeUser bool) error {
	if includeUser {
		return db.Preload("User").Where("title LIKE ?", "%"+query+"%").Find(posts).Error
	}

	return db.Where("title LIKE ?", "%"+query+"%").Find(posts).Error
}

func PostUpdate(input *models.Post, id uint) error {
	return db.Model(&models.Post{}).Where("id = ?", id).Updates(input).Error
}

func PostDelete(id uint) error {
	return db.Delete(&models.Post{}, id).Error
}
