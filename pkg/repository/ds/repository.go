package ds

import (
	"context"
	"database/sql"

	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
)

type Repository interface {
	GetAll(ctx context.Context) ([]basic.User, error)
	GetByID(ctx context.Context, id int) (basic.User, error)
	Exists(ctx context.Context, id int) bool
	Store(ctx context.Context, u basic.User) (basic.User, error)
	Update(ctx context.Context, u basic.User) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
