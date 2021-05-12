package logic

import (
	"context"
	"go-server/grpc/auth/rpc/auth"
)

type TokenLogic struct {
	ctx context.Context
}

func NewTokenLogic(ctx context.Context) *TokenLogic {
	return &TokenLogic{
		ctx: ctx,
	}
}

// 获取系统信息
func (l *TokenLogic) GetTokenAuth(request *auth.GetTokenAuthRequest) (*auth.GetTokenAuthResponse, error) {
	return &auth.GetTokenAuthResponse{}, nil
}
