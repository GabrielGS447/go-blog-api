package post

import (
	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/gabrielgaspar447/go-blog-api/repositories"
)

func createPostService(input *models.Post) error {
	return repositories.PostCreate(input)
}

func listPostsService(includeUser bool) ([]models.Post, error) {
	posts, err := repositories.PostList(includeUser)
	if err != nil {
		return posts, err
	}

	if includeUser {
		for i := range posts {
			posts[i].User.Password = ""
			posts[i].User.ID = 0
			posts[i].User.CreatedAt = nil
			posts[i].User.UpdatedAt = nil
		}
	}

	return posts, nil
}

func getPostByIdService(id uint, includeUser bool) (models.Post, error) {
	post, err := repositories.PostGetById(id, includeUser)
	if err != nil {
		return post, err
	}

	if includeUser {
		post.User.Password = ""
		post.User.ID = 0
		post.User.CreatedAt = nil
		post.User.UpdatedAt = nil
	}

	return post, nil
}
