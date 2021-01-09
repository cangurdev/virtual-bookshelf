package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"virtual-bookshelf/database"
	"virtual-bookshelf/handler"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./assets/images")
	app.Use(logger.New())
	group := app.Group("/")
	store := session.New(session.Config{CookiePath: "/", CookieDomain: "localhost"})

	// This panic will be catch by the middleware
	group.Get("/", func(c *fiber.Ctx) error {
		// get session from storage
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}
		// Get value
		name := sess.Get("name")
		fmt.Printf("Welcome %v\n", name)
		if name != nil {
			return c.Next()
		}
		return c.Redirect("/login")
	})
	// Get methods
	app.Get("/home", handler.Index)           // /
	app.Get("/register", handler.GetRegister) // /register
	app.Get("/login", handler.GetLogin)       // /login

	// Post methods
	app.Post("/login", func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		query := fmt.Sprintf("SELECT users.* FROM users WHERE email = '%s'", email)
		user, err := database.GetCluster().Query(query, nil)

		if err != nil {
			fmt.Print("Error")
		}
		var user1 map[string]interface{}
		err = user.One(&user1)
		if err != nil {
			fmt.Print(err)
		}

		err = user.Err()
		if err != nil {
			fmt.Print(err)
		}
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		// Set key/value
		sess.Set("name", email)
		fmt.Printf("name %v", sess.Get("name"))
		// save session
		defer sess.Save()
		if user1["password"] == password {
			return c.Redirect("/home")
		}
		return c.Redirect("/login")
	}) // /login
	app.Post("/register", handler.PostRegister) // /register
	app.Post("/home", handler.PostIndex)
	// Listens server on port 3000
	app.Listen(":3000")
}
