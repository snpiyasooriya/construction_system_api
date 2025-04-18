package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"github.com/snpiyasooriya/construction_design_api/services"
	"net/http"
	"strconv"
)

// ShapeController handles HTTP requests related to shapes
type ShapeController struct {
	shapeService *services.ShapeService
}

// NewShapeController creates a new instance of ShapeController
func NewShapeController(shapeService *services.ShapeService) *ShapeController {
	return &ShapeController{
		shapeService: shapeService,
	}
}

// Create handles the creation of a new shape
// @Summary Create a new shape
// @Description Create a new shape with the provided details
// @Tags shapes
// @Accept json
// @Produce json
// @Param shape body dto.ShapeDTO true "Shape creation data"
// @Success 201 {object} dto.ShapeDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/shapes [post]
func (sc *ShapeController) Create(c *gin.Context) {
	var shapeDTO dto.ShapeDTO
	if err := c.ShouldBindJSON(&shapeDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&shapeDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	// Create shape - the DTO will be updated with the created shape's data
	err := sc.shapeService.CreateShape(&shapeDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, shapeDTO)
}

// Get handles the retrieval of all shapes
// @Summary Get all shapes
// @Description Get a list of all shapes
// @Tags shapes
// @Produce json
// @Success 200 {object} dto.ShapesGetDTO
// @Failure 500 {object} map[string]interface{}
// @Router /api/shapes [get]
func (sc *ShapeController) Get(c *gin.Context) {
	shapes, err := sc.shapeService.GetShapes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shapes": shapes})
}

// GetByID handles the retrieval of a shape by ID
// @Summary Get a shape by ID
// @Description Get a shape's details by its ID
// @Tags shapes
// @Produce json
// @Param id path int true "Shape ID"
// @Success 200 {object} dto.ShapeDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/shapes/{id} [get]
func (sc *ShapeController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shape ID"})
		return
	}

	shape, err := sc.shapeService.GetShapeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shape not found"})
		return
	}

	c.JSON(http.StatusOK, shape)
}

// Delete handles the deletion of a shape by ID
// @Summary Delete a shape
// @Description Delete a shape by its ID
// @Tags shapes
// @Param id path int true "Shape ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/shapes/{id} [delete]
func (sc *ShapeController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shape ID"})
		return
	}

	err = sc.shapeService.DeleteShape(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shape deleted successfully"})
}
