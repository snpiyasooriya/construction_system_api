package models

import (
	entities2 "github.com/snpiyasooriya/construction_design_api/entities"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type Project struct {
	gorm.Model
	Name          string     `gorm:"unique;not null"`
	ProjectID     string     `gorm:"unique;not null;size:8"`
	ProjectTypeID uint       `gorm:"not null"`
	ProjectType   ProjectType
	Schedules     []Schedule
	Users         []User `gorm:"many2many:project_users;"`
	LeaderID      uint
	Leader        User `gorm:"foreignKey:LeaderID"`
	Address       string
	StartDate     time.Time
	EndDate       time.Time
	Note          string     `gorm:"type:text"`
	Status        string     `gorm:"not null;default:'PENDING'"`
}

func (p *Project) ToEntity() entities2.Project {
	return entities2.Project{
		ID:          p.ID,
		Name:        p.Name,
		ProjectID:   p.ProjectID,
		ProjectType: p.ProjectType.ToEntity(),
		LeaderID:    p.LeaderID,
		Leader:      p.Leader.ToEntity(),
		Address:     p.Address,
		StartDate:   p.StartDate,
		EndDate:     p.EndDate,
		Note:        p.Note,
		Status:      p.Status,
		Users: func() []entities2.User {
			users := make([]entities2.User, len(p.Users))
			for i, u := range p.Users {
				users[i] = u.ToEntity()
			}
			return users
		}(),
		Schedules: []entities2.Schedule{},
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: utils.ConvertDeletedAtToTime(p.DeletedAt),
	}
}
