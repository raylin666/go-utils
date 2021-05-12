package server

import (
	"context"
	"go-server/grpc/auth/rpc/auth"
	"go-server/grpc/auth/rpc/srv"
)

// 鉴权服务
type AuthServer struct {
	srvCtx *svc.Context
}

func NewAuthServer(ctx *svc.Context) *AuthServer {
	return &AuthServer{
		srvCtx: ctx,
	}
}

// 获取 Token 认证接口
func (server *AuthServer) GetTokenAuth(ctx context.Context, request *auth.GetTokenAuthRequest) (*auth.GetTokenAuthResponse, error) {
	return &auth.GetTokenAuthResponse{}, nil
}

// 验证 Token 认证接口
func (server *AuthServer) VerifyTokenAuth(ctx context.Context, request *auth.VerifyTokenAuthRequest) (*auth.VerifyTokenAuthResponse, error) {
	return &auth.VerifyTokenAuthResponse{}, nil
}

// 刷新 Token 认证接口
func (server *AuthServer) RefreshTokenAuth(ctx context.Context, request *auth.RefreshTokenAuthRequest) (*auth.RefreshTokenAuthResponse, error) {
	return &auth.RefreshTokenAuthResponse{}, nil
}

// 删除 Token 认证接口
func (server *AuthServer) DeleteTokenAuth(ctx context.Context, request *auth.DeleteTokenAuthRequest) (*auth.DeleteTokenAuthResponse, error) {
	return &auth.DeleteTokenAuthResponse{}, nil
}

