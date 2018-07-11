package model

import "time"

// todo
type Article struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `sql:"index" json:"deletedAt"`
	Name          string     `json:"name"`
	BrowseCount   uint       `json:"browseCount"`
	CommentCount  uint       `json:"commentCount"`
	CollectCount  uint       `json:"collectCount"`
	Status        int        `json:"status"`
	Content       string     `json:"content"`
	HTMLContent   string     `json:"htmlContent"`
	ContentType   int        `json:"contentType"`

	Comments []Comment `json:"comments"` // gorm
	Categories []Category `json:"categories"` // gorm
}