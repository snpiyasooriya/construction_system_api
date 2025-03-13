package migrations

import (
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	// Run project fields migration first
	if err := MigrateProjectFields(db); err != nil {
		panic("failed to migrate project fields: " + err.Error())
	}

	// Run auto migrations for all models
	if err := db.AutoMigrate(&models.User{}, &models.Project{}, &models.ProjectType{}, &models.Schedule{}, &models.ScheduleItemCrome{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
