package service

import (
	"context"
	"sheepim-auth-service/kitex_gen/auth"
)

func (s *AuthService) Login(ctx context.Context, req *auth.LoginReq) (resp *auth.LoginResp, err error) {
	return
}
func (s *AuthService) Logout(ctx context.Context, req *auth.LogoutReq) (resp *auth.LogoutResp, err error) {
	return
}
func (s *AuthService) GenerateActivateCode(ctx context.Context, req *auth.GenerateActiveCodeReq) (resp *auth.GenerateActiveCodeResp, err error) {
	return
}
func (s *AuthService) Register(ctx context.Context, req *auth.RegisterReq) (resp *auth.RegisterResp, err error) {
	return
}
