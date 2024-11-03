// Code generated by Kitex v0.7.2. DO NOT EDIT.

package authservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	auth "sheepim-auth-service/kitex_gen/auth"
)

func serviceInfo() *kitex.ServiceInfo {
	return authServiceServiceInfo
}

var authServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AuthService"
	handlerType := (*auth.AuthService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Login":              kitex.NewMethodInfo(loginHandler, newAuthServiceLoginArgs, newAuthServiceLoginResult, false),
		"Logout":             kitex.NewMethodInfo(logoutHandler, newAuthServiceLogoutArgs, newAuthServiceLogoutResult, false),
		"GenerateActiveCode": kitex.NewMethodInfo(generateActiveCodeHandler, newAuthServiceGenerateActiveCodeArgs, newAuthServiceGenerateActiveCodeResult, false),
		"Register":           kitex.NewMethodInfo(registerHandler, newAuthServiceRegisterArgs, newAuthServiceRegisterResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "auth",
		"ServiceFilePath": `idl\auth.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceLoginArgs)
	realResult := result.(*auth.AuthServiceLoginResult)
	success, err := handler.(auth.AuthService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceLoginArgs() interface{} {
	return auth.NewAuthServiceLoginArgs()
}

func newAuthServiceLoginResult() interface{} {
	return auth.NewAuthServiceLoginResult()
}

func logoutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceLogoutArgs)
	realResult := result.(*auth.AuthServiceLogoutResult)
	success, err := handler.(auth.AuthService).Logout(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceLogoutArgs() interface{} {
	return auth.NewAuthServiceLogoutArgs()
}

func newAuthServiceLogoutResult() interface{} {
	return auth.NewAuthServiceLogoutResult()
}

func generateActiveCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceGenerateActiveCodeArgs)
	realResult := result.(*auth.AuthServiceGenerateActiveCodeResult)
	success, err := handler.(auth.AuthService).GenerateActiveCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceGenerateActiveCodeArgs() interface{} {
	return auth.NewAuthServiceGenerateActiveCodeArgs()
}

func newAuthServiceGenerateActiveCodeResult() interface{} {
	return auth.NewAuthServiceGenerateActiveCodeResult()
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthServiceRegisterArgs)
	realResult := result.(*auth.AuthServiceRegisterResult)
	success, err := handler.(auth.AuthService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthServiceRegisterArgs() interface{} {
	return auth.NewAuthServiceRegisterArgs()
}

func newAuthServiceRegisterResult() interface{} {
	return auth.NewAuthServiceRegisterResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, req *auth.LoginReq) (r *auth.LoginResp, err error) {
	var _args auth.AuthServiceLoginArgs
	_args.Req = req
	var _result auth.AuthServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Logout(ctx context.Context, req *auth.LogoutReq) (r *auth.LogoutResp, err error) {
	var _args auth.AuthServiceLogoutArgs
	_args.Req = req
	var _result auth.AuthServiceLogoutResult
	if err = p.c.Call(ctx, "Logout", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GenerateActiveCode(ctx context.Context, req *auth.GenerateActiveCodeReq) (r *auth.GenerateActiveCodeResp, err error) {
	var _args auth.AuthServiceGenerateActiveCodeArgs
	_args.Req = req
	var _result auth.AuthServiceGenerateActiveCodeResult
	if err = p.c.Call(ctx, "GenerateActiveCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, req *auth.RegisterReq) (r *auth.RegisterResp, err error) {
	var _args auth.AuthServiceRegisterArgs
	_args.Req = req
	var _result auth.AuthServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
