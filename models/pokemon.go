package models

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Type  string `json:"type" form:"type"`
	Level int    `json:"level" form:"level"`
}
