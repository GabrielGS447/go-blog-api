package services

import (
	"context"

	"github.com/gabrielgaspar447/go-blog-api/database"
	"github.com/gabrielgaspar447/go-blog-api/errs"
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func PostCreate(ctx context.Context, input *models.Post) error {
	return database.PostCreate(ctx, input)
}

func PostList(ctx context.Context, includeUser bool) (*[]models.Post, error) {
	posts, err := database.PostList(ctx, includeUser)
	if err != nil {
		return nil, err
	}

	if includeUser {
		for i := range *posts {
			(*posts)[i].User.Password = ""
			(*posts)[i].User.Id = 0
			(*posts)[i].User.CreatedAt = nil
			(*posts)[i].User.UpdatedAt = nil
		}
	}

	return posts, nil
}

func PostGetById(ctx context.Context, id uint, includeUser bool) (*models.Post, error) {
	post, err := database.PostGetById(ctx, id, includeUser)
	if err != nil {
		return nil, err
	}

	if post.Id == 0 {
		return nil, errs.ErrPostNotFound
	}

	if includeUser {
		post.User.Password = ""
		post.User.Id = 0
		post.User.CreatedAt = nil
		post.User.UpdatedAt = nil
	}

	return post, nil
}

func PostSearch(ctx context.Context, query string, includeUser bool) (*[]models.Post, error) {
	posts, err := database.PostSearch(ctx, query, includeUser)
	if err != nil {
		return nil, err
	}

	if includeUser {
		for i := range *posts {
			(*posts)[i].User.Password = ""
			(*posts)[i].User.Id = 0
			(*posts)[i].User.CreatedAt = nil
			(*posts)[i].User.UpdatedAt = nil
		}
	}

	return posts, nil
}

func PostUpdate(ctx context.Context, input *models.Post, userId uint) error {
	post, err := database.PostGetById(ctx, input.Id, false)
	if err != nil {
		return err
	}

	if post.Id == 0 {
		return errs.ErrPostNotFound
	} else if post.UserId != userId {
		return errs.ErrPostNotOwned
	}

	return database.PostUpdate(ctx, input, input.Id)
}

func PostDelete(ctx context.Context, id uint, userId uint) error {
	post, err := database.PostGetById(ctx, id, false)
	if err != nil {
		return err
	}

	if post.Id == 0 {
		return errs.ErrPostNotFound
	} else if post.UserId != userId {
		return errs.ErrPostNotOwned
	}

	return database.PostDelete(ctx, id)
}
