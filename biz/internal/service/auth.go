package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt/v5"
	"sheepim-auth-service/biz/constant"
	"sheepim-auth-service/kitex_gen/auth"
	"sheepim-auth-service/kitex_gen/base"
	"sheepim-user-service/kitex_gen/user"
	"time"
)

func (s *AuthService) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {

	klog.CtxInfof(ctx, "收到用户 %s 的登录请求", req.Username)
	checkUsernameAndPasswdResp, err := s.UserRpcClient.CheckUsernameAndPasswd(ctx, &user.CheckUsernameAndPasswdReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		klog.CtxErrorf(ctx, " %s 调用user用户名密码核验服务失败： 错误: %v", req.Username, err)
		return nil, err
	}
	resp := &auth.LoginResp{
		BaseResp: &base.BaseResp{
			Code: constant.Success,
		},
	}

	if checkUsernameAndPasswdResp.BaseResp.Code != 0 {
		klog.CtxInfof(ctx, "用户 %s 登录失败，用户名密码错误", req.Username)
		resp.BaseResp.Code = checkUsernameAndPasswdResp.BaseResp.Code
		resp.BaseResp.Message = checkUsernameAndPasswdResp.BaseResp.Message
		return resp, nil
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": checkUsernameAndPasswdResp.UserId,
			"exp":    time.Now().Add(14 * 24 * time.Hour).Unix(),
		})
	token, err := t.SignedString([]byte(s.SecretKeys.JWTKey))
	if err != nil {
		klog.CtxErrorf(ctx, "jwt 加密失败："+err.Error())
		return nil, err
	}
	resp.Token = &token
	klog.CtxInfof(ctx, "用户 %s 登录成功", req.Username)
	return resp, nil
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
