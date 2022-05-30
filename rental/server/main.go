package main

import (
	rentalPb "cool_car/rental/api/v1/pb"
	rental "cool_car/rental/service"
	"cool_car/shared/server"
	"log"

	"google.golang.org/grpc"

	"go.uber.org/zap"
)

func main() {
	// 日志
	logger, err := server.NewLogger()
	if err != nil {
		log.Fatalf("create logger failed: %v", err)
	}

	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              ":6667",
		AuthPublicKeyFile: "/Users/luoxiao/Goland/cool_car/shared/auth/public.key",
		RegisterFunc: func(server *grpc.Server) {
			rentalPb.RegisterTripServiceServer(server, &rental.Service{
				Logger: logger,
			})
		},
		Logger: logger,
	})

	if err != nil {
		logger.Fatal("server start failed: ", zap.Error(err))
	}
}
