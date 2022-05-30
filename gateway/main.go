package main

import (
	"context"
	authPb "cool_car/auth/api/v1/pb"
	rentalPb "cool_car/rental/api/v1/pb"
	"log"
	"net/http"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	// 创建上下文
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 创建 grpc 多路调用器
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:  true,
				UseEnumNumbers: true,
			},
		}))

	var err error

	// 注册 grpc 调用
	// auth(授权)
	err = authPb.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:6666",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
	if err != nil {
		log.Fatalf("register auth service failed: %v", err)
	}
	// rental(租赁)
	err = rentalPb.RegisterTripServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:6667",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		})
	if err != nil {
		log.Fatalf("register trip service failed: %v", err)
	}

	log.Fatal(http.ListenAndServe(":7777", mux))
}
