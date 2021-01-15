package router

import (
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/handler"
	"virtual-bookshelf/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Home
	app.Get("/home", middleware.Auth, handler.Index)      // Returns home page
	app.Post("/home", middleware.Auth, handler.PostIndex) // Handles search book operation

	// Register
	app.Get("/register", handler.GetRegister)   // Returns register page
	app.Post("/register", handler.PostRegister) // Handles register process

	// Login
	app.Get("/login", handler.GetLogin)   // Returns login page
	app.Post("/login", handler.PostLogin) // Handles login process

	// Logout
	app.Get("/logout", handler.Logout) // Handles logout process

	// Add
	app.Get("/add", handler.AddBook) // Handles adding book to bookshelf process

	// Remove
	app.Get("/remove/:bookId", handler.RemoveBook) // Handles removing book from bookshelf process

	// Read
	app.Get("/books/:bookId/pages/:pageNumber", handler.ReadBook) // Returns book page

	// Bookmark
	app.Post("/bookmark", handler.Bookmark) // Handles bookmark process
}
