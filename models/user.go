package models

import (
	"github.com/snpiyasooriya/construction_design_api/entities"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string `gorm:"unique"`
	Phone        string `gorm:"unique"`
	DOB          time.Time
	NIC          string `gorm:"unique"`
	EmployeeID   string `gorm:"unique"`
	Password     string
	Role         string
	Projects     []Project `gorm:"many2many:project_users;"`
	LeadProjects []Project `gorm:"foreignKey:LeaderID"`
}

func (u *User) ToEntity() entities.User {
	return entities.User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Email:      u.Email,
		Phone:      u.Phone,
		DOB:        u.DOB,
		NIC:        u.NIC,
		EmployeeID: u.EmployeeID,
		Password:   u.Password,
		Role:       u.Role,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
		DeletedAt:  u.DeletedAt.Time,
	}
}

func FromUserEntity(e entities.User) User {
	return User{
		Model: gorm.Model{
			ID:        e.ID,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
			DeletedAt: utils.ConvertTimeToDeletedAt(e.DeletedAt),
		},
		FirstName:  e.FirstName,
		LastName:   e.LastName,
		Email:      e.Email,
		Phone:      e.Phone,
		DOB:        e.DOB,
		NIC:        e.NIC,
		EmployeeID: e.EmployeeID,
		Password:   e.Password,
		Role:       e.Role,
	}
}
