package models

import (
	"time"
)

type User struct {
	Id          uint       `json:"id,omitempty" gorm:"primaryKey"`
	DisplayName string     `json:"display_name" gorm:"not null"`
	Email       string     `json:"email" gorm:"unique; not null"`
	Password    string     `json:"password,omitempty" gorm:"not null"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Posts       []Post     `json:"posts,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) SanitizeToJson() {
	u.Password = ""

	if u.Posts != nil {
		for j := range u.Posts {
			u.Posts[j].UserId = 0
		}
	}
}

type CreateUserDTO struct {
	DisplayName string `json:"display_name" binding:"required,min=3,max=20"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
}

func (dto *CreateUserDTO) ToModel() *User {
	return &User{
		DisplayName: dto.DisplayName,
		Email:       dto.Email,
		Password:    dto.Password,
	}
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
