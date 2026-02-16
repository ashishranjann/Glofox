package service

import (
	"errors"

	"github.com/ashishranjann/glofox.com/internal/model"
	"github.com/ashishranjann/glofox.com/internal/store"
	"github.com/ashishranjann/glofox.com/internal/utils"
)

type ClassService struct {
	store *store.InMemoryStore
}

func NewClassService(s *store.InMemoryStore) *ClassService {
	return &ClassService{
		store: s,
	}
}

func (s *ClassService) CreateClass(req *model.ClassRequest) error {
	if req.Name == "" || req.Capacity <= 0 {
		return errors.New("invalid request")
	}

	start, err := utils.ParseDate(req.Start_date)
	if err != nil {
		return err
	}

	end, err := utils.ParseDate(req.End_date)
	if err != nil {
		return err
	}

	s.store.Mu.Lock()
	defer s.store.Mu.Unlock()

	s.store.Classes[req.Name] = model.Class{
		Name:       req.Name,
		Start_date: start,
		End_date:   end,
		Capacity:   req.Capacity,
	}
	return nil

}
