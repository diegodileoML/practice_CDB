package rest

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/logger"
	"net/http"

	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

func (h *handler) Store(w http.ResponseWriter, r *http.Request) error{
	ctx := r.Context()

	usrID,err := parseUserFromRequestBODY(ctx,r)
	if err != nil {
		return err
	}
	err = ValidateFields(usrID)
	if err != nil {
		return err
	}

	err = h.basicSrv.Store(ctx,usrID)
	if err != nil {
		logger.Error(ctx, "create-user-error", logger.Tag{
			"error": err,
		})
		return err
	} else {
		return web.RespondJSON(w, "user registered successfully", http.StatusCreated)
	}

}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) error {

	ctx:= r.Context()
	usrID,err := parseUserFromRequestBODY(ctx,r)
	if err != nil {
		return err
	}
	err= h.basicSrv.Update(ctx,*usrID)
	if err!=nil{
		return err
	}

	return web.RespondJSON(w, "user actualized successfully", http.StatusOK)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) error {
	ctx:= r.Context()
	usrID, err := parseUserFromRequestID(r)
	if err!=nil{
		return err
	}
	err = h.basicSrv.Delete(ctx,usrID)
	if err!=nil{
		return err
	}
	return web.RespondJSON(w, "user deleted successfully", http.StatusNoContent)
}


func UserToMap(user *basic.User) (structMap map[string]interface{}) {

	byteCode, _ := json.Marshal(user)

	json.Unmarshal(byteCode, &structMap)

	return

}

func ValidateFields(request *basic.User) error {

	jsonToMap := UserToMap(request)

	for key, value := range jsonToMap {
		if key != "id" && (value == "" || value == 0 || value == 0.0) {

			errorMessage := fmt.Sprintf("The field %s is required", key)

			return web.NewError(422, errorMessage)
		}

	}
	return nil
}
