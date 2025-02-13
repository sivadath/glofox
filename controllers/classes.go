package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sivadath/glofox/internal/errors"
	"github.com/sivadath/glofox/models"
	"github.com/sivadath/glofox/storage"
	"log"
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
		log.Printf("Invalid request, missing request params: %v\n", err)
		errors.ErrInvalidRequest.Respond(c)
		return
	}
	if req.EndDate.Time().Before(req.StartDate.Time()) {
		log.Printf("Invalid request: %v\n", errors.ErrDateMismatch.Message)
		errors.ErrDateMismatch.Respond(c)
		return
	}
	newClass, err := cc.storage.AddClass(c, models.Class{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Capacity:  req.Capacity,
	})
	if err != nil {
		log.Printf("Failed inserting class to db: %v\n", err)
		errors.NewError(err.Error(), http.StatusInternalServerError).Respond(c)
		return
	}
	log.Printf("Class created successfully: %v\n", newClass)
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
		errors.ErrInternalServer.Respond(c)
		return
	}
	log.Printf("Returning available classes\n")
	c.JSON(http.StatusOK, classes)
}
