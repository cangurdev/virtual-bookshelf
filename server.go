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
	app.Static("/", "./assets/images")

	app.Get("/", handler.Index)                 // /
	app.Get("/register", handler.GetRegister)   // /register
	app.Get("/login", handler.Login)            // /login
	app.Post("/register", handler.PostRegister) // /login

	app.Listen(":3000")
}
