package models

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	DisplayName string    `json:"display_name" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique; not null"`
	Password    string    `json:"password" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Posts       []Post    `json:"posts"`
}
