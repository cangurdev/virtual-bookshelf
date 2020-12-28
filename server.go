package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"virtual-bookshelf/handler"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "./public")

	//api := app.Group("/api", handler.Index) // /
	app.Get("/", handler.Index)            // /login
	app.Get("/register", handler.Register) // /register
	app.Get("/home", handler.Home) // Home

	app.Listen(":3000")
}
