package migrations

import (
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Project{}, &models.ProjectType{})
	if err != nil {
		panic("failed to migrate database")
	}
}
