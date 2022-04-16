package rpc

import (
	"context"
	"github.com/raylin666/go-utils/logger"
	"testing"
)

func TestGRPCServer(t *testing.T) {
	ctx := context.Background()
	l, _ := logger.NewJSONLogger()
	addr := ":10098"
	s := NewGRPCServer(
		WithGRPCServerNetwork("tcp"),
		WithGRPCServerAddress(addr),
		WithGRPCStreamInterceptor(
			CtxtagsStreamServerInterceptor(),
			OpentracingStreamServerInterceptor(),
			ZapStreamServerInterceptor(l),
			RecoveryStreamServerInterceptor(),
			),
		WithGRPCUnaryInterceptor(
			CtxtagsUnaryServerInterceptor(),
			OpentracingUnaryServerInterceptor(),
			ZapUnaryServerInterceptor(l),
			RecoveryUnaryServerInterceptor(),
			),
	)

	go func() {
		err := s.Start(ctx)
		if err != nil {
			panic(err)
		}
	}()
}