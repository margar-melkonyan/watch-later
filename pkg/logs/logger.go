package logs

import (
	"log/slog"
	"os"
	"strconv"
)

// Before use this pkg, you firstly should add env-loader to your project.
func Initialize() error {
	debugLevel, leverErr := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if leverErr != nil {
		return leverErr
	}

	loggerOptions := &slog.HandlerOptions{
		Level: slog.Level(debugLevel),
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, loggerOptions))
	slog.SetDefault(logger)

	return nil
}
