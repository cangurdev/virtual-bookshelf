package handler

import "github.com/gofiber/fiber/v2"

func AddBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	return c.Redirect("/home")
}
func RemoveBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	return c.Redirect("/home")
}
