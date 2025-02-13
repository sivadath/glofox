package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sivadath/glofox/models"
	"github.com/sivadath/glofox/storage"
	"net/http"
)

// ClassController defines the interface
type ClassController interface {
	CreateClass(c *gin.Context)
	GetClasses(c *gin.Context)
}

// classController implements the ClassController interface
type classController struct {
	storage storage.Storage
}

// NewClassController returns an instance of ClassController
func NewClassController(s storage.Storage) ClassController {
	return &classController{storage: s}
}

// CreateClassRequest represents the request body for creating a new class
// @Description Create a new class with the provided details
// @Tags Class
// @Accept json
// @Produce json
// @Param request body models.CreateClassRequest true "Class Information"
// @Success 201 {object} models.Class "Class created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /classes [post]
func (cc *classController) CreateClass(c *gin.Context) {
	var req models.CreateClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.EndDate.Time().Before(req.StartDate.Time()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end time cannot be later to start time"})
		return
	}
	newClass, err := cc.storage.AddClass(c, models.Class{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Capacity:  req.Capacity,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, newClass)
}

// GetClasses retrieves a list of all classes
// @Description Get a list of all classes
// @Tags Class
// @Accept json
// @Produce json
// @Success 200 {array} models.Class "List of classes"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /classes [get]
func (cc *classController) GetClasses(c *gin.Context) {
	classes, err := cc.storage.GetClasses(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, classes)
}
