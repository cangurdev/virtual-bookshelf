package handler

import (
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/service"
)

func GetRegister(c *fiber.Ctx) error {
	return c.Render("sign_up", fiber.Map{
		"Title": "register",
	})
}
func PostRegister(c *fiber.Ctx) error {
	err := service.Register(c)
	if err != nil {
		return c.Redirect("/register")
	}
	return c.Redirect("/home")
}
