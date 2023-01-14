package database

import (
	"context"
	"time"

	"github.com/gabrielgaspar447/go-blog-api/models"
)

func seedPosts() {
	var data = []models.Post{
		{
			Id:      1,
			Title:   "First Post",
			Content: "This is the first post",
			UserId:  1,
		},
		{
			Id:      2,
			Title:   "Second Post",
			Content: "This is the second post",
			UserId:  2,
		},
	}

	db.Create(&data)
}

func PostCreate(ctx context.Context, input *models.Post) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return db.WithContext(timeoutCtx).Create(input).Error
}

func PostList(ctx context.Context, includeUser bool) (*[]models.Post, error) {
	posts := make([]models.Post, 0)
	var err error

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if includeUser {
		err = db.WithContext(timeoutCtx).Preload("User").Find(&posts).Error
	} else {
		err = db.WithContext(timeoutCtx).Find(&posts).Error
	}

	return &posts, err
}

func PostGetById(ctx context.Context, id uint, includeUser bool) (*models.Post, error) {
	post := &models.Post{}
	var err error

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if includeUser {
		err = db.WithContext(timeoutCtx).Preload("User").Find(post, id).Error
	} else {
		err = db.WithContext(timeoutCtx).Find(post, id).Error
	}

	return post, err
}

func PostSearch(ctx context.Context, query string, includeUser bool) (*[]models.Post, error) {
	posts := make([]models.Post, 0)
	var err error

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if includeUser {
		err = db.WithContext(timeoutCtx).Preload("User").Where("title LIKE ?", "%"+query+"%").Find(&posts).Error
	} else {
		err = db.WithContext(timeoutCtx).Where("title LIKE ?", "%"+query+"%").Find(&posts).Error
	}

	return &posts, err
}

func PostUpdate(ctx context.Context, input *models.Post, id uint) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return db.WithContext(timeoutCtx).Model(&models.Post{}).Where("id = ?", id).Updates(input).Error
}

func PostDelete(ctx context.Context, id uint) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return db.WithContext(timeoutCtx).Delete(&models.Post{}, id).Error
}
