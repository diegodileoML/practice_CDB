package rest

import (
	"context"
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fakeBasicService struct{
	//basic.Service
	GetByIDFn func(ctx context.Context, id string) (basic.User, error)
	ExistsFn func(ctx context.Context, id string) bool
	StoreFn func(ctx context.Context, u basic.User) error
	UpdateFn func(ctx context.Context, u basic.User) error
	DeleteFn func(ctx context.Context, id string) error
}


func (fbs *fakeBasicService) GetByID(ctx context.Context, id string) (basic.User, error){
	return fbs.GetByIDFn(ctx,id)
}

func (fbs *fakeBasicService) Exists(ctx context.Context, id string) bool{
	return fbs.ExistsFn(ctx,id)
}
func (fbs *fakeBasicService) Store(ctx context.Context, u basic.User) error{
	return fbs.StoreFn(ctx,u)
}
func (fbs *fakeBasicService) Update(ctx context.Context, u basic.User) error{
	return fbs.UpdateFn(ctx,u)
}
func (fbs *fakeBasicService) Delete(ctx context.Context, id string) error{
	return fbs.DeleteFn(ctx,id)
}

func TestNewHandler(t *testing.T){
	fbs:= &fakeBasicService{}
	h := NewHandler(fbs)

	assert.NotNil(t, h)
}

func IniciarDependencias() (*handler, *fakeBasicService ){
	fbs := &fakeBasicService{}
	h := &handler{basicSrv: fbs}

	return h,fbs
}