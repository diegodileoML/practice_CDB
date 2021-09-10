package basic

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Store_Success(t *testing.T) {
	ctx, fc, s:= IniciarDependencias()

	fc.Storage.StoreFn = func(ctx context.Context, u User) error {
		assert.Equal(t,"1",u.ID)
		return nil
	}

	err:=s.Store(ctx,User{ID:"1"})

	assert.Nil(t,err)
	//assert.Equal(t,User{ID:"1"},usr)

}

func TestService_Store_Error(t *testing.T) {
	ctx, fc, s:= IniciarDependencias()

	fc.Storage.StoreFn = func(ctx context.Context, u User)  error {
		return errors.New("forced error")
	}

	err:=s.Store(ctx,User{ID:"1"})

	assert.EqualError(t,err,"forced error")
	//assert.Empty(t,usr)
}

func TestService_Update_Success(t *testing.T) {
	ctx, fc,s := IniciarDependencias()

	fc.Storage.UpdateFn = func(ctx context.Context, u User) error {
		assert.Equal(t,u.Dni, 456)
		u = User{ID:"1",Dni: 456}
		return nil
	}
	/*
	fc.Storage.GetByIDFn = func(ctx context.Context,id int) (User,error) {
		assert.Equal(t, 1, id)
		return User{ID: 1,Dni: 456}, nil
	}
	 */

	err:=s.Update(ctx,User{ID:"1",Dni: 456})

	assert.Nil(t, err)
}

func TestService_Update_Error(t *testing.T) {
	ctx, fc,s := IniciarDependencias()

	fc.Storage.UpdateFn = func(ctx context.Context, u User) error {
		assert.Equal(t,u.Dni, 456)
		u = User{ID:"1",Dni: 456}
		return errors.New("forced error")
	}
	/*
	fc.Storage.GetByIDFn = func(ctx context.Context,id int) (User,error) {
		assert.Equal(t, 1, id)
		return User{ID: 1,Dni: 456}, nil
	}
	 */

	err:=s.Update(ctx,User{ID:"1",Dni: 456})

	assert.EqualError(t, err,"conflict: forced error")
}

func TestService_Delete_Success(t *testing.T) {
	ctx, fc,s := IniciarDependencias()
	fc.Storage.GetByIDFn = func(ctx context.Context,id string) (User,error) {
		assert.Equal(t, "1", id)
		return User{ID: "1"}, nil
	}
	fc.Storage.DeleteFn = func(ctx context.Context, id string) error {
		assert.Equal(t, "1", id)
		return nil
	}

	err:= s.Delete(ctx,"1")

	assert.Nil(t, err)
}

func TestService_Delete_Error(t *testing.T) {
	ctx, fc,s := IniciarDependencias()
	fc.Storage.GetByIDFn = func(ctx context.Context,id string) (User,error) {
		assert.Equal(t, "1", id)
		return User{ID: "1"}, nil
	}
	fc.Storage.DeleteFn = func(ctx context.Context, id string) error {
		assert.Equal(t, "1", id)
		return errors.New("forced error")
	}

	err:= s.Delete(ctx,"1")

	assert.EqualError(t, err,"internal_server_error: forced error")
}

func TestService_Delete_Error_IDinexistente(t *testing.T) {
	ctx, fc,s := IniciarDependencias()

	fc.Storage.GetByIDFn = func(ctx context.Context,id string) (User,error) {
		assert.Equal(t, "2", id)
		return User{}, errors.New("ID inexistente")
	}
	fc.Storage.DeleteFn = func(ctx context.Context, id string) error {
		assert.Equal(t, "2", id)
		return errors.New("forced error en update")
	}

	err:= s.Delete(ctx,"2")

	assert.EqualError(t, err,"not_found: ID inexistente")
}