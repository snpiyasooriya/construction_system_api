package dto

import (
	"time"

	"github.com/snpiyasooriya/construction_design_api/constants"
)

type ScheduleCreateInputDTO struct {
	ID           uint                     `json:"id,omitempty" binding:"-"`
	ScheduleID   string                   `json:"schedule_id"`
	ProjectID    uint                     `json:"project_id"`
	Description  string                   `json:"description,omitempty" binding:"required"`
	RequiredDate time.Time                `json:"required_date" binding:"required"`
	SchedularID  uint                     `json:"schedular_id" binding:"required"`
	Status       constants.ScheduleStatus `json:"status,omitempty"`
	Note         string                   `json:"note,omitempty"`
	CreatedAt    time.Time                `json:"created_at,omitempty" binding:"-"`
	UpdatedAt    time.Time                `json:"updated_at,omitempty" binding:"-"`
}

type ScheduleGetByProjectOutputDTO struct {
	ID           uint                     `json:"id,omitempty"`
	ScheduleID   string                   `json:"schedule_id"`
	ProjectID    uint                     `json:"project_id"`
	Description  string                   `json:"description,omitempty"`
	RequiredDate time.Time                `json:"required_date"`
	SchedularID  uint                     `json:"schedular_id"`
	Schedular    string                   `json:"schedular"`
	ReviewerID   uint                     `json:"reviewer_id"`
	Reviewer     string                   `json:"reviewer"`
	Status       constants.ScheduleStatus `json:"status,omitempty"`
	Note         string                   `json:"note,omitempty"`
	CreatedAt    time.Time                `json:"created_at,omitempty"`
	UpdatedAt    time.Time                `json:"updated_at,omitempty"`
}

// ScheduleGetByIDOutputDTO represents the output for getting a schedule by ID
type ScheduleGetByIDOutputDTO struct {
	ID           uint                     `json:"id"`
	ScheduleID   string                   `json:"schedule_id"`
	ProjectID    uint                     `json:"project_id"`
	Description  string                   `json:"description"`
	RequiredDate time.Time                `json:"required_date"`
	SchedularID  uint                     `json:"schedular_id"`
	Schedular    string                   `json:"schedular"`
	ReviewerID   uint                     `json:"reviewer_id"`
	Reviewer     string                   `json:"reviewer"`
	Status       constants.ScheduleStatus `json:"status"`
	Note         string                   `json:"note"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
}

// ScheduleUpdateDTO represents the input for updating a schedule
type ScheduleUpdateDTO struct {
	ID           uint      `json:"id"`
	Description  string    `json:"description"`
	RequiredDate time.Time `json:"required_date"`
	ReviewerID   uint      `json:"reviewer_id"`
	Status       string    `json:"status"`
	Note         string    `json:"note"`
}

// ScheduleUpdateOutputDTO represents the output after updating a schedule
type ScheduleUpdateOutputDTO struct {
	ID           uint      `json:"id"`
	ScheduleID   string    `json:"schedule_id"`
	ProjectID    uint      `json:"project_id"`
	Description  string    `json:"description"`
	RequiredDate time.Time `json:"required_date"`
	SchedularID  uint      `json:"schedular_id"`
	Schedular    string    `json:"schedular"`
	ReviewerID   uint      `json:"reviewer_id"`
	Reviewer     string    `json:"reviewer"`
	Status       string    `json:"status"`
	Note         string    `json:"note"`
	UpdatedAt    time.Time `json:"updated_at"`
}
