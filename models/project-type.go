package models

import (
	entities2 "github.com/snpiyasooriya/construction_design_api/entities"
	"gorm.io/gorm"
)

type ProjectType struct {
	gorm.Model
	Type     string `gorm:"type:varchar(20);not null;unique"`
	Projects []Project
}

func (p *ProjectType) ToEntity() entities2.ProjectType {
	return entities2.ProjectType{
		ID:   p.ID,
		Type: p.Type,
		Projects: func() []entities2.Project {
			projects := make([]entities2.Project, len(p.Projects))
			for i, proj := range p.Projects {
				projects[i] = proj.ToEntity()
			}
			return projects
		}(),
	}
}
