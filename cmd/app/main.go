package main

import (
	"github.com/diegodileoML/practice_CDB/cmd/config"
	"github.com/diegodileoML/practice_CDB/pkg/domain/basic"
	"github.com/diegodileoML/practice_CDB/pkg/repository/kvs"
	"github.com/diegodileoML/practice_CDB/pkg/rest"
	rest2 "github.com/diegodileoML/practice_CDB/pkg/rest/reader"
)

func main() {
	conf := config.Get()

	userRepo := kvs.NewRepository(conf.Service.Kvs)
	userCont := basic.Container{
		Storage: userRepo,
	}
	userService := basic.NewService(&userCont)
	userHandler := rest2.NewHandler(userService)


	if err := rest.API(&userHandler).Run(); err != nil {
		panic(err.Error())
	}

}
