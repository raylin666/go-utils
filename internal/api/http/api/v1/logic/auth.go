package logic

import (
	"go-server/internal/api/context"
	"go-server/internal/api/http/api/v1/types/params"
	"go-server/internal/constant"
	"go-server/internal/model"
)

type AuthLogic struct {
	ctx *context.Context
}

func NewAuthLogic(ctx *context.Context) *AuthLogic {
	return &AuthLogic{
		ctx: ctx,
	}
}

func (l *AuthLogic) GetTokenAuthLogic(req params.GetTokenAuthReq) (*params.GetTokenAuthResp, bool) {
	existSecret := model.Get().JwtSecretModel.ExistSecret(req.Secret)
	if !existSecret {
		l.ctx.ResponseBuilder.WithCode(constant.StatusAuthKeySecretNotFound)
		return nil, false
	}
	return &params.GetTokenAuthResp{}, true
}
