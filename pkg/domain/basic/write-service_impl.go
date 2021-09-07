package basic

import (
	"context"
	"net/http"

	"github.com/diegodileoML/practice_CDB/pkg/web"
)

func (s *service) Store(ctx context.Context, u User) (User, error) {

	if s.Exists(ctx, u.ID) {
		return User{}, &web.Error{Status: 409, Code: "409", Message: "ID de usuario repetido"}
	}

	usr, err := s.Storage.Store(ctx, u)
	if err != nil {
		return User{}, err
	}

	/*
	idUserNuevo, err := s.GetByID(ctx, usr.ID)
	if err != nil {
		return User{}, err
	}
	*/

	return usr, nil
}

func (s *service) Update(ctx context.Context, u User) error {
	_, err := s.GetAll(ctx)
	if err != nil {
		return web.NewError(http.StatusNotFound, err.Error())
	}

	err = s.Storage.Update(ctx, u)
	if err != nil {
		return web.NewError(http.StatusConflict, err.Error())
	}
	_, err = s.Storage.GetByID(ctx, u.ID)
	if err != nil {
		return web.NewError(http.StatusNotFound, err.Error())
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id int) error {
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
