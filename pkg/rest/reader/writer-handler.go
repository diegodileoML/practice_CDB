package rest

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/logger"
	"net/http"

	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

func (h *handler) Store(w http.ResponseWriter, r *http.Request) error{ // gin.HandlerFunc {
	/*
	type response struct {
		Data basic.User `json:"data"`
	}

	 */

	//return func(c *gin.Context) {
	ctx := r.Context()
	//var req basic.User
	//err := c.Bind(&req)
	usrID,err := parseUserFromRequestBODY(ctx,r)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	err = ValidateFields(usrID)
	if err != nil {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return err
	}

	err = h.basicSrv.Store(ctx,*usrID)
	if err != nil {
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		logger.Error(ctx, "create-user-error", logger.Tag{
			"error": err,
		})
		return err
	}
	//c.JSON(http.StatusCreated, &response{Data: newUser})
	return web.RespondJSON(w, "user registered successfully", http.StatusCreated)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) error {//gin.HandlerFunc {

	/*
	//return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid ID"})
			return
		}

		var req basic.User
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		userSinModificar, err := h.basicSrv.GetByID(c, int(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		userIntoMap := UserToMap(&userSinModificar)

		for key, value := range UserToMap(&req) {
			if key == "id" || value == 0.0 || value == "" || value == 0 {
				continue
			} else {
				userIntoMap[key] = value
			}
		}
		patchedUser := MapToUser(userIntoMap)
		err = h.basicSrv.Update(c, *patchedUser)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	}
	 */
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

/*
func MapToUser(structMap map[string]interface{}) (user *basic.User) {

	byteCode, _ := json.Marshal(structMap)

	json.Unmarshal(byteCode, &user)

	return

}
*/

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
