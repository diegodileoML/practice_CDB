package writer

import (
	"context"

	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
)

type writer_handler struct {
	basicSrv basic.Service
}

func NewHandler(basic_serv *basic.Service) *writer_handler {
	return &writer_handler{
		basicSrv: *basic_serv,
	}
}

func (h *writer_handler) Store(ctx context.Context, u basic.User) (basic.User, error) {
	return h.basicSrv.Store(ctx, u)
}
func (h *writer_handler) Update(ctx context.Context, u basic.User) error {
	return h.basicSrv.Update(ctx, u)
}
func (h *writer_handler) Delete(ctx context.Context, id int) error {
	return h.basicSrv.Delete(ctx, id)
}
