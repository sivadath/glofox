package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sivadath/glofox/controllers"
	"github.com/sivadath/glofox/storage"
	"path"
)

const (
	Version          = "/api/v1"
	EndPointBookings = "bookings"
)

func RegisterBookingRoutes(r *gin.Engine, db storage.Storage) {
	bc := controllers.NewBookingController(db)
	bookingRoutes := r.Group(path.Join(Version, EndPointBookings))

	bookingRoutes.POST("", bc.CreateBooking)
	bookingRoutes.GET("", bc.GetBookings)
}
