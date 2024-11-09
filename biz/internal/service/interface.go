package service

import (
	"context"
	"github.com/li1553770945/sheepim-auth-service/biz/infra/config"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user/userservice"
)

type AuthServiceImpl struct {
	UserRpcClient userservice.Client
	SecretKeys    *config.SecretKeys
}

type IAuthService interface {
	Login(ctx context.Context, req *auth.LoginReq) (resp *auth.LoginResp, err error)
	Logout(ctx context.Context, req *auth.LogoutReq) (resp *auth.LogoutResp, err error)
	GenerateActivateCode(ctx context.Context, req *auth.GenerateActiveCodeReq) (resp *auth.GenerateActiveCodeResp, err error)
	Register(ctx context.Context, req *auth.RegisterReq) (resp *auth.RegisterResp, err error)
	GetUserId(ctx context.Context, req *auth.GetUserIdReq) (resp *auth.GetUserIdResp, err error)
}

func NewAuthService(userRpcClient userservice.Client, secretKeys *config.SecretKeys) IAuthService {
	return &AuthServiceImpl{
		UserRpcClient: userRpcClient,
		SecretKeys:    secretKeys,
	}
}
