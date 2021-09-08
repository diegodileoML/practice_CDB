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
	GetAllFn func(ctx context.Context) ([]User, error)
	GetByIDFn func(ctx context.Context, id int) (User, error)
	ExistsFn func(ctx context.Context, id int) bool
	StoreFn func(ctx context.Context, u User) (User, error)
	UpdateFn func(ctx context.Context, u User) error
	DeleteFn func(ctx context.Context, id int) error
}

func (r *fakeStorage) GetAll(ctx context.Context) ([]User, error){
	return r.GetAllFn(ctx)
}

func (r *fakeStorage) GetByID(ctx context.Context, id int) (User, error){
	return r.GetByIDFn(ctx,id)
}

func (r *fakeStorage) Exists(ctx context.Context, id int) bool{
	return r.ExistsFn(ctx,id)
}

func (r *fakeStorage) Store(ctx context.Context, u User) (User, error){
	return r.StoreFn(ctx,u)
}

func (r *fakeStorage) Update(ctx context.Context, u User) error{
	return r.UpdateFn(ctx,u)
}

func (r *fakeStorage) Delete(ctx context.Context, id int) error{
	return r.DeleteFn(ctx,id)
}
