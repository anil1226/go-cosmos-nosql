package user

import (
	"context"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserStore interface {
	GetUser(context.Context, string) (User, error)
	CreateUser(context.Context, User) error
}

type UserService struct {
	UserStore
}

func (s *UserService) GetUser(ctx context.Context, id string) (User, error) {
	emp, err := s.GetUser(ctx, id)
	if err != nil {
		return User{}, err
	}
	return emp, nil
}

func (s *UserService) CreateUser(ctx context.Context, emp User) error {
	err := s.CreateUser(ctx, emp)
	if err != nil {
		return err
	}
	return nil
}
