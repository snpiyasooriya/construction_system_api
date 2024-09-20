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
