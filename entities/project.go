package entities

import "time"

type Project struct {
	ID            uint
	Name          string
	ProjectID     string
	ProjectTypeID uint
	ProjectType   ProjectType
	Address       string
	LeaderID      uint
	Leader        User
	StartDate     time.Time
	EndDate       time.Time
	Note          string
	Status        string
	Users         []User
	Schedules     []Schedule
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}
