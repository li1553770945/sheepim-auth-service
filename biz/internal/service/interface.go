package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type AuthService struct {
	Repo *Re
}

type IAuthService interface {
	Login(ctx context.Context, c *app.RequestContext)
	Logout(ctx context.Context, c *app.RequestContext)
	GenerateActivateCode(ctx context.Context, c *app.RequestContext)
	Register(ctx context.Context, c *app.RequestContext)
}

func NewAuthService(repo repo.IRepository) IAuthService {
	return &AuthService{
		Repo: repo,
	}
}
