package basic

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func IniciarDependencias() (context.Context,*fakeContainer, Service){
	ctx := context.Background()
	fc := NewFakeContainer()
	s:= NewService(fc.toContainer())
	return ctx, fc, s
}
/*
func TestService_GetAll_Success(t *testing.T) {

	ctx,fc,s := IniciarDependencias()

	fc.Storage.GetAllFn = func(ctx context.Context) ([]User,error){
		return []User{{ID:1},{ID:2}},nil
	}
	usr, err := s.GetAll(ctx)

	assert.NotEmpty(t,usr)
	assert.Nil(t,err)
}

func TestService_GetAll_ErrorRepetido(t *testing.T) {

	ctx,fc,s := IniciarDependencias()

	fc.Storage.GetAllFn = func(ctx context.Context) ([]User,error){
		return []User{{ID:1},{ID:1}},nil
	}

	fc.Storage.DeleteFn = func (ctx context.Context,id int) error{
		return errors.New("usuario repetido")
	}

	usr, err := s.GetAll(ctx)

	assert.Empty(t,usr)
	assert.EqualError(t,err,"usuario repetido")
}

func TestService_GetAll_Error(t *testing.T) {

	ctx,fc,s := IniciarDependencias()

	fc.Storage.GetAllFn = func(ctx context.Context) ([]User,error){
		return []User{},errors.New("forced error")
	}

	usr, err := s.GetAll(ctx)

	assert.Equal(t,usr,[]User{})
	assert.EqualError(t, err,"forced error")
}
 */

func TestService_GetByID_Success(t *testing.T) {
	ctx,fc,s:= IniciarDependencias()

	fc.Storage.GetByIDFn = func(ctx context.Context,id string) (User,error) {
		assert.Equal(t, "1", id)
		return User{ID: "1"}, nil
	}

	usr ,err:= s.GetByID(ctx,"1")

	assert.Equal(t,User{ID:"1"},usr)
	assert.Nil(t,err)
}

func TestService_GetByID_Error(t *testing.T) {
	ctx,fc,s:= IniciarDependencias()

	fc.Storage.GetByIDFn = func(ctx context.Context,id string) (User,error){
		return User{},errors.New("forced error")
	}

	usr ,err:= s.GetByID(ctx,"1")

	assert.Empty(t,usr)
	assert.EqualError(t, err,"forced error")
}

func TestService_Exists(t *testing.T) {
	ctx, fc, s:= IniciarDependencias()

	fc.Storage.ExistsFn = func(ctx context.Context, id string) bool {
		assert.Equal(t,"1",id)
		return false
	}

	err := s.Exists(ctx,"1")

	assert.Equal(t,false,err)

}