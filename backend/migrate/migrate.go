package main

import (
	"log"
	"mcp_demo/config"
	"mcp_demo/models/entities"
)

func main() {
	// Initialize environment and database
	config.LoadEnvVariables()
	db := config.ConnectDatabase()

	// Cleanup database connections before exit
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Run migrations
	err = db.AutoMigrate(
		&entities.Note{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Migrations completed successfully")
}
