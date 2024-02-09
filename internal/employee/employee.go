package employee

import (
	"context"
)

type Employee struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	BranchD Branch `json:"branch"`
}

type Branch struct {
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zip"`
}

type Store interface {
	GetEmployee(context.Context, string) (Employee, error)
}

type Service struct {
	Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetEmployee(ctx context.Context, id string) (Employee, error) {
	emp, err := s.Store.GetEmployee(ctx, id)
	if err != nil {
		return Employee{}, err
	}
	return emp, nil
}

func (s *Service) CreateEmployee(ctx context.Context, emp Employee) (Employee, error) {
	return Employee{}, nil
}

func (s *Service) UpdateEmployee(ctx context.Context, emp Employee) error {
	return nil
}

func (s *Service) DeleteEmployee(ctx context.Context, id string) error {
	return nil
}
