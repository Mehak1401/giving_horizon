package models

import "gorm.io/gorm"

type Fundraiser struct {
	gorm.Model
	CreatorID uint   `json:"creator_id"`
	Name      string `json:"name"`
	Amount    float64 `json:"amount"`
	Description string `json:"description"`
}