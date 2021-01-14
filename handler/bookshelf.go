package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AddBook(c *fiber.Ctx) error {

	err := bookService.AddBook(c)
	if err != nil {
		return err
	}
	return c.Redirect("/home")
}
func RemoveBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	userId := c.Cookies("username")
	err := bookService.RemoveBook(userId, bookId)
	if err != nil {
		return err
	}
	return c.Redirect("/home")
}
func ReadBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	page, _ := strconv.Atoi(c.Params("pageNumber"))
	book, err := bookService.ReadBook(c.Cookies("username"), bookId)
	if err != nil {
		return err
	}
	nextPage := page + 1
	previousPage := page - 1
	book = book[(page*10)-9 : (page * 10)]
	return c.Render("book", fiber.Map{
		"Book":         book,
		"Page":         page,
		"BookId":       bookId,
		"nextPage":     nextPage,
		"previousPage": previousPage,
	})
}
func Bookmark(c *fiber.Ctx) error {
	page := c.Query("p")
	bookId := c.Query("b")
	id := c.Cookies("username")
	err := bookService.Bookmark(bookId, id, page)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("/books/%s/pages/%s", bookId, page)
	return c.Redirect(url)
}
