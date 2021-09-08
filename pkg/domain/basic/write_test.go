package basic

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Store_Success(t *testing.T) {
	ctx, fc, s:= IniciarDependencias()

	fc.Storage.StoreFn = func(ctx context.Context, u User) (User, error) {
		return User{ID:1},nil
	}

	usr,err:=s.Store(ctx,User{ID:1})

	assert.Nil(t,err)
	assert.Equal(t,User{ID:1},usr)

}

func TestService_Store_Error(t *testing.T) {
	ctx, fc, s:= IniciarDependencias()

	fc.Storage.StoreFn = func(ctx context.Context, u User) (User, error) {
		return User{},errors.New("forced error")
	}

	usr,err:=s.Store(ctx,User{ID:1})

	assert.EqualError(t,err,"forced error")
	assert.Empty(t,usr)

}
