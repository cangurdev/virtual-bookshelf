package service

import (
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/model"
	"virtual-bookshelf/repository"
)

type BookService interface {
	AddBook(c *fiber.Ctx) error
	ReadBook(id string) error
	GetBook(id string) error
}
type bookService struct {
}

var bookRepository repository.BookRepository

func NewBookService(repository repository.BookRepository) BookService {
	bookRepository = repository
	return &bookService{}
}
func (*bookService) AddBook(c *fiber.Ctx) error {
	book := model.Book{}
	book.Id = c.Query("id")
	book.Title = c.Query("title")
	book.Subtitle = c.Query("subtitle")
	book.Description = c.Query("description")
	book.Image = c.Query("image")
	book.Url = c.Query("url")
	book.Bookmark = "1"
	id := c.Cookies("username")
	err := bookRepository.AddBook(id, book)
	if err != nil {
		return err
	}
	return nil
}
func (*bookService) ReadBook(id string) error {
	return nil
}
func (*bookService) GetBook(id string) error {

	return nil
}
