//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"sheepim-auth-service/biz/infra/config"
	"sheepim-auth-service/biz/infra/rpc"
	"sheepim-auth-service/biz/internal/service"
)

func GetContainer(env string) *Container {
	panic(wire.Build(

		//infra
		config.GetConfig,
		config.GetSecretKey,

		//rpc
		rpc.NewUserClient,

		//service
		service.NewAuthService,

		NewContainer,
	))
}
