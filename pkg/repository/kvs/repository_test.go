package kvs

import (
	"context"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/kvs"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fakeRepo struct{
	GetFn func(ctx context.Context, key string, value interface{}) error

	SetFn func(ctx context.Context, key string, value interface{}) error

	SetWithTTLFn func(ctx context.Context, key string, value interface{}, ttl int64) error

	SetWithDefaultTTLFn func(ctx context.Context, key string, value interface{}) error

	DeleteFn func(ctx context.Context, key string) error

	FlushCacheFn func(ctx context.Context) error
}

func (fr fakeRepo) Get(ctx context.Context, key string, value interface{}) error {
	return fr.GetFn(ctx,key,value)
}

func (fr fakeRepo) Set(ctx context.Context, key string, value interface{}) error {
	return fr.SetFn(ctx,key,value)
}

func (fr fakeRepo) SetWithTTL(ctx context.Context, key string, value interface{}, ttl int64) error {
	return fr.SetWithTTLFn(ctx,key,value,ttl)
}

func (fr fakeRepo) SetWithDefaultTTL(ctx context.Context, key string, value interface{}) error {
	return fr.SetWithDefaultTTLFn(ctx,key,value)
}

func (fr fakeRepo) Delete(ctx context.Context, key string) error {
	return fr.DeleteFn(ctx,key)
}

func (fr fakeRepo) FlushCache(ctx context.Context) error {
	return fr.FlushCacheFn(ctx)
}

func TestNewRepository(t *testing.T) {
	r := NewRepository(kvs.Config{Name: ""})
	assert.NotNil(t, r)
}




