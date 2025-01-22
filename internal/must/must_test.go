//go:build unit

package must

import (
	"log/slog"
	"testing"
)

func TestGetenv(t *testing.T) {
	tc := []struct {
		desc        string
		key         string
		expect      string
		expectPanic bool
	}{
		{
			desc:        "on success should return env var",
			key:         "success",
			expect:      "true",
			expectPanic: false,
		},
		{
			desc:        "on empty env var should panic",
			key:         "panic",
			expect:      "",
			expectPanic: true,
		},
	}

	for i, c := range tc {
		t.Setenv(c.key, c.expect)

		defer func() {
			if r := recover(); r == nil && c.expectPanic {
				t.Fatalf("[%d] %s: expected function to panic", i, c.desc)
			}
		}()

		if got := Getenv(c.key); got != c.expect {
			t.Errorf("[%d] %s: expected %s; got %s", i, c.desc, c.expect, got)
		}
	}
}

func TestParseSlogLevel(t *testing.T) {
	tc := []struct {
		desc        string
		level       string
		expect      slog.Level
		expectPanic bool
	}{
		{
			desc:        "on success should return slog level",
			level:       "DEBUG",
			expect:      slog.LevelDebug,
			expectPanic: false,
		},
		{
			desc:        "on success should return slog level",
			level:       "INFO",
			expect:      slog.LevelInfo,
			expectPanic: false,
		},
		{
			desc:        "on success should return slog level",
			level:       "WARN",
			expect:      slog.LevelWarn,
			expectPanic: false,
		},
		{
			desc:        "on success should return slog level",
			level:       "ERROR",
			expect:      slog.LevelError,
			expectPanic: false,
		},
		{
			desc:        "on failure should panic",
			level:       "HIGH",
			expect:      slog.LevelInfo,
			expectPanic: true,
		},
	}

	for i, c := range tc {
		defer func() {
			if r := recover(); r == nil && c.expectPanic {
				t.Fatalf("[%d] %s: expected function to panic", i, c.desc)
			}
		}()

		if got := ParseSlogLevel(c.level); got != c.expect {
			t.Errorf("[%d] %s: expected %s; got %s", i, c.desc, c.expect, got)
		}
	}
}
