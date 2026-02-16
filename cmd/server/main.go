package main

import (
	"log"
	"net/http"

	"github.com/ashishranjann/glofox.com/internal/handler"
	"github.com/ashishranjann/glofox.com/internal/service"
	"github.com/ashishranjann/glofox.com/internal/store"
)

func main() {
	store := store.NewInMemoryStore()
	class_service := service.NewClassService(store)
	booking_service := service.NewBookingService(store)

	class_handler := handler.NewClasshandler(class_service)
	booking_handler := handler.NewBookinghandler(booking_service)

	mux := http.NewServeMux()

	mux.HandleFunc("/classes", class_handler.CreateClassHandler)
	mux.HandleFunc("/bookings", booking_handler.CreateBookingHandler)

	log.Println("Glofox Server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
