package container

import (
	"github.com/li1553770945/sheepim-auth-service/biz/infra/config"
	"github.com/li1553770945/sheepim-auth-service/biz/internal/service"
	"sync"
)

type Container struct {
	Config      *config.Config
	SecretKeys  *config.SecretKeys
	AuthService service.IAuthService
}

var APP *Container
var once sync.Once

func GetGlobalContainer() *Container {
	if APP == nil {
		panic("APP在使用前未初始化")
	}
	return APP
}

func InitGlobalContainer(config *config.Config) {
	once.Do(func() {
		APP = GetContainer(config)
	})
}

func NewContainer(config *config.Config, secretKeys *config.SecretKeys, authService service.IAuthService) *Container {
	return &Container{
		Config:      config,
		SecretKeys:  secretKeys,
		AuthService: authService,
	}

}
