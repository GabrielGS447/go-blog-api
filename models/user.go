package models

import (
	"time"
)

type User struct {
	Id          uint       `json:"id,omitempty" gorm:"primaryKey" binding:"isdefault"`
	DisplayName string     `json:"display_name" gorm:"not null" binding:"required,min=3,max=20"`
	Email       string     `json:"email" gorm:"unique; not null" binding:"required,email"`
	Password    string     `json:"password,omitempty" gorm:"not null" binding:"required,min=6"`
	CreatedAt   *time.Time `json:"created_at,omitempty" binding:"isdefault"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" binding:"isdefault"`
	Posts       []Post     `json:"posts,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" binding:"isdefault"`
}

func (u *User) SanitizeToJson() {
	u.Password = ""

	if u.Posts != nil {
		for j := range u.Posts {
			u.Posts[j].UserId = 0
		}
	}
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
