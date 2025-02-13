package models

// Booking represents the structure of a booking
// @Description Booking information
// @Tags Booking
// @Accept json
// @Produce json
// @Property ID string "ID of the booking"
// @Property ClassID string "ID of the associated class"
// @Property Name string "Name of the person who made the booking"
// @Property Date string "Date of the booking in YYYY-MM-DD format"
// @Router /bookings [get]
type Booking struct {
	ID      string `json:"id"`       // @Description ID of the booking
	ClassID string `json:"class_id"` // @Description ID of the associated class
	Name    string `json:"name"`     // @Description Name of the person who made the booking
	Date    string `json:"date"`     // @Description Date of the booking in YYYY-MM-DD format
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
