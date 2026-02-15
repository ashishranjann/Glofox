package store

import (
	"sync"

	"github.com/ashishranjann/glofox.com/internal/model"
)

type InMemoryStore struct {
	Classes        map[string]model.Class
	BookingHistory []model.BookingRequest
	Mu             sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		Classes:        make(map[string]model.Class),
		BookingHistory: []model.BookingRequest{},
	}
}
