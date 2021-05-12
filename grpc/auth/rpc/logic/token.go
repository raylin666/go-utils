package logic

import (
	"context"
	"go-server/grpc/auth/rpc/auth"
	"go-server/grpc/auth/rpc/srv"
)

type TokenLogic struct {
	ctx    context.Context
	srvCtx *svc.Context
}

func NewTokenLogic(ctx context.Context, srvCtx *svc.Context) *TokenLogic {
	return &TokenLogic{
		ctx:    ctx,
		srvCtx: srvCtx,
	}
}

// 获取系统信息
func (l *TokenLogic) GetTokenAuth(request *auth.GetTokenAuthRequest) (*auth.GetTokenAuthResponse, error) {
	return &auth.GetTokenAuthResponse{}, nil
}
