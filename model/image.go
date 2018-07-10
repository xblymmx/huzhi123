package model

type Image struct {
	ID            uint   `gorm:"primary_key" json:"id"`
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	URL           string `json:"url"`
	Width         uint   `json:"width"`
	Height        uint   `json:"height"`
	Mime          string `json:"mime"`
}
