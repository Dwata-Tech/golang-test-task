package model

import "gorm.io/gorm"

// Comment represents a database model for a comment on an article
type Comment struct {
	gorm.Model
	Nickname  string `json:"nickname"`
	Content   string `json:"content"`
	ArticleID uint   `json:"articleID"`
}
