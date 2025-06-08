package repository

import (
	"github.com/snpiyasooriya/construction_design_api/entities"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

type GORMShapeRepository struct {
	db *gorm.DB
}

// NewGORMShapeRepository creates a new instance of GORMShapeRepository
func NewGORMShapeRepository(db *gorm.DB) repository.ShapeRepositoryInterFace {
	return &GORMShapeRepository{db: db}
}

// Create creates a new shape
func (g *GORMShapeRepository) Create(shape entities.Shape) (*entities.Shape, error) {
	// Convert entity to model
	var shapeModel models.Shape
	shapeModel.FromEntity(shape)

	// Save to database
	if err := g.db.Create(&shapeModel).Error; err != nil {
		return nil, err
	}

	// Convert back to entity
	result := shapeModel.ToEntity()
	return &result, nil
}

// GetByID retrieves a shape by ID
func (g *GORMShapeRepository) GetByID(id uint) (*entities.Shape, error) {
	var shape models.Shape
	if err := g.db.First(&shape, id).Error; err != nil {
		return nil, err
	}

	result := shape.ToEntity()
	return &result, nil
}

// Get retrieves all shapes
func (g *GORMShapeRepository) Get() ([]entities.Shape, error) {
	var shapes []models.Shape
	if err := g.db.Find(&shapes).Error; err != nil {
		return nil, err
	}

	// Convert to entities
	result := make([]entities.Shape, len(shapes))
	for i, shape := range shapes {
		result[i] = shape.ToEntity()
	}

	return result, nil
}

// DeleteByID deletes a shape by ID
func (g *GORMShapeRepository) DeleteByID(id uint) error {
	result := g.db.Delete(&models.Shape{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
