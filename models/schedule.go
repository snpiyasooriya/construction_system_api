package models

import (
	"github.com/snpiyasooriya/construction_design_api/entities"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ProjectID   uint    `json:"project_id"`
	Project     Project `json:"project"`
}

func (s *Schedule) ToEntity() entities.Schedule {
	return entities.Schedule{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Project:     s.Project.ToEntity(),
	}
}
