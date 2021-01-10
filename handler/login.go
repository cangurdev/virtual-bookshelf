package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
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
	res, err := database.GetCluster().Query(query, nil)

	if err != nil {
		fmt.Print("Error")
	}

	var user map[string]interface{}
	err = res.One(&user)
	if err != nil {
		fmt.Print(err)
	}
	err = res.Err()
	if err != nil {
		fmt.Print(err)
	}

	hashedPassword, _ := json.Marshal(user["password"])
	s, _ := strconv.Unquote(string(hashedPassword))

	err = bcrypt.CompareHashAndPassword([]byte(s), []byte(password))
	if err != nil {
		return c.Redirect("/login")
	}
	c.Cookie(&fiber.Cookie{
		Name:     "username",
		Value:    email,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})
	return c.Redirect("/home")

}
