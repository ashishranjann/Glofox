package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ashishranjann/glofox.com/internal/model"
	"github.com/ashishranjann/glofox.com/internal/service"
)

type ClassHandler struct{
	service *service.ClassService
}

func NewClasshandler(s *service.ClassService) *ClassHandler{
	return &ClassHandler{service: s}
}

func (h *ClassHandler)CreateClassHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req model.ClassRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "inValid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateClass(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("message: class created successfully")
}