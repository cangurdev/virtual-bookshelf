package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"virtual-bookshelf/router"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve files from "./assets/images" directory
	app.Static("/", "./assets/images")

	router.SetupRoutes(app)

	// Listens server on port 3000
	err := app.Listen(":3000")
	if err != nil {
		fmt.Print("There was an error occurred on listening port")
	}
}
