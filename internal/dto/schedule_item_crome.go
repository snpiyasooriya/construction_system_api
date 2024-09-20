package dto

import (
	"gorm.io/datatypes"
	"time"
)

type ScheduleItemCromeCreateInputDTO struct {
	Name            string
	ShapeID         uint
	ScheduleID      uint
	ShapeDimensions datatypes.JSON
}

type ScheduleItemCromeCreateOutputDTO struct {
	ID              uint
	Name            string
	ShapeID         uint
	ScheduleID      uint
	ShapeDimensions datatypes.JSON
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
