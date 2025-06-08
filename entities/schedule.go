package entities

import (
	"github.com/snpiyasooriya/construction_design_api/constants"
	"time"
)

type Schedule struct {
	ID           uint
	ProjectID    uint
	Project      Project
	ScheduleID   string
	Description  string
	RequiredDate time.Time
	SchedularID  uint
	Schedular    User
	ReviewerID   uint
	Reviewer     User
	ReviewedDate time.Time
	Status       constants.ScheduleStatus
	Note         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
