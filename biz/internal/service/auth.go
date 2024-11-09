package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt/v4"
	"github.com/li1553770945/sheepim-auth-service/biz/constant"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/base"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user"
	"time"
)

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {

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
func (s *AuthServiceImpl) GetUserId(ctx context.Context, req *auth.GetUserIdReq) (resp *auth.GetUserIdResp, err error) {
	klog.CtxInfof(ctx, "解析用户的 JWT 令牌")

	// 获取令牌
	tokenStr := req.GetToken()
	resp = &auth.GetUserIdResp{
		BaseResp: &base.BaseResp{
			Code:    constant.Success,
			Message: "成功",
		},
	}

	// 解析令牌
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的签名方法是 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("不支持的签名", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(s.SecretKeys.JWTKey), nil
	})

	if err != nil {
		klog.CtxErrorf(ctx, "令牌解析失败: %v", err)
		resp.BaseResp.Code = constant.Unauthorized
		resp.BaseResp.Message = "无效的令牌"
		return resp, nil
	}

	// 检查令牌是否有效并提取声明
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 提取 userId
		if userId, ok := claims["userId"].(float64); ok {
			resp.UserId = int32(userId)
			klog.CtxInfof(ctx, "令牌解析成功，用户ID: %d", resp.UserId)
			return resp, nil
		}
		// 如果无法提取 userId
		klog.CtxErrorf(ctx, "令牌解析失败：无法提取 userId")
		resp.BaseResp.Code = constant.Unauthorized
		resp.BaseResp.Message = "无效的令牌"
		return resp, nil
	} else {
		klog.CtxErrorf(ctx, "令牌无效")
		resp.BaseResp.Code = constant.Unauthorized
		resp.BaseResp.Message = "无效的令牌"
		return resp, nil
	}

}
func (s *AuthServiceImpl) Logout(ctx context.Context, req *auth.LogoutReq) (resp *auth.LogoutResp, err error) {
	return
}
func (s *AuthServiceImpl) GenerateActivateCode(ctx context.Context, req *auth.GenerateActiveCodeReq) (resp *auth.GenerateActiveCodeResp, err error) {
	return
}
func (s *AuthServiceImpl) Register(ctx context.Context, req *auth.RegisterReq) (resp *auth.RegisterResp, err error) {
	return
}
