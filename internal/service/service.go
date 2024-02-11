package service

import (
	"github.com/anil1226/go-employee/internal/service/employee"
	"github.com/anil1226/go-employee/internal/service/user"
)

type Store interface {
	employee.EmpStore
	user.UserStore
}

type Service struct {
	Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}
