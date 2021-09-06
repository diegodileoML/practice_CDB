package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/diegodileoML/practice_CDB/pkg/web"
	"github.com/gin-gonic/gin"
)

func (h *handler) Store() gin.HandlerFunc {
	type response struct {
		Data basic.User `json:"data"`
	}

	return func(c *gin.Context) {
		var req basic.User
		err := c.Bind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = ValidateFields(&req)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		newUser, err := h.basicSrv.Store(c, req)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, &response{Data: newUser})
	}
}
func (h *handler) Update() gin.HandlerFunc {

	return func(c *gin.Context) {
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
}
func (h *handler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func UserToMap(user *basic.User) (structMap map[string]interface{}) {

	byteCode, _ := json.Marshal(user)

	json.Unmarshal(byteCode, &structMap)

	return

}

func MapToUser(structMap map[string]interface{}) (user *basic.User) {

	byteCode, _ := json.Marshal(structMap)

	json.Unmarshal(byteCode, &user)

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
