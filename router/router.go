package router

import (
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/handler"
	"virtual-bookshelf/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/home", middleware.Auth, handler.Index) // /
	app.Get("/register", handler.GetRegister)        // /register
	app.Get("/login", handler.GetLogin)              // /login

	// Post methods
	app.Post("/login", handler.PostLogin)       // /login
	app.Post("/register", handler.PostRegister) // /register
}
