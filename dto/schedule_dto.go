package dto

import "time"

type ScheduleCreateInputDTO struct {
	Name        string
	Description string
	ProjectID   uint
}
type ScheduleCreateOutputDTO struct {
	ID          uint
	Name        string
	Description string
	ProjectID   uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ScheduleGetByProjectOutputDTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProjectID   uint      `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
