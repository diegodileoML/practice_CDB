package main

import (
	"github.com/diegodileoML/practice_CDB/pkg/db"
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/diegodileoML/practice_CDB/pkg/repository/ds"
	rest "github.com/diegodileoML/practice_CDB/pkg/rest/reader"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.StorageDB
	router := gin.Default()

	userRepo := ds.NewRepository(db)
	userCont := basic.Container{
		Storage: userRepo,
	}
	userService := basic.NewService(&userCont)
	userHandler := rest.NewHandler(&userService)
	userRoutes := router.Group("api/practice/users")
	{
		userRoutes.POST("/",userHandler.Store())
		userRoutes.GET("/",userHandler.GetAll())
		userRoutes.GET("/:id",userHandler.GetByID())
		userRoutes.PUT("/:id",userHandler.Update())
		userRoutes.DELETE("/:id",userHandler.Delete())
	}
}
