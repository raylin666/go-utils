package rpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/raylin666/go-utils/http"
	"google.golang.org/grpc"
	nethttp "net/http"
)

var _ GRPCGateway = (*gRPCGateway)(nil)

type GRPCGatewayServerOption func(*gRPCGateway)

func WithGRPCGatewayContext(ctx context.Context) GRPCGatewayServerOption {
	return func(gateway *gRPCGateway) {
		gateway.context = ctx
	}
}

func WithGRPCGatewayGatewayAddress(gatewayAddress string) GRPCGatewayServerOption {
	return func(gateway *gRPCGateway) {
		gateway.gatewayAddress = gatewayAddress
	}
}

func WithGRPCGatewayGRPCAddress(grpcAddress string) GRPCGatewayServerOption {
	return func(gateway *gRPCGateway) {
		gateway.grpcAddress = grpcAddress
	}
}

func WithGRPCGatewayGRPCClientDialOption(opts ...grpc.DialOption) GRPCGatewayServerOption {
	return func(gateway *gRPCGateway) {
		gateway.grpcClientDialOption = append(gateway.grpcClientDialOption, opts...)
	}
}

type GRPCGateway interface {
	Context() context.Context
	ClientConn() *grpc.ClientConn
	RegisterServerHandlers(ctx context.Context, mux *runtime.ServeMux, srvHandlers ...func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) *GatewayServerHandler
	ListenHTTPServe() error
	ListenHTTPShutdown() error
}

type gRPCGateway struct {
	grpcClientConn *grpc.ClientConn
	httpServer     *http.Server

	grpcClientDialOption []grpc.DialOption
	context              context.Context
	gatewayAddress       string
	grpcAddress          string
	gatewayServerHandler *GatewayServerHandler
}

type GatewayServerHandler struct {
	context     context.Context
	mux         *runtime.ServeMux
	srvHandlers []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
	errors      []error
}

// NewGRPCGatewayServer 创建 GRPC 网关服务
func NewGRPCGatewayServer(opts ...GRPCGatewayServerOption) (GRPCGateway, error) {
	var gateway = &gRPCGateway{
		context: context.Background(),
	}
	for _, opt := range opts {
		opt(gateway)
	}

	conn, err := GRPCDialContext(gateway.context, gateway.grpcAddress, gateway.grpcClientDialOption...)

	if err != nil {
		return nil, err
	}

	gateway.grpcClientConn = conn
	gateway.httpServer = http.NewServer(&nethttp.Server{}, http.WithServerAddress(gateway.gatewayAddress))
	return gateway, nil
}

// Context 网关上下文
func (gateway *gRPCGateway) Context() context.Context {
	return gateway.context
}

// ClientConn 获取给定目标的客户端连接
func (gateway *gRPCGateway) ClientConn() *grpc.ClientConn {
	return gateway.grpcClientConn
}

// RegisterServerHandlers 注册业务服务处理器
func (gateway *gRPCGateway) RegisterServerHandlers(ctx context.Context, mux *runtime.ServeMux, srvHandlers ...func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) *GatewayServerHandler {
	gateway.gatewayServerHandler = new(GatewayServerHandler)
	gateway.gatewayServerHandler.context = ctx
	gateway.gatewayServerHandler.mux = mux
	gateway.gatewayServerHandler.srvHandlers = srvHandlers

	// HTTP 处理器 (注意: 不能遗漏该注册, 否则将无法访问到网关方法, 都会报错404.)
	gateway.httpServer.Handler = mux
	// 注册处理函数
	for _, handler := range srvHandlers {
		errHandler := handler(ctx, mux, gateway.ClientConn())
		if errHandler != nil {
			gateway.gatewayServerHandler.errors = append(gateway.gatewayServerHandler.errors, errHandler)
		}
	}

	return gateway.gatewayServerHandler
}

// ListenHTTPServeStart 启动网关 HTTP 服务监听
func (gateway *gRPCGateway) ListenHTTPServe() error {
	return gateway.httpServer.Start(gateway.context)
}

// ListenHTTPServeStop 停止网关 HTTP 服务监听
func (gateway *gRPCGateway) ListenHTTPShutdown() error {
	return gateway.httpServer.Stop(gateway.context)
}

// Errors 获取网关处理器错误信息
func (h *GatewayServerHandler) Errors() []error {
	return h.errors
}
