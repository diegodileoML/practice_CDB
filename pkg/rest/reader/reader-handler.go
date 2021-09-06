package reader

import (
	"net/http"
	"strconv"

	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/gin-gonic/gin"
)

type reader_handler struct {
	basicSrv basic.Service
}

func NewHandler(basic_serv *basic.Service) *reader_handler {
	return &reader_handler{
		basicSrv: *basic_serv,
	}
}

func (h *reader_handler) GetAll() gin.HandlerFunc {
	type response struct {
		Data []basic.User `json:"data"`
	}
	return func(c *gin.Context) {
		usuarios, err := h.basicSrv.GetAll(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(usuarios) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Usuarios no encontrados",
			})
			return
		}
		c.JSON(http.StatusOK, &response{Data: usuarios})
	}
}

func (h *reader_handler) GetByID() gin.HandlerFunc {
	h.GetAll()
	type response struct {
		Data basic.User `json:"data"`
	}
	return func(c *gin.Context) {
		id := c.Param("id")
		userID, _ := strconv.Atoi(id)
		usuario, err := h.basicSrv.GetByID(c, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, &response{Data: usuario})
		}

	}
}
