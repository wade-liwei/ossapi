package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically

	// "github.com/minhblues/api/pkg/configs"
	"github.com/wade-liwei/ossapi/pkg/routes"
	// "github.com/minhblues/api/pkg/utils"
)

func main() {
	// Define Fiber config.

	//config := configs.FiberConfig()

	// Define a new Fiber app with config.
	//app := fiber.New(config)

	app := fiber.New()

	routes.PublicRoutes(app) // Register a public routes for app.

	log.Fatal(app.Listen(":13000"))

	// Start server (with graceful shutdown).
	//utils.StartServer(app)

}
