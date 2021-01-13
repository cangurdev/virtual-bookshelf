package handler

import (
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/repository"
	"virtual-bookshelf/service"
)

func AddBook(c *fiber.Ctx) error {
	bookService := service.NewBookService(repository.NewBookRepository())
	bookService.AddBook(c)
	return c.Redirect("/home")
}
func RemoveBook(c *fiber.Ctx) error {
	//bookId := c.Params("bookId")
	return c.Redirect("/home")
} /*
func ReadBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	book := service.ReadBook(bookId)
	return c.Render("index", fiber.Map{
		"Book": book,
	})
}*/
