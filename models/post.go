package models

import "time"

type Post struct {
	Id        uint       `json:"id,omitempty" gorm:"primaryKey"`
	Title     string     `json:"title" gorm:"not null"`
	Content   string     `json:"content" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UserId    uint       `json:"user_id,omitempty" gorm:"not null"`
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

type CreatePostDTO struct {
	Title   string `json:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required,min=3"`
}

func (dto *CreatePostDTO) ToModel(userId uint) *Post {
	return &Post{
		UserId:  userId,
		Title:   dto.Title,
		Content: dto.Content,
	}
}

type UpdatePostDTO struct {
	Title   string `json:"title" binding:"omitempty,min=3,max=100"`
	Content string `json:"content" binding:"omitempty,min=3"`
}

func (dto *UpdatePostDTO) ToModel() *Post {
	return &Post{
		Title:   dto.Title,
		Content: dto.Content,
	}
}
