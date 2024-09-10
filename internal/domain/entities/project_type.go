package entities

import "time"

type ProjectType struct {
	ID        uint
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
