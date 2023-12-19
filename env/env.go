package env

import (
	"fmt"
	"os"
)

func FallbackGetenv(key string, fallback string) string {
	e := os.Getenv(key)
	if e == "" {
		return fallback
	}
	return e
}

func MustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Errorf("error environment variable %s required", key))
	}
	return v
}
