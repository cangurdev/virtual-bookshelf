package test

import (
	"github.com/stretchr/testify/mock"
	"virtual-bookshelf/model"
)

type MockBookRepository struct {
	mock.Mock
}

func (mock *MockBookRepository) AddBook(email string, document model.Book) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockBookRepository) GetBooks(id string) ([]model.Book, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Book), args.Error(1)
}

func (mock *MockBookRepository) GetBook(id, userId string) ([]string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]string), args.Error(1)
}

func (mock *MockBookRepository) Bookmark(bookId, id, page string) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockBookRepository) RemoveBook(userId, bookId string) error {
	args := mock.Called()
	return args.Error(0)
}
