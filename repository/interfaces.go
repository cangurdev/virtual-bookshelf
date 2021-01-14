package repository

import "virtual-bookshelf/model"

type BookRepository interface {
	AddBook(email string, document model.Book) error
	GetBooks(id string) ([]model.Book, error)
	Bookmark(bookId, id, page string) error
	RemoveBook(userId, bookId string) error
}
type AuthRepository interface {
	GetUser(email string) (map[string]interface{}, error)
	SaveUser(uuid string, document model.User) error
}
