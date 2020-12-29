package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/database"
)

func GetLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
func PostLogin(c *fiber.Ctx) error {
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
	fmt.Printf("for %s", user1)

	if user1["password"] == password {
		fmt.Print(user1)
		return c.Redirect("/")
	}
	err = user.Err()
	if err != nil {
		fmt.Print(err)
	}
	return c.Redirect("/login")
}
