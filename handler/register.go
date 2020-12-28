package handler

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	return c.Render("sign_up", fiber.Map{
		"Title": "register",
	})
}
