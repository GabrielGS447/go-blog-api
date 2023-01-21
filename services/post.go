package services

import (
	"context"

	"github.com/gabrielgs447/go-blog-api/database"
	"github.com/gabrielgs447/go-blog-api/errs"
	"github.com/gabrielgs447/go-blog-api/models"
)

type PostServiceInterface interface {
	Create(ctx context.Context, input *models.Post) error
	List(ctx context.Context, includeUser bool) (*[]models.Post, error)
	GetById(ctx context.Context, id uint, includeUser bool) (*models.Post, error)
	Search(ctx context.Context, query string, includeUser bool) (*[]models.Post, error)
	Update(ctx context.Context, input *models.Post, userId uint) error
	Delete(ctx context.Context, id uint, userId uint) error
}

type postService struct {
	postRepository database.PostRepositoryInterface
}

func NewPostService(r database.PostRepositoryInterface) PostServiceInterface {
	return &postService{
		r,
	}
}

func (s *postService) Create(ctx context.Context, input *models.Post) error {
	return s.postRepository.Create(ctx, input)
}

func (s *postService) List(ctx context.Context, includeUser bool) (*[]models.Post, error) {
	posts, err := s.postRepository.List(ctx, includeUser)
	if err != nil {
		return nil, err
	}

	for _, post := range *posts {
		post.SanitizeToJson()
	}

	return posts, nil
}

func (s *postService) GetById(ctx context.Context, id uint, includeUser bool) (*models.Post, error) {
	post, err := s.postRepository.GetById(ctx, id, includeUser)
	if err != nil {
		return nil, err
	}

	if post.Id == 0 {
		return nil, errs.ErrPostNotFound
	}

	post.SanitizeToJson()

	return post, nil
}

func (s *postService) Search(ctx context.Context, query string, includeUser bool) (*[]models.Post, error) {
	posts, err := s.postRepository.Search(ctx, query, includeUser)
	if err != nil {
		return nil, err
	}

	for _, post := range *posts {
		post.SanitizeToJson()
	}

	return posts, nil
}

func (s *postService) Update(ctx context.Context, input *models.Post, userId uint) error {
	post, err := s.postRepository.GetById(ctx, input.Id, false)
	if err != nil {
		return err
	}

	if post.Id == 0 {
		return errs.ErrPostNotFound
	} else if post.UserId != userId {
		return errs.ErrPostNotOwned
	}

	return s.postRepository.Update(ctx, input, input.Id)
}

func (s *postService) Delete(ctx context.Context, id uint, userId uint) error {
	post, err := s.postRepository.GetById(ctx, id, false)
	if err != nil {
		return err
	}

	if post.Id == 0 {
		return errs.ErrPostNotFound
	} else if post.UserId != userId {
		return errs.ErrPostNotOwned
	}

	return s.postRepository.Delete(ctx, id)
}
