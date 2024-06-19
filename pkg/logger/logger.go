package logger

import (
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger

var lock = &sync.Mutex{}

func getLogger() *zap.Logger {
	if logger == nil {
		lock.Lock()
		defer lock.Unlock()
		if logger == nil {
			logger, _ = zap.NewProduction()
		}
	}
	return logger
}

func Info(msg string, fields ...zap.Field) {
	defer getLogger().Sync()
	getLogger().Info(msg, fields...)
}
