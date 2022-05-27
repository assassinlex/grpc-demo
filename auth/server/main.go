package main

import (
	authPb "cool_car/auth/api/v1/pb"
	auth "cool_car/auth/service"
	"cool_car/auth/service/token"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"google.golang.org/grpc"

	"go.uber.org/zap"
)

func main() {
	// 日志
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("create logger failed: %v", err)
	}

	// 监听
	lis, err := net.Listen("tcp", ":6666")
	if err != nil {
		logger.Fatal("create listener failed: ", zap.Error(err))
	}

	// 读取签名私钥
	keyFile, err := os.Open("/Users/luoxiao/Goland/cool_car/auth/server/private.key")
	if err != nil {
		logger.Fatal("无法打开私钥文件: ", zap.Error(err))
	}
	keyBytes, err := ioutil.ReadAll(keyFile)
	if err != nil {
		logger.Fatal("无法读取私钥文件: ", zap.Error(err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		logger.Fatal("无法解析私钥文件: ", zap.Error(err))
	}

	// 注册 grpc 服务
	server := grpc.NewServer()
	authPb.RegisterAuthServiceServer(server, &auth.Service{
		Logger:         logger,
		TokenExpire:    2 * time.Hour,
		TokenGenerator: token.NewJwtTokenGenerator("cool_cat/auth", privateKey),
	})

	// 启动服务
	err = server.Serve(lis)
	if err != nil {
		logger.Fatal("server start failed: ", zap.Error(err))
	}
}
