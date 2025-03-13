package dto

import (
	"gorm.io/datatypes"
	"time"
)

type ScheduleItemCromeCreateInputDTO struct {
	Name            string         `json:"name" binding:"required"`
	ShapeID         uint           `json:"shape_id" binding:"required"`
	ScheduleID      uint           `json:"schedule_id" binding:"required"`
	ShapeDimensions datatypes.JSON `json:"shape_dimensions" binding:"required"`
}

type ScheduleItemCromeCreateOutputDTO struct {
	ID              uint           `json:"id"`
	Name            string         `json:"name"`
	ShapeID         uint           `json:"shape_id"`
	ScheduleID      uint           `json:"schedule_id"`
	ShapeDimensions datatypes.JSON `json:"shape_dimensions"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
