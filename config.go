package tint

import (
	"log/slog"
	"os"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05.000000"
)

func ConfigFromEnv(args ...any) {
	stdout := os.Stdout
	opts := Options{
		AddSource:  true,
		Level:      slog.LevelInfo,
		TimeFormat: TIME_FORMAT,
	}

	switch os.Getenv("LOG_SOURCE") {
	case "false", "FALSE", "0":
		opts.AddSource = false
	}

	switch os.Getenv("LOG_LEVEL") {
	case "debug", "DEBUG":
		opts.Level = slog.LevelDebug
	case "info", "INFO":
		opts.Level = slog.LevelInfo
	case "warn", "WARN":
		opts.Level = slog.LevelWarn
	case "error", "ERROR":
		opts.Level = slog.LevelError
	}

	colors := true
	switch os.Getenv("NO_COLOR") {
	case "true", "TRUE", "1":
		colors = false
	}
	opts.NoColor = !colors

	handler := NewHandler(stdout, &opts)
	logger := slog.New(handler).With(args...)
	slog.SetDefault(logger)
}
