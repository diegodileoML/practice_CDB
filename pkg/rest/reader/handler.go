package rest

import (
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
)

type handler struct {
	basicSrv basic.Service
}


func NewHandler(basic_serv basic.Service) handler {
	return handler{
		basicSrv: basic_serv,
	}
}
