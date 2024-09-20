package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name          string `gorm:"unique;not null"`
	ProjectTypeID uint   `gorm:"not null"`
	ProjectType   ProjectType
	Schedules     []Schedule
}
