package models

import (
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	ID    uint   `json:"id" form:"id" gorm:"primaryKey" query:"id" xml:"id"`
	Name  string `json:"name" form:"name" query:"name" xml:"name"`
	Type  string `json:"type" form:"type" query:"type" xml:"type"`
	Level int    `json:"level" form:"level" query:"level" xml:"level"`
}
