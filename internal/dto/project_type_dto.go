package dto

import "time"

type ProjectTypeCreateInputDTO struct {
	Type string `json:"type" validate:"required"`
}

type ProjectTypeCreateOutputDTO struct {
	ID        uint   `json:"id"`
	Type      string `json:"type"`
	CreatedAt time.Time
}

type ProjectTypeViewInputDTO struct {
	ID uint `json:"id" validate:"required"`
}

type ProjectTypeViewOutputDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Type      uint      `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectTypeViewAllOutputDTO struct {
	ProjectTypes []ProjectTypeViewOutputDTO `json:"project_types"`
}
