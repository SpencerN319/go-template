package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/SpencerN319/go-template/internal/pkg/server"
)

type PrettyJSONHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyJSONHandler) Handle(ctx context.Context, r slog.Record) error {
	fields := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	t := time.Now().Format(time.RFC3339)
	h.l.Println(t, r.Level.String(), r.Message, string(b))

	return nil
}

func NewPrettyJSONHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return &PrettyJSONHandler{slog.NewJSONHandler(w, opts), log.New(w, "", 0)}
}

func Mustgetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Errorf("missing env var: %s", key))
	}
	return v
}

func MustParseLogLevel(l string) slog.Level {
	switch l {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		panic(fmt.Errorf("unknown log level %s", l))
	}
}

func run(ctx context.Context, l *slog.Logger, args []string, getenv func(string) string) error {
	slog.SetDefault(l)

	httpServer := &http.Server{
		Addr:              getenv("ADDR"),
		Handler:           server.New(ctx),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	go func() {
		slog.Debug("http server listening", slog.String("addr", httpServer.Addr))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving on %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		cctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(cctx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()

	return nil
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	level := MustParseLogLevel(Mustgetenv("LOG_LEVEL"))
	opts := &slog.HandlerOptions{AddSource: level == slog.LevelDebug, Level: level}
	l := slog.New(NewPrettyJSONHandler(os.Stdout, opts))

	if err := run(ctx, l, os.Args, Mustgetenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
