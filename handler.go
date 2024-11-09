package main

import (
	"context"
	"github.com/li1553770945/sheepim-auth-service/biz/infra/container"
	auth "github.com/li1553770945/sheepim-auth-service/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// Login implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginReq) (resp *auth.LoginResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.AuthService.Login(ctx, req)
	return
}

// Logout implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Logout(ctx context.Context, req *auth.LogoutReq) (resp *auth.LogoutResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.AuthService.Logout(ctx, req)
	return
}

// GenerateActiveCode implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GenerateActiveCode(ctx context.Context, req *auth.GenerateActiveCodeReq) (resp *auth.GenerateActiveCodeResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.AuthService.GenerateActivateCode(ctx, req)
	return
}

// Register implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Register(ctx context.Context, req *auth.RegisterReq) (resp *auth.RegisterResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.AuthService.Register(ctx, req)
	return
}

// GetUserId implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GetUserId(ctx context.Context, req *auth.GetUserIdReq) (resp *auth.GetUserIdResp, err error) {
	// TODO: Your code here...
	App := container.GetGlobalContainer()
	resp, err = App.AuthService.GetUserId(ctx, req)
	return
}
