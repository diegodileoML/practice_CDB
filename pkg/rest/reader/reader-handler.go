package rest

import (
	"fmt"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/base/logger"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"net/http"
)

/*
func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) error {
	//func (h *handler) GetAll() gin.HandlerFunc{

	type response struct {
		Data []basic.User `json:"data"`
	}



	//return func(c *gin.Context) {
		ctx := r.Context()
		usuarios, err := h.basicSrv.GetAll(ctx)
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})


			return err
		}
		if len(usuarios) == 0 {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Usuarios no encontrados",
			})


			return &web.Error{Message: "Usuarios no encontrados"}
		}

		//ctx.JSON(http.StatusOK, &response{Data: usuarios})
		return web.RespondJSON(w, usuarios, http.StatusOK)
	}
	//}
 */



func (h *handler) GetByID(w http.ResponseWriter, r *http.Request) error {
	//h.GetAll()
	/*
	type response struct {
		Data basic.User `json:"data"`
	}

	 */
	//return func(c *gin.Context) {
	var err error
	ctx := r.Context()

	//id := ctx.Param("id")
	//userID, _ := strconv.Atoi(id)

	usrID, err := parseUserFromRequestID(r)
	if err!=nil{
		return err
	}

	usuario, err := h.basicSrv.GetByID(ctx, usrID)
	if err != nil {
			/*
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Usuario con ese ID no se encuentra",
			})

			 */
		logger.Error(ctx, "get_user_error", logger.Tag{
			"details": fmt.Sprintf("Error getting User with ID %s", usrID),
			"error":   err,
			"user_id":  usrID,
		})
		return err
	}
	return web.RespondJSON(w,usuario,http.StatusOK)
}
