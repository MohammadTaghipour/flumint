package inout

import (
	"sync"

	"go.uber.org/zap"
)

var (
	l    *zap.SugaredLogger
	once sync.Once
)

func Logger() *zap.SugaredLogger {
	once.Do(func() {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic("Logger can not be initialized.")
		}
		defer logger.Sync()
		l = logger.Sugar()
	})
	return l
}
