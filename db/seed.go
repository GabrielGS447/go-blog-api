package db

import (
	"github.com/gabrielgaspar447/go-blog-api/models"
	"gorm.io/gorm"
)

var usersSeed = []models.User{
	{
		ID:          1,
		DisplayName: "John Doe",
		Email:       "johndoe@go.com",
		Password:    "123456",
	},
	{
		ID:          2,
		DisplayName: "Jane Doe",
		Email:       "janedoe@go.com",
		Password:    "123456",
	},
}

var postsSeed = []models.Post{
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

func seed(db *gorm.DB) {
	db.Create(&usersSeed)
	db.Create(&postsSeed)
}
