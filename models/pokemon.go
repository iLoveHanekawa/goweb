package models

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	Name string
	Type string
}
