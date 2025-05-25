package models

import (
	"time"
)

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderRequest struct {
	DiscountCodes string `json:"discount_codes,omitempty"`
	Items      []struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required"`
	} `json:"items" binding:"required"`
}

type Order struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Items        []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Products     []Product  `json:"products" gorm:"many2many:order_items"`
	OriginalTotal float64   `json:"original_total"`
	DiscountCode string     `json:"discount_code"`
	FinalTotal   float64    `json:"final_total"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}