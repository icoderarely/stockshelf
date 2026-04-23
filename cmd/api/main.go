package main

import (
	"fmt"
	"log"

	"github.com/icoderarely/stockshelf/internal/config"
	"github.com/icoderarely/stockshelf/internal/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg := config.New()

	dbpool, err := db.NewPool(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	_ = dbpool
	fmt.Println("DB Connected")
}
