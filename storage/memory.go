package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/sivadath/glofox/models"
	"time"
)

type InMemoryStorage struct {
	classes  []models.Class
	bookings []models.Booking
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		classes:  []models.Class{},
		bookings: []models.Booking{},
	}
}

func (s *InMemoryStorage) AddClass(ctx context.Context, class models.Class) (models.Class, error) {
	class.ID = uuid.NewString()
	s.classes = append(s.classes, class)
	return class, nil
}

func (s *InMemoryStorage) GetClasses(ctx context.Context) ([]models.Class, error) {
	return s.classes, nil
}

func (s *InMemoryStorage) GetClassesByDate(ctx context.Context, date time.Time) (classes []models.Class, err error) {
	for _, class := range s.classes {
		if date.After(class.StartDate.Time()) && date.Before(class.EndDate.Time()) {
			classes = append(classes, class)
			break
		}
	}
	return
}

func (s *InMemoryStorage) AddBooking(ctx context.Context, booking models.Booking) (models.Booking, error) {
	booking.ID = uuid.NewString()
	s.bookings = append(s.bookings, booking)
	return booking, nil
}

func (s *InMemoryStorage) GetBookings(ctx context.Context) ([]models.Booking, error) {
	return s.bookings, nil
}
