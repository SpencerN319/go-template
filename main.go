package main

import (
	"github.com/SpencerN319/go-template/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	logLevel, _ := zerolog.ParseLevel(env.Getenv("LOG_LEVEL", "0"))
	zerolog.SetGlobalLevel(logLevel)
}

func main() {
	log.Print(hello())
}