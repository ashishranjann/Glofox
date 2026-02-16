package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ashishranjann/glofox.com/internal/model"
	"github.com/ashishranjann/glofox.com/internal/service"
)

type BookingHandler struct {
	service *service.BookingService
}

func NewBookinghandler(s *service.BookingService) *BookingHandler {
	return &BookingHandler{service: s}
}

func (h *BookingHandler) CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req model.BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "inValid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateBooking(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("message: Booking completed successfully")
}
