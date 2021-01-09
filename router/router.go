package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"virtual-bookshelf/handler"
)

func SetupRoutes(app *fiber.App)  {
	// Middleware
	app.Get("/", logger.New(),handler.GetLogin)

	app.Get("/", handler.Index)                 // /
	app.Get("/register", handler.GetRegister)   // /register
	app.Get("/login", handler.GetLogin)         // /login

	// Post methods
	app.Post("/login", handler.PostLogin)       // /login
	app.Post("/register", handler.PostRegister) // /register
}
