package migrations

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/entities"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&entities.User{})
	if err != nil {
		panic("failed to migrate database")
	}
}
