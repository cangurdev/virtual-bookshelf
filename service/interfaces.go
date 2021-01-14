package service

import (
	"github.com/gofiber/fiber/v2"
	"virtual-bookshelf/model"
)

type BookService interface {
	AddBook(c *fiber.Ctx) error
	ReadBook(userId, id string) ([]string, error)
	GetBookshelf(id string) ([]model.Book, error)
	Bookmark(bookId, id, page string) error
	RemoveBook(userId, bookId string) error
}
type AuthService interface {
	Login(email, password string) (string, error)
	Register(email, password string) (string, error)
}
