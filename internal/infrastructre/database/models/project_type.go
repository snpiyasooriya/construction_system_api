package models

import "gorm.io/gorm"

type ProjectType struct {
	gorm.Model
	Type     string `gorm:"type:varchar(20);not null;unique"`
	Projects []Project
}
