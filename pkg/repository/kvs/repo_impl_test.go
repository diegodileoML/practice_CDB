package kvs

import (
	"context"
	"errors"
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/test"
	"github.com/stretchr/testify/assert"
	"testing"
)
func IniciarDependenciasRepo() (context.Context, repository, *fakeRepo){
	ctx := context.Background()
	fr := fakeRepo{}
	return ctx, repository{&fr}, &fr
}

func TestRepository_GetByID_Success(t *testing.T) {
	ctx,r,fr := IniciarDependenciasRepo()

	fr.GetFn = func(ctx context.Context, key string, value interface{}) error {
		if key == "1"{
			test.UpdateParam(value, basic.User{ID: "1"})
		}
		return nil
	}

	usr, err := r.GetByID(ctx,"1")

	assert.Equal(t,usr,basic.User{ID:"1"})
	assert.Nil(t,err)
}

func TestRepository_GetByID_Error(t *testing.T) {
	ctx, r, fr := IniciarDependenciasRepo()

	fr.GetFn = func(ctx context.Context, key string, value interface{}) error {
		return errors.New("get-user-error")
	}

	cdb, err := r.GetByID(ctx, "1")

	assert.Nil(t, cdb)
	assert.EqualError(t, err, "get-user-error")
}
