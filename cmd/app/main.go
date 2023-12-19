package main

import (
	"log/slog"
	"os"

	"github.com/SpencerN319/go-template/env"
	"github.com/SpencerN319/go-template/hello"
)

func init() {
	parseLogLevel := func(l string) slog.Level {
		switch l {
		default:
			return slog.LevelInfo
		}
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: parseLogLevel(env.FallbackGetenv("LOG_LEVEL", "INFO")) == slog.LevelDebug,
		Level:     parseLogLevel(env.FallbackGetenv("LOG_LEVEL", "INFO")),
	})))
}

func main() {
	slog.Info(hello.Hello())
}
