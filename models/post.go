package models

import "time"

type Post struct {
	Id        uint       `json:"id,omitempty" gorm:"primaryKey" binding:"isdefault"`
	Title     string     `json:"title" gorm:"not null" binding:"required,min=3,max=100"`
	Content   string     `json:"content" gorm:"not null" binding:"required,min=3"`
	CreatedAt *time.Time `json:"created_at,omitempty" binding:"isdefault"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" binding:"isdefault"`
	UserId    uint       `json:"user_id,omitempty" gorm:"not null" binding:"isdefault"`
	User      *User      `json:"user,omitempty" gorm:"foreignKey:UserId"`
}

func (p *Post) SanitizeToJson() {
	if p.User != nil {
		p.User.SanitizeToJson()
		p.User.Id = 0
		p.User.CreatedAt = nil
		p.User.UpdatedAt = nil
	}
}
