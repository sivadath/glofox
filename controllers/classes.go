package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sivadath/glofox/internal/errors"
	"github.com/sivadath/glofox/models"
	"github.com/sivadath/glofox/storage"
	"net/http"
)

// ClassController defines handlers for class-related API endpoints.
type ClassController interface {
	CreateClass(c *gin.Context)
	GetClasses(c *gin.Context)
}

// classController implements the ClassController interface.
type classController struct {
	storage storage.Storage
}

// NewClassController initializes and returns an instance of ClassController.
func NewClassController(s storage.Storage) ClassController {
	return &classController{storage: s}
}

// CreateClass handles the creation of a new class.
// @Description Create a new class with the provided details
// @Tags Class
// @Accept json
// @Produce json
// @Param request body models.CreateClassRequest true "Class Information"
// @Success 201 {object} models.Class "Class created successfully"
// @Failure 400 {object} errors.APIError "Invalid input"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Router /classes [post]
func (cc *classController) CreateClass(c *gin.Context) {
	var req models.CreateClassRequest

	// Parse request body and validate JSON format.
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("Failed to parse request body: %v", err)
		errors.ErrInvalidRequest.Respond(c)
		return
	}

	// Ensure the end date is not earlier than the start date.
	if req.EndDate.Time().Before(req.StartDate.Time()) {
		log.Errorf("Invalid class creation request: End date (%s) is before start date (%s)",
			req.EndDate.Time().Format("2006-01-02"), req.StartDate.Time().Format("2006-01-02"))
		errors.ErrDateMismatch.Respond(c)
		return
	}

	newClass := models.Class{
		Name:      req.Name,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Capacity:  req.Capacity,
	}

	// Store the class in the database.
	storedClass, err := cc.storage.AddClass(c, newClass)
	if err != nil {
		log.Errorf("Database insertion error: %v", err)
		errors.NewError(err.Error(), http.StatusInternalServerError).Respond(c)
		return
	}

	log.Infof("Class successfully created: %+v", storedClass)
	c.JSON(http.StatusCreated, storedClass)
}

// GetClasses retrieves a list of all available classes.
// @Description Get a list of all classes
// @Tags Class
// @Accept json
// @Produce json
// @Success 200 {array} models.Class "List of classes"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Router /classes [get]
func (cc *classController) GetClasses(c *gin.Context) {
	// Fetch class list from the database.
	classes, err := cc.storage.GetClasses(c)
	if err != nil {
		log.Errorf("Failed to retrieve class list: %v", err)
		errors.ErrInternalServer.Respond(c)
		return
	}

	log.Infof("Returning %d available classes", len(classes))
	c.JSON(http.StatusOK, classes)
}
