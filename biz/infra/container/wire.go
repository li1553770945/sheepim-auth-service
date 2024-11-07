//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/li1553770945/sheepim-auth-service/biz/infra/config"
	"github.com/li1553770945/sheepim-auth-service/biz/infra/rpc"
	"github.com/li1553770945/sheepim-auth-service/biz/internal/service"
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
