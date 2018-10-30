package logging

import (
	"os"
	"path/filepath"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

func serviceNameFromBin() string {
	programPath := os.Args[0]
	programName := filepath.Base(programPath)
	return programName
}

// SetupZapSDLogging - Setup logging with some defaults for GCP
func SetupZapSDLogging(devel bool) {
	var (
		logger *zap.Logger
		err    error
	)

	if devel {
		logger, err = zapdriver.NewDevelopment()
	} else {
		logger, err = zapdriver.NewProduction()
	}

	zap.ReplaceGlobals(logger)

	if err != nil {
		panic(err)
	}

}
