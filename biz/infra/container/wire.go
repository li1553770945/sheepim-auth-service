//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"sheepim-auth-service/biz/infra/config"
	"sheepim-auth-service/biz/infra/rpc"
	"sheepim-auth-service/biz/internal/service"
)

func GetContainer(cfg *config.Config) *Container {
	panic(wire.Build(

		//infra
		config.GetSecretKey,

		//rpc
		rpc.NewUserClient,

		//service
		service.NewAuthService,

		NewContainer,
	))
}
