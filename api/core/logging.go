package core

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func SetupLogger() {
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		panic("Failed to setup logger")
	}
}
