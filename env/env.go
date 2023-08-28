package env 

import "os"

func Getenv(key string, fallback string) string {
	e := os.Getenv(key)
	if e == "" {
		return fallback
	}
	return e
}