package server

import (
	"context"
	"go-server/grpc/auth/rpc/auth"
)

// 鉴权服务
type Auth struct {}

// 获取 Token 认证接口
func (server *Auth) GetTokenAuth(ctx context.Context, request *auth.GetTokenAuthRequest) (*auth.GetTokenAuthResponse, error) {
	return &auth.GetTokenAuthResponse{}, nil
}

// 验证 Token 认证接口
func (server *Auth) VerifyTokenAuth(ctx context.Context, request *auth.VerifyTokenAuthRequest) (*auth.VerifyTokenAuthResponse, error) {
	return &auth.VerifyTokenAuthResponse{}, nil
}

// 刷新 Token 认证接口
func (server *Auth) RefreshTokenAuth(ctx context.Context, request *auth.RefreshTokenAuthRequest) (*auth.RefreshTokenAuthResponse, error) {
	return &auth.RefreshTokenAuthResponse{}, nil
}

// 删除 Token 认证接口
func (server *Auth) DeleteTokenAuth(ctx context.Context, request *auth.DeleteTokenAuthRequest) (*auth.DeleteTokenAuthResponse, error) {
	return &auth.DeleteTokenAuthResponse{}, nil
}

