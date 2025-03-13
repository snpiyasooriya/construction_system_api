package dto

import "time"

type ProjectCreateInputDTO struct {
	Name      string    `json:"name" binding:"required"`
	ProjectID string    `json:"project_id" binding:"required,min=4,max=8"`
	TypeID    uint      `json:"type_id" binding:"required"`
	LeaderID  uint      `json:"leader_id" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
	Note      string    `json:"note"`
	Status    string    `json:"status" binding:"required,oneof=PENDING IN_PROGRESS COMPLETED CANCELLED"`
	UserIDs   []uint    `json:"user_ids"`
}

type ProjectCreateOutputDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	ProjectID string    `json:"project_id"`
	TypeID    uint      `json:"type_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Note      string    `json:"note"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UserIDs   []uint    `json:"user_ids"`
}

type ProjectsGetDTO struct {
	Projects []ProjectGetDTO
}

type ProjectGetDTO struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	ProjectID     string    `json:"project_id"`
	ProjectType   string    `json:"project_type"`
	LeaderID      uint      `json:"leader_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Note          string    `json:"note"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ScheduleCount int       `json:"schedule_count"`
	UserIDs       []uint    `json:"user_ids"`
}
