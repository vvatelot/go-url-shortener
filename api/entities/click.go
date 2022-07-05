package entities

import "gorm.io/gorm"

type Click struct {
	gorm.Model
	LinkID       int `json:"link_id"`
	ResponseCode int `json:"response_code"`
}
