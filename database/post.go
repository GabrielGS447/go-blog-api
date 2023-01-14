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

func PostList(includeUser bool) (*[]models.Post, error) {
	posts := make([]models.Post, 0)
	var err error

	if includeUser {
		err = db.Preload("User").Find(&posts).Error
	} else {
		err = db.Find(&posts).Error
	}

	return &posts, err
}

func PostGetById(id uint, includeUser bool) (*models.Post, error) {
	post := &models.Post{}
	var err error

	if includeUser {
		err = db.Preload("User").Find(post, id).Error
	} else {
		err = db.Find(post, id).Error
	}

	return post, err
}

func PostSearch(query string, includeUser bool) (*[]models.Post, error) {
	posts := make([]models.Post, 0)
	var err error

	if includeUser {
		err = db.Preload("User").Where("title LIKE ?", "%"+query+"%").Find(&posts).Error
	} else {
		err = db.Where("title LIKE ?", "%"+query+"%").Find(&posts).Error
	}

	return &posts, err
}

func PostUpdate(input *models.Post, id uint) error {
	return db.Model(&models.Post{}).Where("id = ?", id).Updates(input).Error
}

func PostDelete(id uint) error {
	return db.Delete(&models.Post{}, id).Error
}
