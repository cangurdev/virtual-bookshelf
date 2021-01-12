package handler

import (
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/service"
)

func GetLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}
func PostLogin(c *fiber.Ctx) error {
	err := service.Login(c)
	if err != nil {
		return c.Redirect("/login")
	}
	return c.Redirect("/home")

}
