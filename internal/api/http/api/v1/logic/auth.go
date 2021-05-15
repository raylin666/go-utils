package logic

import (
	"go-server/config"
	"go-server/internal/api/context"
	"go-server/internal/api/http/api/v1/types/params"
	"go-server/internal/constant"
	"go-server/internal/model"
	"go-server/pkg/jwt"
	"time"
)

type AuthLogic struct {
	ctx *context.Context
}

func NewAuthLogic(ctx *context.Context) *AuthLogic {
	return &AuthLogic{
		ctx: ctx,
	}
}

// 获取 Token 认证
func (l *AuthLogic) GetTokenAuthLogic(req params.GetTokenAuthReq) (*params.GetTokenAuthResp, bool) {
	getKeySecret := model.Get().JwtSecretModel.GetKeySecretFirst(req.Key, req.Secret)
	if getKeySecret.ID <= 0 {
		l.ctx.ResponseBuilder.WithCode(constant.StatusAuthKeySecretNotFound)
		return nil, false
	}

	// 判断 Key Secret 是否已过期
	if getKeySecret.ExpiredAt.Before(time.Now()) {
		l.ctx.ResponseBuilder.WithCode(constant.StatusAuthKeySecretExpire)
		return nil, false
	}

	if req.TTL <= 0 {
		req.TTL = config.Get().Jwt.TTL
	}

	secretUser := model.Get().JwtUsersModel.GetSecretUser(req.UserID, int(getKeySecret.ID))
	if secretUser.ID <= 0 {
		// 生成 Token
		newJwt := jwt.New(getKeySecret.App, getKeySecret.Key, getKeySecret.Secret)
		token, err := newJwt.GenerateToken(req.UserID, time.Duration(req.TTL))
		if err != nil {
			l.ctx.ResponseBuilder.WithCode(constant.StatusAuthTokenGenerateError)
			return nil, false
		}

		secretUser = &model.JwtUsers{
			SecretId:  int(getKeySecret.ID),
			UserID:    req.UserID,
			Token:     token,
			TTL:       req.TTL,
			ExpiredAt: time.Unix(time.Now().Unix()+int64(req.TTL), 0),
		}

		// 创建数据
		createId := model.Get().JwtUsersModel.Create(secretUser)
		if createId <= 0 {
			// 创建失败
			l.ctx.ResponseBuilder.WithCode(constant.StatusWriteDataError)
			return nil, false
		}

		secretUser.ID = createId
	}

	// 判断是否过期
	if secretUser.ExpiredAt.Before(time.Now()) {
		l.ctx.ResponseBuilder.WithCode(constant.StatusAuthTokenExpire)
		return nil, false
	}

	// 判断是否已删除
	if secretUser.DeletedAt != nil {
		l.ctx.ResponseBuilder.WithCode(constant.StatusAuthUserDeleted)
		return nil, false
	}

	return &params.GetTokenAuthResp{
		Key:       getKeySecret.Key,
		Secret:    getKeySecret.Secret,
		UserID:    secretUser.UserID,
		TTL:       secretUser.TTL,
		Token:     secretUser.Token,
		ExpiredAt: secretUser.ExpiredAt,
	}, true
}
