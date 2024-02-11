package employee

import (
	"context"
)

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Branch Branch `json:"branch"`
}

type Branch struct {
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zip"`
}

type EmpService struct {
	EmpStore
}

type EmpStore interface {
	GetEmployee(context.Context, string) (Employee, error)
	CreateEmployee(context.Context, Employee) error
	UpdateEmployee(context.Context, Employee) error
	DeleteEmployee(context.Context, string) error
}

func (s *EmpService) GetEmployee(ctx context.Context, id string) (Employee, error) {
	emp, err := s.GetEmployee(ctx, id)
	if err != nil {
		return Employee{}, err
	}
	return emp, nil
}

func (s *EmpService) CreateEmployee(ctx context.Context, emp Employee) error {
	err := s.CreateEmployee(ctx, emp)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmpService) UpdateEmployee(ctx context.Context, emp Employee) error {
	err := s.UpdateEmployee(ctx, emp)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmpService) DeleteEmployee(ctx context.Context, id string) error {
	err := s.DeleteEmployee(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
