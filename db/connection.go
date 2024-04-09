package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres password=admin dbname=go_crud port=5433"
var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	log.Println("Connected to database")
}
