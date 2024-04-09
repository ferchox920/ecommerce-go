package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID          int
	Order_Cart  []Cart
	Quantity    int
	Status      string
	TotalAmount float64
	CreatedAt    int64   `gorm:"column:created_at" json:"created_at" bson:"created_at"`
    UpdatedAt    int64   `gorm:"column:updated_at" json:"updated_at" bson:"updated_at"`
}
