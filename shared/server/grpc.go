package server

import (
	"cool_car/shared/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(server *grpc.Server)
	Logger            *zap.Logger
}

// RunGRPCServer grpc 服务注册 & 启动
func RunGRPCServer(cfg *GRPCConfig) error {
	name := zap.String("name", cfg.Name)
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		cfg.Logger.Fatal("create listener failed", name, zap.Error(err))
	}

	// 部分接口无需 auth 中间件, 比如登录
	var opts []grpc.ServerOption

	if cfg.AuthPublicKeyFile != "" {
		interceptor, err := auth.Interceptor(cfg.AuthPublicKeyFile)
		if err != nil {
			cfg.Logger.Fatal("create auth interceptor failed: ", name, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}

	server := grpc.NewServer(opts...)
	cfg.RegisterFunc(server)

	cfg.Logger.Info("server started")
	return server.Serve(lis)
}
