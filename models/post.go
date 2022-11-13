package models

import "time"

type Post struct {
	ID        uint       `json:"id,omitempty" gorm:"primaryKey"`
	Title     string     `json:"title" gorm:"not null" binding:"required,min=3,max=100"`
	Content   string     `json:"content" gorm:"not null" binding:"required,min=3"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UserID    uint       `json:"user_id,omitempty" gorm:"not null"`
	User      *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
