package models

import "gorm.io/gorm"

type FoodDemand struct {
	gorm.Model
	RequesterID uint   `json:"requester_id"`
	ItemName    string `json:"item_name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	Location    string `json:"location"`
}
