package entities

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Title  string  `json:"title"`
	URL    string  `json:"url"`
	Active bool    `json:"active" gorm:"default:true"`
	Key    string  `json:"key"`
	Clicks []Click `json:"clicks"`
}
