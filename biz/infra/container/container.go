package container

import (
	"sheepim-auth-service/biz/infra/config"
	user "sheepim-auth-service/biz/internal/service"
)

type Container struct {
	Config      *config.Config
	UserService user.IAuthService
}

func NewContainer(config *config.Config, userService user.IAuthService,
) *Container {
	return &Container{
		Config:      config,
		UserService: userService,
	}

}
