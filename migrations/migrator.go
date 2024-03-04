package migrations

import (
	"goweb/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Pokemon{})
}
