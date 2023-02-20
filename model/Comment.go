package model

import (
	"gorm.io/gorm"
	"time"
)

// Comment represents a database model for a comment on an article
type Comment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Nickname  string    `json:"nickname"`
	Content   string    `json:"content"`
	ArticleID uint      `json:"articleID"`
}

func (m *Comment) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = time.Now()
	return nil
}
