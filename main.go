package main

import (
	"log"

	"github.com/gofibre/fiber/v2"
	"github.com/tanisboxed/alignspace/config"
	"github.com/tanisboxed/alignspace/router"
)

func main() {
	//loads config that calls the env file
	cfg := config.LoadConfig()

	//creates a new fibre instance (this is the web server)
	app := fiber.New()

	//calls router package to attach routes to the app
	router.SetupRoutes(app)

	//starts the server
	log.Printf("🚀 AlignSpace starting on port %s\n", cfg.Port)
	err := app.Listen(":" + cfg.Port)
	if err != nil {
		log.Fatalf("❌ Failed to start server: %v\n", err)
	}
}