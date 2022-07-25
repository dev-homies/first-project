package core

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger
var RawLogger *zap.Logger

func SetupLogger() {
	var err error
	RawLogger, err = zap.NewDevelopment()
	if err != nil {
		panic("Failed to setup logger")
	}

	Logger = RawLogger.Sugar()
}
