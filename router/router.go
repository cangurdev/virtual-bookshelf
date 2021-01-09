package router

import (
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/handler"
	"virtual-bookshelf/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Home
	app.Get("/home", middleware.Auth, handler.Index)      // Home page
	app.Post("/home", middleware.Auth, handler.PostIndex) // Handles search book operation

	// Register
	app.Get("/register", handler.GetRegister)   // Register page
	app.Post("/register", handler.PostRegister) // Handles register process

	// Login
	app.Get("/login", handler.GetLogin)   // Login page
	app.Post("/login", handler.PostLogin) // Handles login process

	// Logout
	app.Get("/logout", handler.Logout) // Handles logout process

}
