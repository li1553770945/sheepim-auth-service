package container

import (
	"github.com/google/wire"
	"sheepim-auth-service/biz/infra/config"
	"sheepim-auth-service/biz/internal/service"
)

func GetContainer(env string) *Container {
	panic(wire.Build(

		//infra
		config.InitConfig,

		//repo

		//service
		service.NewAuthService,

		NewContainer,
	))
}
