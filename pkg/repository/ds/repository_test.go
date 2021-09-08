package ds

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
import "github.com/DATA-DOG/go-sqlmock"
type fakeRepository struct{
	db sqlmock
}

func NewFakeRepository() *repository{
	return &repository{
		db: &fakeDB{},
	}
}

type fakeDB struct{

}
*/

func NewFakeRepository(fr *fakeRepository) *repository{
	return &repository{
		db: *fr,
	}
}

type fakeRepository struct{
	db *sql.DB
	GetAllFn func(ctx context.Context) ([]basic.User, error)
	GetByIDFn func(ctx context.Context, id int) (basic.User, error)
	ExistsFn func(ctx context.Context, id int) bool
	StoreFn func(ctx context.Context, u basic.User) (basic.User, error)
	UpdateFn func(ctx context.Context, u basic.User) error
	DeleteFn func(ctx context.Context, id int) error
}

 func (fr *fakeRepository) GetAll(ctx context.Context) ([]basic.User, error){
	 return fr.GetAllFn(ctx)
 }
func (fr *fakeRepository) GetByID(ctx context.Context, id int) (basic.User, error){
	return fr.GetByIDFn(ctx,id)
}
func (fr *fakeRepository) Exists(ctx context.Context, id int) bool{
	return fr.ExistsFn(ctx,id)
}
func (fr *fakeRepository) Store(ctx context.Context, u basic.User) (basic.User, error){
	return fr.StoreFn(ctx,u)
}
func (fr *fakeRepository) Update(ctx context.Context, u basic.User) error{
	return fr.UpdateFn(ctx,u)
}
func (fr *fakeRepository) Delete(ctx context.Context, id int) error{
	return fr.DeleteFn(ctx,id)
}

func TestNewRepository(t *testing.T) {
	db ,_,_:= sqlmock.New()
	r := NewRepository(db)
	assert.NotNil(t,r)
}




