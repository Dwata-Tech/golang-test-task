package model

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Nickname string    `json:"nickname"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
}

type ArticleListResponse struct {
	Nickname  string    `json:"nickname"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}
