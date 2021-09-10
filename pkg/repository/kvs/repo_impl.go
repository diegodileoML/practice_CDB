package kvs

import (
	"context"

	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
)


func (r repository) GetByID(ctx context.Context, id string) (*basic.User, error) {
	user := &User{}
	err := r.repo.Get(ctx, id, user)
	if err!=nil{
		return nil, err
	}

	return user.ToDomain(),nil

}

func (r repository) Exists(ctx context.Context, id string) bool {
	err := r.repo.Get(ctx, id,nil)
	if err!=nil {
		return false
	}
	return true
}

func (r repository) Store(ctx context.Context, u *basic.User) error {
	return r.repo.Set(ctx,u.ID,userFromDomain(*u))
}

func (r repository) Update(ctx context.Context, u basic.User) error {
	return r.repo.Set(ctx,u.ID,userFromDomain(u))
}

func (r repository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(ctx,id)
}
