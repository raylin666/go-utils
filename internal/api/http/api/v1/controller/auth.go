package controller

import (
	"go-server/internal/api/context"
	"go-server/internal/api/http/api/v1/types/params"
	"go-server/internal/constant"
	"go-server/internal/model"
)

// 获取 Token 认证
func GetTokenAuth(ctx *context.Context)  {
	var req params.GetTokenAuthRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(constant.StatusParamsParseError)
		return
	}

	if ok := ctx.RequestValidate(req); ok {
		existSecret := model.Get().JwtSecretModel.ExistSecret(req.Secret)
		if !existSecret {
			ctx.Error(constant.StatusAuthSecretNotFound)
			return
		}
		ctx.Success(context.H{})
	}
}

