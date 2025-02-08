package models

import "gorm.io/gorm"

type ClothingDemand struct {
	gorm.Model
	RequesterID  uint   `json:"requester_id"`
	ClothingType string `json:"clothing_type"`
	Size         string `json:"size"`
	Quantity     int    `json:"quantity"`
	Description  string `json:"description"`
	Location     string `json:"location"`
}
