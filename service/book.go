package service

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"virtual-bookshelf/model"
	"virtual-bookshelf/repository"
)

type bookService struct {
}

var bookRepository repository.BookRepository

func NewBookService(repository repository.BookRepository) BookService {
	bookRepository = repository
	return &bookService{}
}
func (*bookService) AddBook(id string, book model.Book) error {
	err := bookRepository.AddBook(id, book)
	if err != nil {
		return err
	}
	return nil
}

func (*bookService) ReadBook(id string) ([]string, error) {
	book, err := getBook(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
func getBook(id string) ([]string, error) {
	url := fmt.Sprintf("http://www.gutenberg.org/files/%s/%s-h/%s-h.htm", id, id, id)
	var paragraphs []string
	c := colly.NewCollector()
	c.OnHTML("p", func(e *colly.HTMLElement) {
		paragraphs = append(paragraphs, e.Text)
	})
	err := c.Visit(url)
	if err != nil {
		return nil, nil
	}
	return paragraphs, nil
}
func (*bookService) GetBookshelf(id string) ([]model.Book, error) {
	books, err := bookRepository.GetBooks(id)
	if err != nil {
		return books, err
	}
	return books, nil
}
func (*bookService) Bookmark(bookId, id, page string) error {
	err := bookRepository.Bookmark(bookId, id, page)
	if err != nil {
		return err
	}
	return nil
}
func (*bookService) RemoveBook(userId, bookId string) error {
	err := bookRepository.RemoveBook(userId, bookId)
	if err != nil {
		return err
	}
	return nil
}
