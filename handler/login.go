package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"virtual-bookshelf/database"
)

func GetLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
func PostLogin(c *fiber.Ctx) error {
	store := session.New()

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
	if user1["password"] == password {
		return c.Redirect("/")
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

	return c.Redirect("/")
}
