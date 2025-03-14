package migrations

import (
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	//// For production deployment, we'll use a simpler approach to avoid migration issues
	//// First, check if tables already exist to determine if this is a fresh install
	//var tableCount int64
	//db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tableCount)
	//
	//// If we have tables already, skip the migration to avoid disrupting the production database
	//if tableCount > 0 {
	//	fmt.Println("Database tables already exist, skipping migration to avoid disrupting production data")
	//	return
	//}
	//
	//// For a fresh install, create tables in the correct order
	//fmt.Println("Fresh database detected, performing initial migration")
	//
	//// First create the users table
	//db.Exec(`CREATE TABLE IF NOT EXISTS users (
	//	id SERIAL PRIMARY KEY,
	//	created_at TIMESTAMP WITH TIME ZONE,
	//	updated_at TIMESTAMP WITH TIME ZONE,
	//	deleted_at TIMESTAMP WITH TIME ZONE,
	//	name TEXT,
	//	email TEXT UNIQUE,
	//	password TEXT,
	//	role TEXT
	//);`)
	//
	//// Create the project_types table
	//db.Exec(`CREATE TABLE IF NOT EXISTS project_types (
	//	id SERIAL PRIMARY KEY,
	//	created_at TIMESTAMP WITH TIME ZONE,
	//	updated_at TIMESTAMP WITH TIME ZONE,
	//	deleted_at TIMESTAMP WITH TIME ZONE,
	//	type TEXT
	//);`)
	//
	//// Create the projects table
	//db.Exec(`CREATE TABLE IF NOT EXISTS projects (
	//	id SERIAL PRIMARY KEY,
	//	created_at TIMESTAMP WITH TIME ZONE,
	//	updated_at TIMESTAMP WITH TIME ZONE,
	//	deleted_at TIMESTAMP WITH TIME ZONE,
	//	name TEXT,
	//	project_id TEXT UNIQUE,
	//	project_type_id INTEGER,
	//	leader_id INTEGER,
	//	address TEXT,
	//	start_date TIMESTAMP WITH TIME ZONE,
	//	end_date TIMESTAMP WITH TIME ZONE,
	//	note TEXT,
	//	status TEXT DEFAULT 'PENDING'
	//);`)
	//
	//// Create the project_users join table
	//db.Exec(`CREATE TABLE IF NOT EXISTS project_users (
	//	user_id INTEGER,
	//	project_id INTEGER,
	//	PRIMARY KEY (user_id, project_id)
	//);`)
	//
	//// Create the schedules table
	//db.Exec(`CREATE TABLE IF NOT EXISTS schedules (
	//	id SERIAL PRIMARY KEY,
	//	created_at TIMESTAMP WITH TIME ZONE,
	//	updated_at TIMESTAMP WITH TIME ZONE,
	//	deleted_at TIMESTAMP WITH TIME ZONE,
	//	name TEXT,
	//	description TEXT,
	//	project_id INTEGER
	//);`)
	//
	//// Create the schedule_item_cromes table
	//db.Exec(`CREATE TABLE IF NOT EXISTS schedule_item_cromes (
	//	id SERIAL PRIMARY KEY,
	//	created_at TIMESTAMP WITH TIME ZONE,
	//	updated_at TIMESTAMP WITH TIME ZONE,
	//	deleted_at TIMESTAMP WITH TIME ZONE,
	//	name TEXT,
	//	description TEXT,
	//	schedule_id INTEGER
	//);`)
	//
	//// Add foreign key constraints after all tables are created
	//db.Exec(`ALTER TABLE projects ADD CONSTRAINT fk_projects_leader FOREIGN KEY (leader_id) REFERENCES users(id);`)
	//db.Exec(`ALTER TABLE project_users ADD CONSTRAINT fk_project_users_user FOREIGN KEY (user_id) REFERENCES users(id);`)
	//db.Exec(`ALTER TABLE project_users ADD CONSTRAINT fk_project_users_project FOREIGN KEY (project_id) REFERENCES projects(id);`)
	//db.Exec(`ALTER TABLE schedules ADD CONSTRAINT fk_schedules_project FOREIGN KEY (project_id) REFERENCES projects(id);`)
	//db.Exec(`ALTER TABLE schedule_item_cromes ADD CONSTRAINT fk_schedule_item_cromes_schedule FOREIGN KEY (schedule_id) REFERENCES schedules(id);`)
	//
	//fmt.Println("Initial database migration completed successfully")

	err := db.AutoMigrate(&models.User{}, &models.Project{}, &models.ProjectType{})
	if err != nil {
		panic("failed to migrate database")
	}

}
