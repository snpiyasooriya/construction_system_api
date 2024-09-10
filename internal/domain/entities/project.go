package entities

import "time"

type Project struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Name          string
	ID            uint
	ProjectTypeID uint
	ProjectType   ProjectType
}
