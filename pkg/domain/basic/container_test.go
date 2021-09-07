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
	GetAllFake func(ctx context.Context) ([]User, error)
	GetByIDFake func(ctx context.Context, id int) (User, error)
	ExistsFake func(ctx context.Context, id int) bool
	StoreFake func(ctx context.Context, u User) (User, error)
	UpdateFake func(ctx context.Context, u User) error
	DeleteFake func(ctx context.Context, id int) error
}

func (r *fakeStorage) GetAll(ctx context.Context) ([]User, error){
	return r.GetAllFake(ctx)
}

func (r *fakeStorage) GetByID(ctx context.Context, id int) (User, error){
	return r.GetByIDFake(ctx,id)
}

func (r *fakeStorage) Exists(ctx context.Context, id int) bool{
	return r.ExistsFake(ctx,id)
}

func (r *fakeStorage) Store(ctx context.Context, u User) (User, error){
	return r.StoreFake(ctx,u)
}

func (r *fakeStorage) Update(ctx context.Context, u User) error{
	return r.UpdateFake(ctx,u)
}

func (r *fakeStorage) Delete(ctx context.Context, id int) error{
	return r.DeleteFake(ctx,id)
}
