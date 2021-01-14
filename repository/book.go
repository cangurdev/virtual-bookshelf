package repository

import (
	"github.com/couchbase/gocb/v2"
	"virtual-bookshelf/database"
	"virtual-bookshelf/model"
)

type bookRepository struct {
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}
func (*bookRepository) AddBook(id string, document model.Book) error {
	updateGetResult, err := database.GetCollection().Get(id, nil)
	if err != nil {
		return err
	}
	var doc model.User
	err = updateGetResult.Content(&doc)
	if err != nil {
		return err
	}
	doc.Books = append(doc.Books, document)
	_, err = database.GetCollection().Replace(id, doc, &gocb.ReplaceOptions{
		Cas: updateGetResult.Cas()})
	return nil
}
func (*bookRepository) GetBooks(id string) ([]model.Book, error) {
	var doc model.User
	updateGetResult, err := database.GetCollection().Get(id, nil)
	if err != nil {
		return nil, err
	}
	err = updateGetResult.Content(&doc)
	if err != nil {
		return nil, err
	}
	return doc.Books, nil
}
func (*bookRepository) Bookmark(bookId, id, page string) error {
	updateGetResult, err := database.GetCollection().Get(id, nil)
	if err != nil {
		return err
	}
	var doc model.User
	err = updateGetResult.Content(&doc)
	if err != nil {
		return err
	}
	books := doc.Books
	for i := range books {
		if books[i].Id == bookId {
			books[i].Bookmark = page
			break
		}
	}
	doc.Books = books
	_, err = database.GetCollection().Replace(id, doc, &gocb.ReplaceOptions{
		Cas: updateGetResult.Cas()})
	if err != nil {
		return err
	}
	return nil
}
func (*bookRepository) RemoveBook(userId, bookId string) error {
	var doc model.User
	updateGetResult, err := database.GetCollection().Get(userId, nil)
	if err != nil {
		return err
	}
	err = updateGetResult.Content(&doc)
	var books []model.Book
	for _, book := range doc.Books {
		if book.Id != bookId {
			books = append(books, book)
		}
	}
	doc.Books = books
	_, err = database.GetCollection().Replace(userId, doc, &gocb.ReplaceOptions{
		Cas: updateGetResult.Cas()})
	if err != nil {
		return err
	}
	return nil
}
