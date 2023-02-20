package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Nickname string    `json:"nickname"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
}
