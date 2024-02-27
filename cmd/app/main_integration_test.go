//go:build integration

package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"
)

func testenv(key string) string {
	switch key {
	case "ADDR":
		return ":8080"
	default:
		fmt.Fprintf(os.Stderr, "getenv unknown key: %s", key)
		return ""
	}
}

func WaitForReady(ctx context.Context, timeout time.Duration, endpoint string) error {
	client := http.Client{}
	startTime := time.Now().UTC()
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		res, err := client.Do(req)
		if err != nil {
			continue
		}

		switch res.StatusCode {
		case http.StatusOK:
			res.Body.Close()
			return nil
		default:
			res.Body.Close()
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if time.Since(startTime) >= timeout {
				return fmt.Errorf("timeout reached while waiting for endpoint")
			}
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func TestRun(t *testing.T) {
	cctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	go func() {
		if err := run(cctx, slog.Default(), []string{}, testenv); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}()
	if err := WaitForReady(cctx, 5*time.Second, "http://localhost:8080"); err != nil {
		t.Fatal(err)
	}
}
