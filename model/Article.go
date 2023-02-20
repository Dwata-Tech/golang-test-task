package model

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Nickname  string    `json:"nickname"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Comments  []Comment `json:"comments"`
}

func (m *Article) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = time.Now()
	return nil
}
