package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ScheduleItem struct {
	gorm.Model
	Name            string `gorm:"type:varchar(20);not null"`
	ShapeID         uint
	Shape           Shape
	ScheduleID      uint
	Schedule        Schedule
	ShapeDimensions datatypes.JSON
}
