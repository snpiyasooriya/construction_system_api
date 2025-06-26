package dto

import (
	"time"

	"gorm.io/datatypes"
)

// ScheduleItemCreateInputDTO represents the input for creating a schedule item
type ScheduleItemCreateInputDTO struct {
	Name            string         `json:"name" binding:"required" example:"Steel Beam"`
	ShapeID         uint           `json:"shape_id" binding:"required" example:"1"`
	ScheduleID      uint           `json:"schedule_id" binding:"required" example:"1"`
	ShapeDimensions datatypes.JSON `json:"-" binding:"-"`
}

// ScheduleItemCreateOutputDTO represents the output after creating a schedule item
type ScheduleItemCreateOutputDTO struct {
	ID              uint           `json:"id" example:"1"`
	Name            string         `json:"name" example:"Steel Beam"`
	ShapeID         uint           `json:"shape_id" example:"1"`
	ScheduleID      uint           `json:"schedule_id" example:"1"`
	ShapeDimensions datatypes.JSON `json:"shape_dimensions" swaggertype:"object"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// ScheduleItemGetByIDOutputDTO represents the output for getting a schedule item by ID
type ScheduleItemGetByIDOutputDTO struct {
	ID              uint           `json:"id" example:"1"`
	Name            string         `json:"name" example:"Steel Beam"`
	ShapeID         uint           `json:"shape_id" example:"1"`
	ShapeName       string         `json:"shape_name" example:"Rectangle"`
	ScheduleID      uint           `json:"schedule_id" example:"1"`
	ScheduleName    string         `json:"schedule_name" example:"Foundation Schedule"`
	ShapeDimensions datatypes.JSON `json:"shape_dimensions" swaggertype:"object"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// ScheduleItemGetByScheduleOutputDTO represents the output for getting schedule items by schedule ID
type ScheduleItemGetByScheduleOutputDTO struct {
	ID              uint           `json:"id" example:"1"`
	Name            string         `json:"name" example:"Steel Beam"`
	ShapeID         uint           `json:"shape_id" example:"1"`
	ShapeName       string         `json:"shape_name" example:"Rectangle"`
	ShapeDimensions datatypes.JSON `json:"shape_dimensions" swaggertype:"object"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// ScheduleItemUpdateInputDTO represents the input for updating a schedule item
type ScheduleItemUpdateInputDTO struct {
	ID              uint           `json:"id" binding:"required" example:"1"`
	Name            string         `json:"name" binding:"required" example:"Updated Steel Beam"`
	ShapeID         uint           `json:"shape_id" binding:"required" example:"2"`
	ShapeDimensions datatypes.JSON `json:"shape_dimensions" binding:"required" swaggertype:"object"`
}

// ScheduleItemUpdateOutputDTO represents the output after updating a schedule item
type ScheduleItemUpdateOutputDTO struct {
	ID              uint           `json:"id" example:"1"`
	Name            string         `json:"name" example:"Updated Steel Beam"`
	ShapeID         uint           `json:"shape_id" example:"2"`
	ScheduleID      uint           `json:"schedule_id" example:"1"`
	ShapeDimensions datatypes.JSON `json:"shape_dimensions" swaggertype:"object"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// ScheduleItemDeleteOutputDTO represents the output after deleting a schedule item
type ScheduleItemDeleteOutputDTO struct {
	ID        uint      `json:"id" example:"1"`
	DeletedAt time.Time `json:"deleted_at"`
	Message   string    `json:"message" example:"Schedule item deleted successfully"`
}
