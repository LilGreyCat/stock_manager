package database

import (
	"log"
	"stock_manager_back-end/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitializeDatabase initializes the database connection and performs automatic migration.
//
// It opens a connection to the SQLite database file "farm_stock.db" using the GORM library.
// If the connection fails, it returns the error.
//
// It then performs automatic migration for the models.Product, models.Site, and models.Stock structs.
// If the migration fails, it returns the error.
//
// Finally, it assigns the database connection to the global variable DB and logs a message indicating that the database setup is complete.
// It returns nil if everything is successful.
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
