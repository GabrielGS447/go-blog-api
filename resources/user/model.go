package user

import (
	"time"

	"github.com/gabrielgaspar447/go-blog-api/resources/post"
)

type User struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	DisplayName string      `json:"display_name" gorm:"unique; not null"`
	Email       string      `json:"email" gorm:"unique; not null"`
	Password    string      `json:"password" gorm:"not null"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Posts       []post.Post `json:"posts"`
}
