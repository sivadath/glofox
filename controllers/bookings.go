package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sivadath/glofox/internal/errors"
	"github.com/sivadath/glofox/models"
	"github.com/sivadath/glofox/storage"
	"net/http"
	"time"
)

// BookingController defines handlers for booking-related endpoints.
type BookingController interface {
	CreateBooking(c *gin.Context)
	GetBookings(c *gin.Context)
}

type bookingController struct {
	storage storage.Storage
}

// NewBookingController initializes a new booking controller.
func NewBookingController(s storage.Storage) BookingController {
	return &bookingController{storage: s}
}

// CreateBooking handles the creation of a new booking.
// @Summary Create a new booking
// @Tags Booking
// @Accept json
// @Produce json
// @Param request body models.CreateBookingRequest true "Booking Information"
// @Success 201 {object} models.Booking "Booking created successfully"
// @Failure 400 {object} errors.APIError "Invalid input"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Router /bookings [post]
func (bc *bookingController) CreateBooking(c *gin.Context) {
	var req models.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("Invalid request body: %v", err)
		errors.ErrInvalidRequest.Respond(c)
		return
	}

	// Parse the date string into a standard format
	date, err := time.Parse(time.DateOnly, req.Date)
	if err != nil {
		log.Errorf("Invalid date format (%s), expected YYYY-MM-DD: %v", req.Date, err)
		errors.NewError("Invalid date format, expected YYYY-MM-DD", http.StatusBadRequest).Respond(c)
		return
	}

	// Retrieve available classes for the given date
	classes, err := bc.storage.GetClassesByDate(c, date)
	if err != nil {
		log.Errorf("Database error while retrieving classes for date %s: %v", date, err)
		errors.NewError(err.Error(), http.StatusInternalServerError).Respond(c)
		return
	}

	// If no classes are available, return an appropriate response
	if len(classes) == 0 {
		log.Errorf("No available classes for the given date: %s", date)
		errors.ErrNoClassesFound.Respond(c)
		return
	}

	// Assign the first available class (simplified logic for now)
	booking := models.Booking{
		Name:    req.Name,
		Date:    req.Date,
		ClassID: classes[0].ID,
	}

	// Store the new booking in the database
	newBooking, err := bc.storage.AddBooking(c, booking)
	if err != nil {
		log.Errorf("Failed to insert booking into database: %v", err)
		errors.NewError(err.Error(), http.StatusInternalServerError).Respond(c)
		return
	}

	log.Infof("Booking successfully created: %+v", newBooking)
	c.JSON(http.StatusCreated, newBooking)
}

// GetBookings retrieves all existing bookings.
// @Description Get a list of all bookings
// @Tags Booking
// @Accept json
// @Produce json
// @Success 200 {array} models.Booking "List of bookings"
// @Failure 500 {object} errors.APIError "Internal server error"
// @Router /bookings [get]
func (bc *bookingController) GetBookings(c *gin.Context) {
	// Fetch all bookings from storage
	bookings, err := bc.storage.GetBookings(c)
	if err != nil {
		log.Errorf("Database error while retrieving bookings: %v", err)
		errors.NewError(err.Error(), http.StatusInternalServerError).Respond(c)
		return
	}

	log.Infof("Returning %d bookings", len(bookings))
	c.JSON(http.StatusOK, bookings)
}
