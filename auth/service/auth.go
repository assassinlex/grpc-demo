package auth

import (
	"context"
	authPb "cool_car/auth/api/v1/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.uber.org/zap"
)

// TokenGenerator token 生成器
type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

// Service 授权服务
type Service struct {
	Logger         *zap.Logger
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
}

// Login 登录
func (s *Service) Login(ctx context.Context, request *authPb.LoginRequest) (*authPb.LoginResponse, error) {
	s.Logger.Info("request params: " + request.String())

	// todo:: 数据库获取 account_id
	accountID := "SEcE8367E5hXC7CU"

	token, err := s.TokenGenerator.GenerateToken(accountID, s.TokenExpire)
	if err != nil {
		// 日志系统错误信息
		s.Logger.Error("token 生成失败: %v", zap.Error(err))
		// 用户展示错误信息
		return nil, status.Error(codes.Internal, "服务暂不可用, 请稍后再试.")
	}

	return &authPb.LoginResponse{
		AccessToken: token,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}
