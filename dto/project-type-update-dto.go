package dto

import "time"

type ProjectTypeUpdateDTO struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type" binding:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}
