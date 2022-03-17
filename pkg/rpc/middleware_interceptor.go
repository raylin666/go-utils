package rpc

import (
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type StreamServerInterceptor grpc.StreamServerInterceptor

func CtxtagsStreamServerInterceptor(opts ...grpc_ctxtags.Option) StreamServerInterceptor {
	return StreamServerInterceptor(grpc_ctxtags.StreamServerInterceptor(opts...))
}

func OpentracingStreamServerInterceptor(opts ...grpc_opentracing.Option) StreamServerInterceptor {
	return StreamServerInterceptor(grpc_opentracing.StreamServerInterceptor(opts...))
}

func ZapStreamServerInterceptor(logger *zap.Logger, opts ...grpc_zap.Option) StreamServerInterceptor {
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	return StreamServerInterceptor(grpc_zap.StreamServerInterceptor(logger, opts...))
}

func AuthStreamServerInterceptor(authFunc grpc_auth.AuthFunc) StreamServerInterceptor {
	return StreamServerInterceptor(grpc_auth.StreamServerInterceptor(authFunc))
}

func RecoveryStreamServerInterceptor(opts ...grpc_recovery.Option) StreamServerInterceptor {
	return StreamServerInterceptor(grpc_recovery.StreamServerInterceptor(opts...))
}

type UnaryServerInterceptor grpc.UnaryServerInterceptor


func CtxtagsUnaryServerInterceptor(opts ...grpc_ctxtags.Option) UnaryServerInterceptor {
	return UnaryServerInterceptor(grpc_ctxtags.UnaryServerInterceptor(opts...))
}

func OpentracingUnaryServerInterceptor(opts ...grpc_opentracing.Option) UnaryServerInterceptor {
	return UnaryServerInterceptor(grpc_opentracing.UnaryServerInterceptor(opts...))
}

func ZapUnaryServerInterceptor(logger *zap.Logger, opts ...grpc_zap.Option) UnaryServerInterceptor {
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	return UnaryServerInterceptor(grpc_zap.UnaryServerInterceptor(logger, opts...))
}

func AuthUnaryServerInterceptor(authFunc grpc_auth.AuthFunc) UnaryServerInterceptor {
	return UnaryServerInterceptor(grpc_auth.UnaryServerInterceptor(authFunc))
}

func RecoveryUnaryServerInterceptor(opts ...grpc_recovery.Option) UnaryServerInterceptor {
	return UnaryServerInterceptor(grpc_recovery.UnaryServerInterceptor(opts...))
}
