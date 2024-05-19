package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
)

func SetupLog(config *environment.Config) {
	var level slog.Level
	var handler slog.Handler

	logLevel := "info"

	if config.ApiConfig.IsDevelopment() {
		logLevel = "debug"
	}

	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		panic(fmt.Errorf("unable to load log level: %v", err))
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler = slog.NewJSONHandler(os.Stdout, opts)

	if config.ApiConfig.IsDevelopment() {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	log := slog.New(handler)
	slog.SetDefault(log)
}
