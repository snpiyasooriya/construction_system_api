package migrations

import (
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Project{}, &models.ProjectType{}, &models.Schedule{}, &models.ScheduleItemCrome{})
	if err != nil {
		panic("failed to migrate database")
	}
}
