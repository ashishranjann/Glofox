package service

import (
	"errors"

	"github.com/ashishranjann/glofox.com/internal/model"
	"github.com/ashishranjann/glofox.com/internal/store"
	"github.com/ashishranjann/glofox.com/internal/utils"
)

type BookingService struct {
	store *store.InMemoryStore
}

func NewBookingService(s *store.InMemoryStore) *BookingService{
	return &BookingService{store: s}
}

func (s *BookingService)CreateBooking(req *model.BookingRequest) error {
	if req.Name == "" || req.Class_name == "" {
		return errors.New("Invalid request")
	}

	date, err := utils.ParseDate(req.Date)
	if err != nil {
		return err
	}

	s.store.Mu.RLock()
	class, exists := s.store.Classes[req.Class_name]
	s.store.Mu.RUnlock()

	if !exists{
		return errors.New("class not found")
	}

	if date.Before(class.Start_date) || date.After(class.End_date) {
		return errors.New("Date invalid for the requested class")
	}

	s.store.Mu.Lock()
	defer s.store.Mu.Unlock()

	s.store.BookingHistory = append(s.store.BookingHistory, *req)
	return  nil
}
