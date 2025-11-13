package common

import (
	"sync"

	"go.uber.org/zap"
)

type Logger struct {
	Log *zap.Logger
}

var mutex = &sync.Mutex{}

var loggerInstance *Logger

func GetLogger() *Logger {
	if loggerInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if loggerInstance == nil {
			zapLogger, _ := zap.NewProduction()
			loggerInstance = &Logger{
				Log: zapLogger,
			}
		}
	}
	return loggerInstance
}
