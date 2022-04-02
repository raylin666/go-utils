package rpc

import (
	"github.com/raylin666/go-utils/logger"
	"net"
	"testing"
)

func TestGRPCServer(t *testing.T) {
	l, _ := logger.NewJSONLogger()
	s := NewGRPCServer(
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

	servAddress := ":10098"
	lis, err := net.Listen("tcp", servAddress)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		s.Serve(lis)
	}()

	conn, err := GRPCDialContext(servAddress)
	if err != nil {
		t.Fatal(err)
	}

	gwAddress := ":10099"
	gw := NewGRPCGatewayServer(gwAddress, conn)
	go func() {
		gw.ListenAndServe()
	}()
}