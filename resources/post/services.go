package post

import (
	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/gabrielgaspar447/go-blog-api/repositories"
)

func createPostService(input *models.Post) error {
	return repositories.PostCreate(input)
}
