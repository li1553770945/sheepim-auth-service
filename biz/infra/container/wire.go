//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"sheepim-user-service/biz/infra/config"
	"sheepim-user-service/biz/infra/database"
	"sheepim-user-service/biz/internal/repo"
)

func GetContainer(env string) *Container {
	panic(wire.Build(

		//infra
		config.InitConfig,

		//repo
		repo.NewRepository,
		database.NewDatabase,

		//service
		user.NewUserService,

		NewContainer,
	))
}
