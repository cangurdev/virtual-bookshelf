package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {

	// Get value
	name := c.Cookies("username")

	fmt.Printf("Welcome %v\n", name)
	if name != "" {
		return c.Next()
	}
	return c.Redirect("/login")
}
