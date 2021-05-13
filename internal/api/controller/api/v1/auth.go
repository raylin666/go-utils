package v1

import (
	"go-server/internal/api/context"
	"go-server/internal/api/validator/request"
	"go-server/internal/constant"
	"go-server/internal/model"
)

// 获取 Token 认证
func GetTokenAuth(ctx *context.Context)  {
	secret := ctx.PostForm("secret")
	ok := ctx.RequestValidate(request.GetTokenAuthValidate{
		Secret: secret,
	})

	if ok {
		existSecret := model.Get().JwtSecretModel.ExistSecret(secret)
		if !existSecret {
			ctx.Error(constant.StatusAuthSecretNotFound)
			return
		}
		ctx.Success(context.H{})
	}
}

