package migrations

import (
	"goweb/models"

	"gorm.io/gorm"
)

func InitializeDB(db *gorm.DB) {
	db.AutoMigrate(&models.Pokemon{})
}
