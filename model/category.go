package model

import "time"

type Category struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`

	Name     string `json:"name"`
	Sequence int    `json:"sequence"` // sorting category if at the same level
	ParentID uint   `json:"parent_id"`
}
