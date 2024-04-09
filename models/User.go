package models

import (
	"github.com/fercho920/ecommerce-go/constants"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID          `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName      string             `gorm:"not null" json:"first_name"`
	LastName       string             `gorm:"not null" json:"last_name"`
	Password       string             `gorm:"not null" json:"password"`
	Email          string             `gorm:"not null;uniqueIndex" json:"email"`
	Phone          string             `json:"phone"`
	Token          string             `json:"token"`
	RefreshToken   string             `json:"refresh_token"`
	AddressDetails []Address          `json:"address_details" gorm:"foreignKey:UserID"`
	UserType       constants.UserType `json:"user_type"`
}

type Address struct {
	gorm.Model
	Address1  string    `json:"address1"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"` // Cambiado a uuid.UUID para usar UUID
	City      string    `json:"city"`
	Country   string    `json:"country"`
	CreatedAt int64     `gorm:"column:created_at" json:"created_at" bson:"created_at"`
	UpdatedAt int64     `gorm:"column:updated_at" json:"updated_at" bson:"updated_at"`
}
type UserCreate struct {
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"not null;uniqueIndex" json:"email"`
	Phone     string `json:"phone"`
}

type UserUpdateClient struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}
