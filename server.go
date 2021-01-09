package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"virtual-bookshelf/handler"
	"virtual-bookshelf/middleware"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./assets/images")

	// Get methods
	app.Get("/home", middleware.Auth, handler.Index)
	app.Get("/register", handler.GetRegister) // /register
	app.Get("/login", handler.GetLogin)       // /login

	// Post methods
	app.Post("/login", handler.PostLogin)       // /login
	app.Post("/register", handler.PostRegister) // /register
	app.Post("/home", middleware.Auth, handler.PostIndex)
	app.Get("/logout", handler.Logout)
	// Listens server on port 3000
	app.Listen(":3000")
}
