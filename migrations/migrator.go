package migrations

import (
	"goweb/db"
	"goweb/models"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	// AutoMigrate will create the table if it does not exist
	db.AutoMigrate(&models.Pokemon{})
}

func InitializeDB() (*gorm.DB, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	MigrateModels(db)
	return db, nil
}
