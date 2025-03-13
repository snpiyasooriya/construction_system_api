package dto

import "time"

type ProjectTypeCreateDTO struct {
	ID        uint   `json:"id"`
	Type      string `json:"type" binding:"required"`
	CreatedAt time.Time
}

type ProjectTypeGetDto struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
