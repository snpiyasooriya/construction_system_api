package dto

import (
	"time"
)

// ShapeDTO represents a shape with appropriate JSON tags to control serialization/deserialization
type ShapeDTO struct {
	// ID is read-only (output only)
	ID uint `json:"id,omitempty" binding:"-"`

	// Required fields for creation
	Name       string                 `json:"name" binding:"required"`
	Path       string                 `json:"path"`
	Dimensions map[string]interface{} `json:"dimensions" binding:"required"`

	// Timestamps are read-only (output only)
	CreatedAt time.Time `json:"created_at,omitempty" binding:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" binding:"-"`
}

// ShapesGetDTO represents a collection of shapes
type ShapesGetDTO struct {
	Shapes []ShapeDTO `json:"shapes"`
}
