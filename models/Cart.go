package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID      string `json:"user_id"`
	Products_ID string `json:"products_id"`
	Checkout    bool   `json:"checkout"`
	CreatedAt   int64  `gorm:"column:created_at" json:"created_at" bson:"created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at" json:"updated_at" bson:"updated_at"`
}
