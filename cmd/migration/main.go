package main

import (
	migrationDown "donTecoTest/cmd/migration/down"
	migrationUp "donTecoTest/cmd/migration/up"
	"donTecoTest/config"
	"donTecoTest/pkg/repository"
	"flag"
	"fmt"
	"log"
)

func main() {
	upFlag := flag.Bool("up", false, "up migration")
	downFlag := flag.Bool("down", false, "down migration")
	flag.Parse()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("can't parse config: %s", err)
		return
	}
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
		return
	}
	if *upFlag && *downFlag {
		fmt.Printf("Got up and down flags, please use only one flag to use migration\n")
		return
	}

	if *upFlag {
		err := migrationUp.Run(db)
		if err != nil {
			log.Fatalf("failed migration up: %s \n", err)
			return
		}
	}

	if *downFlag {
		err := migrationDown.Run(db)
		if err != nil {
			log.Fatalf("failed migration down: %s", err)
			return
		}
	}
}
