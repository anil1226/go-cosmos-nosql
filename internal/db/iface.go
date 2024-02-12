package db

import (
	"context"

	"github.com/anil1226/go-employee/internal/models"
)

type Store interface {
	EmpStore
	UserStore
}

type UserStore interface {
	GetUser(context.Context, models.User) (models.User, error)
	CreateUser(context.Context, models.User) error
}

type EmpStore interface {
	GetEmployee(context.Context, string) (models.Employee, error)
	CreateEmployee(context.Context, models.Employee) error
	UpdateEmployee(context.Context, models.Employee) error
	DeleteEmployee(context.Context, string) error
}
