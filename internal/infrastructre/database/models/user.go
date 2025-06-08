package models

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/entities"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Phone     string `gorm:"unique"`
	DOB       time.Time
	NIC       string `gorm:"unique"`
	Password  string
	Role      string
}

func (u *User) ToEntity() entities.User {
	return entities.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
		DOB:       u.DOB,
		NIC:       u.NIC,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt.Time,
	}
}

func FromUserEntity(e entities.User) User {
	return User{
		Model: gorm.Model{
			ID:        e.ID,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
			DeletedAt: ConvertTimeToDeletedAt(e.DeletedAt),
		},
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Email:     e.Email,
		Phone:     e.Phone,
		DOB:       e.DOB,
		NIC:       e.NIC,
		Password:  e.Password,
		Role:      e.Role,
	}
}

// ConvertTimeToDeletedAt converts a time.Time to gorm.DeletedAt.
func ConvertTimeToDeletedAt(t time.Time) gorm.DeletedAt {
	return gorm.DeletedAt{
		Time:  t,
		Valid: !t.IsZero(), // Set Valid to true if t is not the zero time
	}
}
