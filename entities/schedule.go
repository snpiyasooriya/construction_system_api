package entities

import "time"

type Schedule struct {
	ID                uint `json:"id"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
	Name              string `json:"name"`
	Description       string `json:"description"`
	CreatedBy         User
	ScheduleItemCrome []ScheduleItemCrome
	ProjectID         uint
	Project           Project
}
