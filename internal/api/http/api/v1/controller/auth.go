package controller

import (
	"go-server/internal/api/context"
	"go-server/internal/api/http/api/v1/logic"
	"go-server/internal/api/http/api/v1/types/params"
	"go-server/internal/constant"
)

// 获取 Token 认证
func GetTokenAuth(ctx *context.Context)  {
	var req params.GetTokenAuthReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(constant.StatusParamsParseError)
		return
	}

	if valid := ctx.RequestValidate(req); valid {
		l := logic.NewAuthLogic(ctx)
		resp, ok := l.GetTokenAuthLogic(req)
		if !ok {
			ctx.Error(ctx.ResponseBuilder.Code)
		} else {
			ctx.Success(resp)
		}
	}
}

// 验证 Token 认证
func VerifyTokenAuth(ctx *context.Context)  {
	var req params.VerifyTokenAuthReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(constant.StatusParamsParseError)
		return
	}

	if valid := ctx.RequestValidate(req); valid {
		l := logic.NewAuthLogic(ctx)
		resp, ok := l.VerifyTokenAuthLogic(req)
		if !ok {
			ctx.Error(ctx.ResponseBuilder.Code)
		} else {
			ctx.Success(resp)
		}
	}
}

// 刷新 Token 认证
func RefreshTokenAuth(ctx *context.Context)  {

}


