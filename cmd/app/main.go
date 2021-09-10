package main

import (
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/diegodileoML/practice_CDB/pkg/repository/kvs"
	"github.com/diegodileoML/practice_CDB/pkg/rest"
	rest2 "github.com/diegodileoML/practice_CDB/pkg/rest/reader"

	kvs2 "github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/repository/kvs"

)

func main() {
	/*
	_ = godotenv.Load()
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	*/


	userRepo := kvs.NewRepository(kvs2.Config{})
	userCont := basic.Container{
		Storage: userRepo,
	}
	userService := basic.NewService(&userCont)
	userHandler := rest2.NewHandler(userService)


	if err := rest.API(&userHandler).Run(); err != nil {
		panic(err.Error())
	}

}
