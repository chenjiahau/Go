package util

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLog *zap.Logger

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "caller",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}

	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	config := zap.Config{
		 Level: atom,
		 Development: true,
		//  Encoding: "json",
		 Encoding: "console",
		 EncoderConfig: encoderConfig,
		 OutputPaths: []string{"stdout",},
		 ErrorOutputPaths: []string{"stderr"},
	}

	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder

	var err error
	ZapLog, err = config.Build()
	if err != nil {
		panic(fmt.Sprintf("can't initialize zap logger: %v", err))
	}
}

func WriteInfoLog(message string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ZapLog.Info(message)
}

func WriteWarnLog(message string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ZapLog.Warn(message)
}

func WriteErrorLog(message string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ZapLog.Error(message)
}