package services

import (
	"github.com/gabrielgaspar447/go-blog-api/database"
	"github.com/gabrielgaspar447/go-blog-api/errs"
	"github.com/gabrielgaspar447/go-blog-api/models"
)

func CreatePostService(input *models.Post) error {
	return database.PostCreate(input)
}

func ListPostsService(posts *[]models.Post, includeUser bool) error {
	err := database.PostList(posts, includeUser)
	if err != nil {
		return err
	}

	if includeUser {
		for i := range *posts {
			(*posts)[i].User.Password = ""
			(*posts)[i].User.ID = 0
			(*posts)[i].User.CreatedAt = nil
			(*posts)[i].User.UpdatedAt = nil
		}
	}

	return nil
}

func GetPostByIdService(post *models.Post, id uint, includeUser bool) error {
	err := database.PostGetById(post, id, includeUser)
	if err != nil {
		return err
	}

	if post.ID == 0 {
		return errs.ErrPostNotFound
	}

	if includeUser {
		post.User.Password = ""
		post.User.ID = 0
		post.User.CreatedAt = nil
		post.User.UpdatedAt = nil
	}

	return nil
}

func SearchPostsService(posts *[]models.Post, query string, includeUser bool) error {
	err := database.PostSearch(posts, query, includeUser)
	if err != nil {
		return err
	}

	if includeUser {
		for i := range *posts {
			(*posts)[i].User.Password = ""
			(*posts)[i].User.ID = 0
			(*posts)[i].User.CreatedAt = nil
			(*posts)[i].User.UpdatedAt = nil
		}
	}

	return nil
}

func UpdatePostService(input *models.Post, userId uint) error {
	post := &models.Post{}
	err := database.PostGetById(post, input.ID, false)
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

func DeletePostService(id uint, userId uint) error {
	post := &models.Post{}
	err := database.PostGetById(post, id, false)
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
