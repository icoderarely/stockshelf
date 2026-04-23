package main

import (
	"log"

	"github.com/icoderarely/stockshelf/internal/config"
	"github.com/icoderarely/stockshelf/internal/pkg/db"
)

func main() {
	cfg := config.New()

	dbpool, err := db.NewPool(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

}
