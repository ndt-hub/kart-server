package main

import (
	"kart-server/config"
	"log"
)

func main() {
	log.Println("Starting database migration...")

	config.LoadEnv()

	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := config.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
}
