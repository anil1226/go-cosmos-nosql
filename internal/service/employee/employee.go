package employee

import (
	"context"

	"github.com/anil1226/go-employee/internal/db"
	"github.com/anil1226/go-employee/internal/models"
)

type EmpService struct {
	db.EmpStore
}

func (s *EmpService) GetEmployee(ctx context.Context, id string) (models.Employee, error) {
	emp, err := s.GetEmployee(ctx, id)
	if err != nil {
		return models.Employee{}, err
	}
	return emp, nil
}

func (s *EmpService) CreateEmployee(ctx context.Context, emp models.Employee) error {
	err := s.CreateEmployee(ctx, emp)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmpService) UpdateEmployee(ctx context.Context, emp models.Employee) error {
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
