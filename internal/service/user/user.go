package user

import (
	"context"

	"github.com/anil1226/go-employee/internal/db"
	"github.com/anil1226/go-employee/internal/models"
)

type UserService struct {
	db.UserStore
}

func (s *UserService) GetUser(ctx context.Context, id string) (models.User, error) {
	emp, err := s.GetUser(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return emp, nil
}

func (s *UserService) CreateUser(ctx context.Context, emp models.User) error {
	err := s.CreateUser(ctx, emp)
	if err != nil {
		return err
	}
	return nil
}
