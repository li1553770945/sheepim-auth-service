namespace go auth

include "base.thrift"

struct LoginReq {
    1: required string username
    2: required string password
}

struct LoginResp {
    1: required base.BaseResp baseResp
    2: optional string token
}


struct LogoutReq {
    1: required string token
}

struct LogoutResp {
    1: required base.BaseResp baseResp
}


struct GenerateActiveCodeReq {
    1: required string username
}

struct GenerateActiveCodeResp {
    1: required base.BaseResp baseResp
    2: optional string token
}

struct RegisterReq{
    1: required string username
    2: required string nickname
    3: required string password
}

struct RegisterResp{
    1: required base.BaseResp baseResp
}

struct GetUserIdReq{
    1: required string token
}
struct GetUserIdResp{
    1: required base.BaseResp baseResp
    2: required i64 userId
}
service AuthService {
    LoginResp Login(LoginReq req)
    LogoutResp Logout(LogoutReq req)
    GenerateActiveCodeResp GenerateActiveCode(GenerateActiveCodeReq req)
    RegisterResp Register(RegisterReq req)
    GetUserIdResp GetUserId(GetUserIdReq req)

}
