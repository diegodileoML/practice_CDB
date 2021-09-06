package basic

import (
	"context"
)

type Service interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int) (User, error)
	Exists(ctx context.Context, id int) bool
	Store(ctx context.Context, u User) (User, error)
	Update(ctx context.Context, u User) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	*Container
}

func NewService(cont *Container) Service {
	return &service{
		Container: cont,
	}
}
