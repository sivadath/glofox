package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrInvalidRequest   = NewError("invalid request data", http.StatusBadRequest)
	ErrMissingName      = NewError("name is required", http.StatusBadRequest)
	ErrMissingStartDate = NewError("start_date is required", http.StatusBadRequest)
	ErrMissingEndDate   = NewError("end_date is required", http.StatusBadRequest)
	ErrInvalidCapacity  = NewError("capacity must be greater than zero", http.StatusBadRequest)
	ErrDateMismatch     = NewError("end_date cannot be before start_date", http.StatusBadRequest)
	ErrInternalServer   = NewError("internal server error", http.StatusInternalServerError)
	ErrNoClassesFound   = NewError("no classes found for given request", http.StatusInternalServerError)
)

// APIError represents a standard error response.
type APIError struct {
	Message string `json:"error"`
	Code    int    `json:"-"`
}

// NewError creates a new API error.
func NewError(message string, code int) *APIError {
	return &APIError{Message: message, Code: code}
}

// Respond sends an error response.
func (e *APIError) Respond(c *gin.Context) {
	c.JSON(e.Code, gin.H{"error": e.Message})
}
