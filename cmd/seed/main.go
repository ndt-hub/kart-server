package main

import (
	"kart-server/config"
	"log"
)

func main() {
	log.Println("Starting database seeding...")

	config.LoadEnv()

	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := config.SeedDatabase(); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	log.Println("Database seeding completed successfully")
}
