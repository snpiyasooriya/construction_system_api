package entities

import "time"

type Project struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Name          string
	ID            uint
	ProjectID     string
	ProjectTypeID uint
	ProjectType   ProjectType
	Address       string
	LeaderID      uint
	Leader        User
	Users         []User
	Schedules     []Schedule
}
