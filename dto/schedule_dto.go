package dto

import (
	"github.com/snpiyasooriya/construction_design_api/constants"
	"time"
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
