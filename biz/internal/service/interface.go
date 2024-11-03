package service

import (
	"context"
	"sheepim-auth-service/kitex_gen/auth"
)

type AuthService struct {
}

type IAuthService interface {
	Login(ctx context.Context, req *auth.LoginReq) (resp *auth.LoginResp, err error)
	Logout(ctx context.Context, req *auth.LogoutReq) (resp *auth.LogoutResp, err error)
	GenerateActivateCode(ctx context.Context, req *auth.GenerateActiveCodeReq) (resp *auth.GenerateActiveCodeResp, err error)
	Register(ctx context.Context, req *auth.RegisterReq) (resp *auth.RegisterResp, err error)
}

func NewAuthService() IAuthService {
	return &AuthService{}
}
