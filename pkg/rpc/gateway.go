package rpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net/http"
)

var _ GRPCGateway = (*gRPCGateway)(nil)

type GRPCGateway interface {
	GetServer() *http.Server
	RegisterHandler(f func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) error
	ListenAndServe() error
}

type gRPCGateway struct {
	address string

	*grpc.ClientConn
	*runtime.ServeMux

	httpServer *http.Server
}

// NewGRPCGatewayServer 创建 GRPC 网关服务
func NewGRPCGatewayServer(address string, conn *grpc.ClientConn) GRPCGateway {
	var gw = new(gRPCGateway)
	gw.address = address
	gw.ServeMux = runtime.NewServeMux()
	gw.ClientConn = conn
	gw.httpServer = gw.GetServer()
	return gw
}

func (gw *gRPCGateway) RegisterHandler(f func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) error {
	return f(context.Background(), gw.ServeMux, gw.ClientConn)
}

func (gw *gRPCGateway) GetServer() *http.Server {
	if gw.httpServer == nil {
		gw.httpServer = &http.Server{
			Addr: gw.address,
			Handler: gw.ServeMux,
		}
	}

	return gw.httpServer
}

func (gw *gRPCGateway) ListenAndServe() error {
	return gw.httpServer.ListenAndServe()
}



