package rest

import (
	"context"
	"errors"
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/diegodileoML/practice_CDB/test"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Get_Success(t *testing.T) {
	h,fbs := IniciarDependencias()

	ctx := test.FakeContext{ValueData: web.URIParams{"id": "test-id"}}
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "localhost:8080/{id}", nil)
	w := httptest.NewRecorder()

	fbs.GetByIDFn = func(ctx context.Context, id string) (*basic.User, error) {
		return &basic.User{
			ID:"1",
			FirstName: "diego",
			LastName: "dileo",
			Dni: 40,
			BirthDate: "0305",
			Email: "di@",
			Nacionality: "arg",
			Address: "dr montes",
		},nil
	}

	err:= h.GetByID(w,r)

	assert.Nil(t,err)
	assert.Equal(t, `{"id":"1","first_name":"diego","last_name":"dileo","dni":40,"birth_date":"0305","email":"di@","nacionality":"arg","address":"dr montes"}`,w.Body.String())
}

func TestHandler_Get_Error(t *testing.T) {
	h,fbs := IniciarDependencias()

	ctx := test.FakeContext{ValueData: web.URIParams{"id": "test-id"}}
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "localhost:8080/{id}", nil)
	w := httptest.NewRecorder()

	fbs.GetByIDFn = func(ctx context.Context, id string) (*basic.User, error) {
		return nil,errors.New("forced error")
	}

	err:= h.GetByID(w,r)

	assert.EqualError(t,err,"forced error")
}