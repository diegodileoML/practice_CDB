package basic

import "context"

type fakeContainer struct{
	Storage *fakeStorage
}

func (fc *fakeContainer) toContainer() *Container{
	return &Container{
		Storage: fc.Storage,
	}
}

func NewFakeContainer() *fakeContainer{
	return &fakeContainer{
		Storage: &fakeStorage{},
	}
}

type fakeStorage struct{
	GetByIDFn func(ctx context.Context, id string) (*User, error)
	ExistsFn func(ctx context.Context, id string) bool
	StoreFn func(ctx context.Context, u *User) error
	UpdateFn func(ctx context.Context, u User) error
	DeleteFn func(ctx context.Context, id string) error
}

func (r *fakeStorage) GetByID(ctx context.Context, id string) (*User, error){
	return r.GetByIDFn(ctx,id)
}

func (r *fakeStorage) Exists(ctx context.Context, id string) bool{
	return r.ExistsFn(ctx,id)
}

func (r *fakeStorage) Store(ctx context.Context, u *User) error{
	return r.StoreFn(ctx,u)
}

func (r *fakeStorage) Update(ctx context.Context, u User) error{
	return r.UpdateFn(ctx,u)
}

func (r *fakeStorage) Delete(ctx context.Context, id string) error{
	return r.DeleteFn(ctx,id)
}
