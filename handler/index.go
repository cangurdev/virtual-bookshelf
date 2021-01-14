package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/repository"
	"virtual-bookshelf/service"
)

var bookService = service.NewBookService(repository.NewBookRepository())

func Index(c *fiber.Ctx) error {
	id := c.Cookies("username")
	books, err := bookService.GetBookshelf(id)
	if err != nil {
		fmt.Print(err)
	}
	return c.Render("index", fiber.Map{
		"Books": books,
	})
}
func PostIndex(c *fiber.Ctx) error {
	query := c.FormValue("query")
	books := service.Search(query)
	return c.Render("result", fiber.Map{
		"Books": books,
	})
}
