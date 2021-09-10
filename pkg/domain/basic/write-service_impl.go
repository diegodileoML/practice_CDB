package basic

import (
	"context"
	"net/http"

	"github.com/diegodileoML/practice_CDB/pkg/web"
)

func (s *service) Store(ctx context.Context, u *User) error {

	err := s.Storage.Store(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, u User) error {

	err := s.Storage.Update(ctx, u)
	if err != nil {
		return web.NewError(http.StatusConflict, err.Error())
	}
	return nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	_, err := s.Storage.GetByID(ctx, id)
	if err != nil {
		return web.NewError(http.StatusNotFound, err.Error())
	}

	err = s.Storage.Delete(ctx, id)
	if err != nil {
		return web.NewError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
