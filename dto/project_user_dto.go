package dto

type ProjectAddUserDTO struct {
	ProjectID uint   `json:"project_id" binding:"required"`
	UserID    uint   `json:"user_id" binding:"required"`
}

type ProjectAddUserOutputDTO struct {
	ProjectID uint `json:"project_id"`
	UserID    uint `json:"user_id"`
}
