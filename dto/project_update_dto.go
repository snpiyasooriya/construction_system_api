package dto

import "time"

type ProjectUpdateDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" binding:"required"`
	ProjectID string    `json:"project_id" binding:"required,min=4,max=8"`
	TypeID    uint      `json:"type_id" binding:"required"`
	Address   string    `json:"address"`
	LeaderID  uint      `json:"leader_id"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
	Note      string    `json:"note"`
	Status    string    `json:"status" binding:"required,oneof=PENDING IN_PROGRESS COMPLETED CANCELLED"`
	UpdatedAt time.Time `json:"updated_at"`
	UserIDs   []uint    `json:"user_ids"`
}
