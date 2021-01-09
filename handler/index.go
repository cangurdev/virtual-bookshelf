package handler

import (
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/service"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Index",
	})
}
func PostIndex(c *fiber.Ctx) error {
	query := c.FormValue("query")
	books := service.Search(query)
	return c.Render("index", fiber.Map{
		"books": books,
	})
}
