package logging

import (
	"os"
	"path/filepath"

	zapStackdriver "github.com/tommy351/zap-stackdriver"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func serviceNameFromBin() string {
	programPath := os.Args[0]
	programName := filepath.Base(programPath)
	return programName
}

// SetupZapSDLogging - Setup logging with some defaults for GCP
func SetupZapSDLogging() {
	config := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Encoding:         "json",
		EncoderConfig:    zapStackdriver.EncoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return &zapStackdriver.Core{
			Core: core,
		}
	}), zap.Fields(
		zapStackdriver.LogServiceContext(&zapStackdriver.ServiceContext{
			Service: serviceNameFromBin(),
			Version: "master",
		}),
	))

	zap.ReplaceGlobals(logger)

	if err != nil {
		panic(err)
	}
}
