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

type PostRepositoryInterface interface {
	Create(ctx context.Context, input *models.Post) error
	List(ctx context.Context, includeUser bool) (*[]models.Post, error)
	GetById(ctx context.Context, id uint, includeUser bool) (*models.Post, error)
	Search(ctx context.Context, query string, includeUser bool) (*[]models.Post, error)
	Update(ctx context.Context, input *models.Post, id uint) error
	Delete(ctx context.Context, id uint) error
}

type postRepository struct{}

func NewPostRepository() PostRepositoryInterface {
	return &postRepository{}
}

func (r *postRepository) Create(ctx context.Context, input *models.Post) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return db.WithContext(timeoutCtx).Create(input).Error
}

func (r *postRepository) List(ctx context.Context, includeUser bool) (*[]models.Post, error) {
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

func (r *postRepository) GetById(ctx context.Context, id uint, includeUser bool) (*models.Post, error) {
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

func (r *postRepository) Search(ctx context.Context, query string, includeUser bool) (*[]models.Post, error) {
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

func (r *postRepository) Update(ctx context.Context, input *models.Post, id uint) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return db.WithContext(timeoutCtx).Model(&models.Post{}).Where("id = ?", id).Updates(input).Error
}

func (r *postRepository) Delete(ctx context.Context, id uint) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return db.WithContext(timeoutCtx).Delete(&models.Post{}, id).Error
}
