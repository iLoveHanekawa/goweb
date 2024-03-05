package models

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	Name  string `json:"name"`
	Type  string `json:"type"`
	Level int    `json:"level"`
}
