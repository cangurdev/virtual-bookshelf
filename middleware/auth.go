package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Auth(c *fiber.Ctx) error{
	store := session.New()
	sess, err := store.Get(c)

	if err != nil {
		panic(err)
	}
	defer sess.Save()
	// Get value
	name := sess.Get("name")
	if name == nil{
		return c.Next()
	}
	return nil
}