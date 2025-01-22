package must

import (
	"fmt"
	"log/slog"
	"os"
)

func Getenv(s string) string {
	v := os.Getenv(s)
	if v == "" {
		panic(fmt.Errorf("unknown env var: %s", s))
	}
	return v
}

func ParseSlogLevel(s string) slog.Level {
	switch s {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		panic(fmt.Errorf("unknown log level %s", s))
	}
}
