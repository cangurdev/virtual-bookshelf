package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
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

	err = user.Err()
	if err != nil {
		fmt.Print(err)
	}

	if user1["password"] == password {
		c.Cookie(&fiber.Cookie{
			Name:  "username",
			Value: email,
			// Set expiry date to the past
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			SameSite: "lax",
		})
		return c.Redirect("/home")
	}
	return c.Redirect("/login")
}
