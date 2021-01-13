package handler

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"virtual-bookshelf/repository"
	"virtual-bookshelf/service"
)

func GetRegister(c *fiber.Ctx) error {
	return c.Render("sign_up", fiber.Map{
		"Title": "register",
	})
}
func PostRegister(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	authService := service.NewAuthService(repository.NewAuthRepository())
	err := authService.Register(email, password)
	if err != nil {
		return c.Redirect("/register")
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
