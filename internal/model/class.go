package model

import (
	"time"
)

type Class struct {
	Name       string
	Start_date time.Time
	End_date   time.Time
	Capacity   int
}

type ClassRequest struct {
	Name       string `json:"name"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
	Capacity   int    `json:"capacity"`
}
