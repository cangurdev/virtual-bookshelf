package service

import (
	"virtual-bookshelf/model"
)

type BookService interface {
	AddBook(id string, book model.Book) error
	ReadBook(id string) ([]string, error)
	GetBookshelf(id string) ([]model.Book, error)
	Bookmark(bookId, id, page string) error
	RemoveBook(userId, bookId string) error
}
type UserService interface {
	Login(email, password string) (string, error)
	Register(email, password string) (string, error)
}
