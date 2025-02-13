package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sivadath/glofox/controllers"
	"github.com/sivadath/glofox/storage"
)

const Version = "/api/v1"

func RegisterBookingRoutes(r *gin.Engine) {
	bc := controllers.NewBookingController(storage.DB)
	bookingRoutes := r.Group(Version + "/bookings")
	bookingRoutes.POST("", bc.CreateBooking)
	bookingRoutes.GET("", bc.GetBookings)
}
