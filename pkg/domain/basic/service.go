package basic

import (
	"context"
)

type Service interface {
	GetByID(ctx context.Context, id string) (*User, error)
	Exists(ctx context.Context, id string) bool
	Store(ctx context.Context, u *User) error
	Update(ctx context.Context, u User) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	*Container
}

func NewService(cont *Container) Service {
	return &service{
		Container: cont,
	}
}
