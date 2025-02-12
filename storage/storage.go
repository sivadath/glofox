package storage

import (
	"context"
	"github.com/sivadath/glofox/models"
	"time"
)

type Storage interface {
	ClassSchema
	BookingSchema
}

type ClassSchema interface {
	AddClass(ctx context.Context, class models.Class) (models.Class, error)
	GetClasses(ctx context.Context) ([]models.Class, error)
	GetClassesByDate(ctx context.Context, date time.Time) ([]models.Class, error)
}

type BookingSchema interface {
	AddBooking(ctx context.Context, booking models.Booking) (models.Booking, error)
	GetBookings(ctx context.Context) ([]models.Booking, error)
}

var DB Storage

func SetStorage(s Storage) {
	if DB == nil {
		DB = s
	}
}
