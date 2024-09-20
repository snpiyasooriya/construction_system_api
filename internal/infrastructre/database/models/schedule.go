package models

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ProjectID   uint    `json:"project_id"`
	Project     Project `json:"project"`
}
