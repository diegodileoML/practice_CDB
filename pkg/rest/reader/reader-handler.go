package rest

import (
	"fmt"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/logger"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"net/http"
)

func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) error {

	var err error
	ctx := r.Context()

	usrID, err := parseUserFromRequestID(r)
	if err!=nil{
		return err
	}

	usuario, err := h.basicSrv.GetByID(ctx, usrID)
	if err != nil {
		logger.Error(ctx, "get_user_error", logger.Tag{
			"details": fmt.Sprintf("Error getting User with ID %s", usrID),
			"error":   err,
			"user_id":  usrID,
		})
		return err
	}
	return web.RespondJSON(w,usuario,http.StatusOK)
}
