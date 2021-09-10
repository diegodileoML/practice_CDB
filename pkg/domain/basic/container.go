package basic

import "context"

type Container struct {
	Storage Storage
}

type Storage interface {
	//GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
	Exists(ctx context.Context, id string) bool
	Store(ctx context.Context, u User) error
	Update(ctx context.Context, u User) error
	Delete(ctx context.Context, id string) error
}
