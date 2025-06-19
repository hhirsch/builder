package environment

import (
	"log/slog"
	"os"
)

func SetupLoggingForTesting() {
	slog.Info("Setting up logging for testing.")
	opts := &slog.HandlerOptions{
		AddSource: true,
		//Level:     slog.LevelInfo,
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))
	slog.SetDefault(logger)
}

func SetupLoggingForProduction(fileName string) {
	slog.Info("Setting up logging for production.")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("error opening file using logging for testing", slog.String("errorMessage", err.Error()))
		SetupLoggingForTesting()
		return
	}
	defer file.Close()
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}
	logger := slog.New(slog.NewJSONHandler(file, opts))
	slog.SetDefault(logger)
}
