package container

import (
	"sheepim-user-service/biz/infra/config"
	"sync"
)

type Container struct {
	Config      *config.Config
	UserService user.IUserService
}

var APP *Container
var once sync.Once

func GetGlobalContainer() *Container {
	if APP == nil {
		panic("APP在使用前未初始化")
	}
	return APP
}

func InitGlobalContainer(env string) {
	once.Do(func() {
		APP = GetContainer(env)
	})
}

func NewContainer(config *config.Config, userService user.IUserService,
) *Container {
	return &Container{
		Config:      config,
		UserService: userService,
	}

}
