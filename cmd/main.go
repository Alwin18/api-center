package main

import (
	"api-center/config"
	"api-center/database"
	"api-center/routes"
	"context"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config")
	}

	ctx := context.Background()

	// setup database
	db, err := database.Connect(cfg)
	if err != nil {
		panic("failed to connect database")
	}

	// Setup and run router
	r := routes.SetupRouter(ctx, db)
	r.Run(":9000")
}
