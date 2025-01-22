package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/SpencerN319/go-template/internal/must"
	"github.com/SpencerN319/go-template/internal/server"
)

func run(ctx context.Context, l *slog.Logger, getenv func(string) string) error {
	slog.SetDefault(l)
	var wg sync.WaitGroup

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

	level := must.ParseSlogLevel(must.Getenv("LOG_LEVEL"))
	opts := &slog.HandlerOptions{AddSource: level == slog.LevelDebug, Level: level}
	l := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	if err := run(ctx, l, must.Getenv); err != nil {
		slog.Error("run error", slog.String("error", err.Error()))
		panic(err)
	}
}
