package handler

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"virtual-bookshelf/repository"
	"virtual-bookshelf/service"
)

func GetLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
func PostLogin(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	authService := service.NewAuthService(repository.NewAuthRepository())
	err := authService.Login(email, password)
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
