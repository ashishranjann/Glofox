package service

import (
	"testing"

	"github.com/ashishranjann/glofox.com/internal/model"
	"github.com/ashishranjann/glofox.com/internal/store"
)

func setupClassTestStore() *store.InMemoryStore {
	return &store.InMemoryStore{
		Classes:        make(map[string]model.Class),
		BookingHistory: []model.BookingRequest{},
	}
}

func TestCreateClass_Success(t *testing.T) {
	store := setupClassTestStore()
	service := NewClassService(store)

	req := &model.ClassRequest{
		Name:       "Pilates",
		Start_date: "2026-12-01",
		End_date:   "2026-12-20",
		Capacity:   20,
	}

	err := service.CreateClass(req)
	if err != nil {
		t.Fatalf("expected success result, got %v", err)
	}

	class, exists := store.Classes["Pilates"]
	if !exists {
		t.Fatal("expected class to be created & stored")
	}

	if class.Name != "Pilates" {
		t.Fatalf("expected name Pilates, got %s", class.Name)
	}

	if class.Capacity != 20 {
		t.Fatalf("expected capacity 20, got %d", class.Capacity)
	}
}

func TestCreateClass_InvalidRequest(t *testing.T) {
	store := setupClassTestStore()
	service := NewClassService(store)

	req := &model.ClassRequest{
		Name:       "",
		Start_date: "2026-12-01",
		End_date:   "2026-12-20",
		Capacity:   0,
	}

	err := service.CreateClass(req)
	if err == nil {
		t.Fatal("expected error for invalid request")
	}
}

func TestCreateClass_InvalidStartDate(t *testing.T) {
	store := setupClassTestStore()
	service := NewClassService(store)

	req := &model.ClassRequest{
		Name:       "Yoga",
		Start_date: "invalid-date",
		End_date:   "2026-12-20",
		Capacity:   10,
	}

	err := service.CreateClass(req)
	if err == nil {
		t.Fatal("expected error for invalid start date")
	}
}

