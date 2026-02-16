package service

import (
	"testing"
	"time"

	"github.com/ashishranjann/glofox.com/internal/model"
	"github.com/ashishranjann/glofox.com/internal/store"
)

func setupTestStore() *store.InMemoryStore {
	return &store.InMemoryStore{
		Classes:        make(map[string]model.Class),
		BookingHistory: []model.BookingRequest{},
	}
}

func TestCreateBooking_Success(t *testing.T) {
	store := setupTestStore()

	store.Classes["Pilates"] = model.Class{
		Name:       "Pilates",
		Start_date: time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC),
		End_date:   time.Date(2026, 12, 20, 0, 0, 0, 0, time.UTC),
		Capacity:   20,
	}

	service := NewBookingService(store)

	req := &model.BookingRequest{
		Name:       "Ashish",
		Class_name: "Pilates",
		Date:       "2026-12-16",
	}

	err := service.CreateBooking(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(store.BookingHistory) != 1 {
		t.Fatalf("expected 1 booking, got %d", len(store.BookingHistory))
	}
}

func TestCreateBooking_InvalidRequest(t *testing.T) {
	store := setupTestStore()
	service := NewBookingService(store)

	req := &model.BookingRequest{
		Name:       "",
		Class_name: "",
		Date:       "2026-12-16",
	}

	err := service.CreateBooking(req)
	if err == nil {
		t.Fatal("expected error for invalid request")
	}
}

func TestCreateBooking_ClassNotFound(t *testing.T) {
	store := setupTestStore()
	service := NewBookingService(store)

	req := &model.BookingRequest{
		Name:       "Ashish",
		Class_name: "Unknown",
		Date:       "2026-12-16",
	}

	err := service.CreateBooking(req)
	if err == nil {
		t.Fatal("expected error for class not found")
	}
}

func TestCreateBooking_DateOutsideRange(t *testing.T) {
	store := setupTestStore()

	store.Classes["Pilates"] = model.Class{
		Name:       "Pilates",
		Start_date: time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC),
		End_date:   time.Date(2026, 12, 20, 0, 0, 0, 0, time.UTC),
		Capacity:   20,
	}

	service := NewBookingService(store)

	req := &model.BookingRequest{
		Name:       "Ashish",
		Class_name: "Pilates",
		Date:       "2026-12-25",
	}

	err := service.CreateBooking(req)
	if err == nil {
		t.Fatal("expected error for date outside class range")
	}
}
