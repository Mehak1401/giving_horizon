package models

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	UserID       uint    `json:"user_id"`
	FundraiserID uint    `json:"fundraiser_id"`
	Amount       float64 `json:"amount"`
	PaymentID    string  `json:"payment_id"`
}