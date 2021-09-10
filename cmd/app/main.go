package main

import (
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	rest "github.com/diegodileoML/practice_CDB/pkg/rest/reader"
	"github.com/joho/godotenv"
	"os"

	//"github.com/diegodileoML/practice_CDB/pkg/repository/kvs"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/kvs"

)

func main() {
	_ = godotenv.Load()
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	//db := db.StorageDB
	//router := gin.Default()
	router := nil
	userRepo := kvs.NewRepository()
	userCont := basic.Container{
		Storage: userRepo,
	}
	userService := basic.NewService(&userCont)
	userHandler := rest.NewHandler(userService)
	userRoutes := router.Group("api/practice/users")
	{
		userRoutes.POST("/", userHandler.Store())
		//userRoutes.GET("/", userHandler.GetAll())
		userRoutes.GET("/:id", userHandler.GetByID())
		userRoutes.PUT("/:id", userHandler.Update())
		userRoutes.DELETE("/:id", userHandler.Delete())
	}

	router.Run()
}
