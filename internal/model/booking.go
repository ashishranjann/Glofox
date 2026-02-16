package model

import (
	"time"
)

type Booking struct {
	Name       string
	Date       time.Time
	Class_name string
}

type BookingRequest struct {
	Name       string `json:"name"`
	Date       string `json:"date"`
	Class_name string `json:"class_name"`
}
