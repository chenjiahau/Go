package util

import (
	"go.uber.org/zap"
)

func WriteInfoLog(message string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info(message)
}

func WriteWarnLog(message string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Warn(message)
}

func WriteErrorLog(message string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Error(message)
}

func WriteDebugLog(message string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Debug(message)
}

