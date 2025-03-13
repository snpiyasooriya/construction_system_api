package repository

import (
	"fmt"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

type GORMProjectRepository struct {
	db *gorm.DB
}

func NewGORMProjectRepository(db *gorm.DB) *GORMProjectRepository {
	return &GORMProjectRepository{
		db: db,
	}
}

func (G *GORMProjectRepository) Create(input dto.ProjectCreateInputDTO) (*dto.ProjectCreateOutputDTO, error) {
	// Start a transaction
	tx := G.db.Begin()

	// Create project
	createdProject := &models.Project{
		Model:         gorm.Model{},
		Name:          input.Name,
		ProjectID:     input.ProjectID,
		ProjectTypeID: input.TypeID,
		LeaderID:      input.LeaderID,
		StartDate:     input.StartDate,
		EndDate:       input.EndDate,
		Note:          input.Note,
		Status:        input.Status,
	}

	// Create the project
	if err := tx.Create(&createdProject).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Add users if provided
	var userIDs []uint
	if len(input.UserIDs) > 0 {
		// Get users
		var users []models.User
		if err := tx.Find(&users, input.UserIDs).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Check if all users were found
		if len(users) != len(input.UserIDs) {
			tx.Rollback()
			return nil, fmt.Errorf("some users not found")
		}

		// Add users to project
		if err := tx.Model(&createdProject).Association("Users").Append(&users); err != nil {
			tx.Rollback()
			return nil, err
		}

		userIDs = input.UserIDs
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	projectCreateOutput := dto.ProjectCreateOutputDTO{
		ID:        createdProject.ID,
		Name:      createdProject.Name,
		ProjectID: createdProject.ProjectID,
		TypeID:    createdProject.ProjectTypeID,
		StartDate: createdProject.StartDate,
		EndDate:   createdProject.EndDate,
		Note:      createdProject.Note,
		Status:    createdProject.Status,
		CreatedAt: createdProject.CreatedAt,
		UserIDs:   userIDs,
	}

	return &projectCreateOutput, nil
}

func (G *GORMProjectRepository) UpdateByID(input models.Project) (*models.Project, error) {
	// Start a transaction
	tx := G.db.Begin()

	// Get existing project
	var project models.Project
	if err := tx.First(&project, input.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update project fields
	project.Name = input.Name
	project.ProjectID = input.ProjectID
	project.ProjectTypeID = input.ProjectTypeID
	project.LeaderID = input.LeaderID
	project.Address = input.Address
	project.StartDate = input.StartDate
	project.EndDate = input.EndDate
	project.Note = input.Note
	project.Status = input.Status

	// Save project updates
	if err := tx.Save(&project).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update users if provided
	if len(input.Users) > 0 {
		// Get users
		var users []models.User
		userIDs := make([]uint, len(input.Users))
		for i, user := range input.Users {
			userIDs[i] = user.ID
		}

		if err := tx.Find(&users, userIDs).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Check if all users were found
		if len(users) != len(input.Users) {
			tx.Rollback()
			return nil, fmt.Errorf("some users not found")
		}

		// Replace all users
		if err := tx.Model(&project).Association("Users").Replace(&users); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Reload project with associations
	if err := G.db.Preload("Users").First(&project, project.ID).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

func (G *GORMProjectRepository) DeleteByID(id uint) error {
	result := G.db.Delete(&models.Project{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (G *GORMProjectRepository) GetByID(id uint) (*models.Project, error) {
	var project models.Project
	if err := G.db.Preload("ProjectType").Preload("Schedules").Preload("Users").First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (G *GORMProjectRepository) Get() (*dto.ProjectsGetDTO, error) {
	var projects []models.Project
	var projectsDTO dto.ProjectsGetDTO

	if err := G.db.Preload("ProjectType").Preload("Schedules").Preload("Users").Find(&projects).Error; err != nil {
		return nil, err
	}
	for _, project := range projects {
		// Get user IDs
		userIDs := make([]uint, len(project.Users))
		for i, user := range project.Users {
			userIDs[i] = user.ID
		}

		projectDTO := dto.ProjectGetDTO{
			ID:            project.ID,
			Name:          project.Name,
			ProjectID:     project.ProjectID,
			ProjectType:   project.ProjectType.Type,
			LeaderID:      project.LeaderID,
			StartDate:     project.StartDate,
			EndDate:       project.EndDate,
			Note:          project.Note,
			Status:        project.Status,
			CreatedAt:     project.CreatedAt,
			UpdatedAt:     project.UpdatedAt,
			ScheduleCount: len(project.Schedules),
			UserIDs:       userIDs,
		}
		projectsDTO.Projects = append(projectsDTO.Projects, projectDTO)
	}
	return &projectsDTO, nil
}

func (G *GORMProjectRepository) AddUser(input dto.ProjectAddUserDTO) (*dto.ProjectAddUserOutputDTO, error) {
	// Get the project
	var project models.Project
	if err := G.db.First(&project, input.ProjectID).Error; err != nil {
		return nil, err
	}

	// Get the user
	var user models.User
	if err := G.db.First(&user, input.UserID).Error; err != nil {
		return nil, err
	}

	// Add user to project
	if err := G.db.Model(&project).Association("Users").Append(&user); err != nil {
		return nil, err
	}

	return &dto.ProjectAddUserOutputDTO{
		ProjectID: project.ID,
		UserID:    user.ID,
	}, nil
}
