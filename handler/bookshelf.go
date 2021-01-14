package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"virtual-bookshelf/model"
)

func AddBook(c *fiber.Ctx) error {
	book := model.Book{}
	book.Id = c.Query("id")
	book.Title = c.Query("title")
	book.Subtitle = c.Query("subtitle")
	book.Description = c.Query("description")
	book.Image = c.Query("image")
	book.Url = c.Query("url")
	book.Bookmark = "1"
	book.File = fmt.Sprintf("http://www.gutenberg.org/files/%s/%s-h/%s-h.htm", book.Id, book.Id, book.Id)
	id := c.Cookies("username")
	err := bookService.AddBook(id, book)
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
	book, err := bookService.ReadBook(bookId)
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
