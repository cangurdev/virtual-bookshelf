package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {

	// Get value
	name := c.Cookies("username")

	if name != "" {
		return c.Next()
	}
	return c.Redirect("/login")
}
