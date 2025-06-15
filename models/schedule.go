package models

import (
	"time"

	"github.com/snpiyasooriya/construction_design_api/constants"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	ProjectID    uint                     `json:"project_id"`
	Project      Project                  `json:"project"`
	ScheduleID   string                   `json:"schedule_id" gorm:"unique;not null;size:8"`
	Description  string                   `json:"description"`
	RequiredDate time.Time                `json:"required_date"`
	SchedularID  uint                     `json:"schedular_id"`
	Schedular    User                     `json:"schedular"`
	ReviewerID   *uint                    `json:"reviewer_id" gorm:"default:null"`
	Reviewer     User                     `json:"reviewer" gorm:"foreignKey:ReviewerID"`
	ReviewedDate *time.Time               `json:"reviewed_date"`
	Status       constants.ScheduleStatus `json:"status"`
	Note         string                   `json:"note"`
}
