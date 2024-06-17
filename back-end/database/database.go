package database

import (
	"log"
	"stock_manager_back-end/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() error {
	db, err := gorm.Open(sqlite.Open("farm_stock.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Product{}, &models.Site{}, &models.Stock{}); err != nil {
		return err
	}

	DB = db
	log.Println("Database setup complete")
	return nil
}
