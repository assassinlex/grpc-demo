package rental

import (
	"context"
	rentalPb "cool_car/rental/api/v1/pb"
	"cool_car/shared/auth"

	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
}

// CreateTrip 获取行程
func (s *Service) CreateTrip(ctx context.Context, request *rentalPb.CreateTripRequest) (*rentalPb.CreateTripResponse, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	s.Logger.Info("create trip", zap.String("start", request.Start), zap.String("account_id", aid.String()))
	return nil, status.Error(codes.Unimplemented, "待实现")
}
