package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sivadath/glofox/models"
	"github.com/sivadath/glofox/storage"
	"net/http"
	"time"
)

type BookingController interface {
	CreateBooking(c *gin.Context)
	GetBookings(c *gin.Context)
}

type bookingController struct {
	storage storage.Storage
}

// NewBookingController returns an instance of BookingController
func NewBookingController(s storage.Storage) BookingController {
	return &bookingController{storage: s}
}

// CreateBookingRequest represents the request body for creating a new booking
// @Description Create a new booking with the provided details
// @Tags Booking
// @Accept json
// @Produce json
// @Param request body CreateBookingRequest true "Booking Information"
// @Success 201 {object} models.Booking "Booking created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /bookings [post]
type CreateBookingRequest struct {
	Name string `json:"name" binding:"required"` // @Description Name of the person making the booking
	Date string `json:"date" binding:"required"` // @Description Date for the booking in YYYY-MM-DD format
}

// CreateBooking handles the creation of a new booking
// @Summary Create a new booking
// @Tags Booking
// @Accept json
// @Produce json
// @Param request body CreateBookingRequest true "Booking Information"
// @Success 201 {object} models.Booking "Booking created successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /bookings [post]
func (bc *bookingController) CreateBooking(c *gin.Context) {
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking := models.Booking{
		Name: req.Name,
		Date: req.Date,
	}
	date, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	classes, err := bc.storage.GetClassesByDate(c, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(classes) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// For simplicity considering only first class fetched among all available classes
	booking.ClassID = classes[0].ID

	newBooking, err := bc.storage.AddBooking(c, booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, newBooking)
}

// GetBookings retrieves a list of all bookings
// @Description Get a list of all bookings
// @Tags Booking
// @Accept json
// @Produce json
// @Success 200 {array} models.Booking "List of bookings"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /bookings [get]
func (bc *bookingController) GetBookings(c *gin.Context) {
	bookings, err := bc.storage.GetBookings(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, bookings)
}

// ErrorResponse defines the structure of an error response
// @Description Standard error response structure
// @Tags General
// @Accept json
// @Produce json
// @Property error string "Error message"
// @Property details string "Additional error details" optional
type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}
