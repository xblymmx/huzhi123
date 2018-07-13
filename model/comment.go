package model

import "time"

// todo
type Comment struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `sql:"index" json:"deletedAt"`
	Status      int        `json:"status"`
	Content     string     `json:"content"`
	HTMLContent string     `json:"htmlContent"`
	ContentType int        `json:"contentType"`
	ParentID    uint       `json:"parentID"` // direct parent comment id
	Parents     []Comment  `json:"parents"`  // all parent comments
	SourceName  string     `json:"sourceName"`
	SourceID    uint       `json:"sourceID"`
	UserID      uint       `json:"userID"`
	User        User       `json:"user"`
}
