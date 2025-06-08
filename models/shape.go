package models

import (
	"github.com/snpiyasooriya/construction_design_api/entities"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Shape struct {
	gorm.Model
	Name       string         `json:"name"`
	Path       string         `json:"path"`
	Dimensions datatypes.JSON `json:"dimensions" gorm:"type:json"`
}

func (s *Shape) ToEntity() entities.Shape {
	// Convert datatypes.JSON to map[string]interface{}
	dimensions := utils.DataTypeJsonToInterface(s.Dimensions)
	return entities.Shape{
		ID:         s.ID,
		Name:       s.Name,
		Path:       s.Path,
		Dimensions: dimensions,
	}
}

func (s *Shape) FromEntity(shape entities.Shape) {
	// Convert map[string]interface{} to datatypes.JSON
	dimensions := utils.InterfaceToDataTypeJson(shape.Dimensions)
	s.ID = shape.ID
	s.Name = shape.Name
	s.Path = shape.Path
	s.Dimensions = dimensions
}
