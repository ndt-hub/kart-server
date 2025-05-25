package models

import (
	"time"
)

type Discount struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	DiscountCode  string    `json:"discount_code" gorm:"unique"`
	DiscountValue float64   `json:"discount_value"`
	ValidFrom     time.Time `json:"valid_from"`
	ValidTo       time.Time `json:"valid_to"`
	RemainingCount int      `json:"remaining_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type DiscountResponse struct {
	DiscountCode   string    `json:"discount_code"`
	DiscountValue  float64   `json:"discount_value"`
	ValidFrom      time.Time `json:"valid_from"`
	ValidTo        time.Time `json:"valid_to"`
	RemainingCount int       `json:"remaining_count"`
}