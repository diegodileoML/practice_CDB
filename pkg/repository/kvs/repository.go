package kvs

import (
	"context"
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/kvs"
)

type KVSRepository interface {
	//GetAll(ctx context.Context) ([]basic.User, error)
	GetByID(ctx context.Context, id string) (basic.User, error)
	Exists(ctx context.Context, id string) bool
	Store(ctx context.Context, u basic.User) error
	Update(ctx context.Context, u basic.User) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	repo kvs.Repository
}

func NewRepository(conf kvs.Config) KVSRepository {
	return &repository{kvs.NewRepository(conf)}
}
