package logger

import (
	configs "myblog/configs"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	SetDebugMode(configs.IsDevelopment())
}

func SetDebugMode(isDebug bool) {
	if isDebug {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
