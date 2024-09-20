package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Shape struct {
	gorm.Model
	Name       string         `json:"name"`
	Path       string         `json:"path"`
	Dimensions datatypes.JSON `json:"dimensions" gorm:"type:json"`
}
