package db

import (
    "github.com/fercho920/ecommerce-go/models"
)

func Migrate() error {
    // Realizar migraciones de todos los modelos necesarios
    if err := DB.AutoMigrate(&models.User{}, &models.Address{}); err != nil {
        return err
    }

    return nil
}
