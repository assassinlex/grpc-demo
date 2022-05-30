package server

import (
	"go.uber.org/zap"
)

// NewLogger 获取日志器
func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()

	if err != nil {
		return nil, err
	}

	return logger, nil
}
