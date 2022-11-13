package models

import (
	"time"
)

type User struct {
	ID          uint       `json:"id,omitempty" gorm:"primaryKey"`
	DisplayName string     `json:"display_name" gorm:"not null" binding:"required,min=3,max=20"`
	Email       string     `json:"email" gorm:"unique; not null" binding:"required,email"`
	Password    string     `json:"password,omitempty" gorm:"not null" binding:"required,min=6"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Posts       []Post     `json:"posts,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
