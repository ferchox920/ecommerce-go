package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    ProductID    string  `gorm:"column:product_id" json:"id" bson:"id"`
    ProductName  string  `gorm:"column:product_name" json:"name" bson:"name"`
    Description  string  `gorm:"column:description" json:"description" bson:"description"`
    Price        float64 `gorm:"column:price" json:"price" bson:"price"`
    ImageURL     string  `gorm:"column:image_url" json:"image_url" bson:"image_url"`
    Rating       float64 `gorm:"column:rating" json:"rating" bson:"rating"`
    CreatedAt    int64   `gorm:"column:created_at" json:"created_at" bson:"created_at"`
    UpdatedAt    int64   `gorm:"column:updated_at" json:"updated_at" bson:"updated_at"`
}
