package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/sivadath/glofox/models"
	"sync"
	"time"
)

type InMemoryStorage struct {
	classLock    sync.RWMutex
	bookingsLock sync.RWMutex
	classes      []models.Class
	bookings     []models.Booking
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		classes:  []models.Class{},
		bookings: []models.Booking{},
	}
}

func (s *InMemoryStorage) AddClass(ctx context.Context, class models.Class) (models.Class, error) {
	class.ID = uuid.NewString()
	s.classLock.Lock()
	// This will cause a panic if the number of bookings exceeds the maximum value of an int (INTMAX),
	// typically due to excessive slice growth during append operations.
	s.classes = append(s.classes, class)
	s.classLock.Unlock()
	return class, nil
}

func (s *InMemoryStorage) GetClasses(ctx context.Context) ([]models.Class, error) {
	s.classLock.RLock()
	classes := make([]models.Class, len(s.classes))
	copy(classes, s.classes)
	s.classLock.RUnlock()
	return classes, nil
}

func (s *InMemoryStorage) GetClassesByDate(ctx context.Context, date time.Time) (classes []models.Class, err error) {
	s.classLock.RLock()
	for _, class := range s.classes {
		if date.After(class.StartDate.Time()) && date.Before(class.EndDate.Time()) {
			classes = append(classes, class)
			break
		}
	}
	s.classLock.RUnlock()
	return
}

func (s *InMemoryStorage) AddBooking(ctx context.Context, booking models.Booking) (models.Booking, error) {
	booking.ID = uuid.NewString()
	s.bookingsLock.Lock()
	// This will cause a panic if the number of bookings exceeds the maximum value of an int (INTMAX),
	// typically due to excessive slice growth during append operations.
	s.bookings = append(s.bookings, booking)
	s.bookingsLock.Unlock()
	return booking, nil
}

func (s *InMemoryStorage) GetBookings(ctx context.Context) ([]models.Booking, error) {
	s.bookingsLock.RLock()
	bookings := make([]models.Booking, len(s.bookings))
	copy(bookings, s.bookings)
	s.bookingsLock.RUnlock()
	return bookings, nil
}
