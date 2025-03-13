package migrations

import (
	"gorm.io/gorm"
)

func MigrateProjectFields(db *gorm.DB) error {
	// Drop unique constraint on project_id if it exists
	db.Exec("ALTER TABLE projects DROP CONSTRAINT IF EXISTS uni_projects_project_id")

	// Add new columns
	type Project struct {
		ProjectID string  `gorm:"size:8"`
		StartDate *string `gorm:"type:timestamp"`
		EndDate   *string `gorm:"type:timestamp"`
		Note      *string `gorm:"type:text"`
		Status    string  `gorm:"default:PENDING"`
	}

	if err := db.Migrator().AutoMigrate(&Project{}); err != nil {
		return err
	}

	// Update existing records with default values
	if err := db.Exec(`
		UPDATE projects 
		SET status = 'PENDING',
		    project_id = CONCAT('PRJ', LPAD(CAST(id AS VARCHAR), 5, '0'))
		WHERE project_id IS NULL OR project_id = '';
	`).Error; err != nil {
		return err
	}

	// Add unique constraint on project_id
	if err := db.Exec("ALTER TABLE projects ADD CONSTRAINT uni_projects_project_id UNIQUE (project_id);").Error; err != nil {
		return err
	}

	return nil
}
