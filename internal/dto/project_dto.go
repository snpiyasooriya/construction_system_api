package dto

import "time"

type ProjectCreateInputDTO struct {
	Name   string `json:"name" validate:"required"`
	TypeID uint   `json:"type_id" validate:"required"`
}

type ProjectCreateOutputDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	TypeID    uint      `json:"type_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ProjectsGetDTO struct {
	Projects []ProjectGetDTO
}

type ProjectGetDTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	ProjectType string    `json:"project_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
