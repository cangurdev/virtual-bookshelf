package handler

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:  "username",
		Value: "",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
		SameSite: "lax",
	})
	return c.Redirect("/login")
}
