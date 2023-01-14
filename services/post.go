package services

import (
	"github.com/gabrielgaspar447/go-blog-api/database"
	"github.com/gabrielgaspar447/go-blog-api/errs"
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func PostCreate(input *models.Post) error {
	return database.PostCreate(input)
}

func PostList(includeUser bool) (*[]models.Post, error) {
	posts, err := database.PostList(includeUser)
	if err != nil {
		return nil, err
	}

	if includeUser {
		for i := range *posts {
			(*posts)[i].User.Password = ""
			(*posts)[i].User.ID = 0
			(*posts)[i].User.CreatedAt = nil
			(*posts)[i].User.UpdatedAt = nil
		}
	}

	return posts, nil
}

func PostGetById(id uint, includeUser bool) (*models.Post, error) {
	post, err := database.PostGetById(id, includeUser)
	if err != nil {
		return nil, err
	}

	if post.ID == 0 {
		return nil, errs.ErrPostNotFound
	}

	if includeUser {
		post.User.Password = ""
		post.User.ID = 0
		post.User.CreatedAt = nil
		post.User.UpdatedAt = nil
	}

	return post, nil
}

func PostSearch(query string, includeUser bool) (*[]models.Post, error) {
	posts, err := database.PostSearch(query, includeUser)
	if err != nil {
		return nil, err
	}

	if includeUser {
		for i := range *posts {
			(*posts)[i].User.Password = ""
			(*posts)[i].User.ID = 0
			(*posts)[i].User.CreatedAt = nil
			(*posts)[i].User.UpdatedAt = nil
		}
	}

	return posts, nil
}

func PostUpdate(input *models.Post, userId uint) error {
	post, err := database.PostGetById(input.ID, false)
	if err != nil {
		return err
	}

	if post.ID == 0 {
		return errs.ErrPostNotFound
	} else if post.UserID != userId {
		return errs.ErrPostNotOwned
	}

	return database.PostUpdate(input, input.ID)
}

func PostDelete(id uint, userId uint) error {
	post, err := database.PostGetById(id, false)
	if err != nil {
		return err
	}

	if post.ID == 0 {
		return errs.ErrPostNotFound
	} else if post.UserID != userId {
		return errs.ErrPostNotOwned
	}

	return database.PostDelete(id)
}
