package entities

import "time"

type ScheduleItemCrome struct {
	ID         uint
	ScheduleID uint
	Name       string
	ShapeID    Shape
	Shape      Shape
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
