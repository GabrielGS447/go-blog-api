package models

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	DisplayName string    `json:"display_name" gorm:"not null" binding:"required,min=3,max=20"`
	Email       string    `json:"email" gorm:"unique; not null" binding:"required,email"`
	Password    string    `json:"password" gorm:"not null" binding:"required,min=6"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Posts       []Post    `json:"posts"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
