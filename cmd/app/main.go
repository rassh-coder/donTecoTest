package main

import (
	"donTecoTest/config"
	"donTecoTest/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("can't parse config: %s", err)
		return
	}

	err = app.Run(cfg)
	if err != nil {
		log.Fatalf("can't run server: %s", err)
		return
	}
}
