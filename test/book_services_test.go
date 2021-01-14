package test

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"virtual-bookshelf/model"
	"virtual-bookshelf/service"
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

func TestAddBook(t *testing.T) {
	var c *fiber.Ctx
	mockRepo := new(MockBookRepository)
	mockRepo.On("AddBook").Return(nil)
	bookService := service.NewBookService(mockRepo)
	err := bookService.AddBook(c)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
}
func TestReadBook(t *testing.T) {
	mockRepo := new(MockBookRepository)
	var result []string
	mockRepo.On("GetBook").Return(result, nil)
	bookService := service.NewBookService(mockRepo)
	_, err := bookService.ReadBook("1", "1")
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
}
func TestGetBookshelf(t *testing.T) {
	mockRepo := new(MockBookRepository)
	var result []model.Book
	mockRepo.On("GetBooks").Return(result, nil)
	bookService := service.NewBookService(mockRepo)
	_, err := bookService.GetBookshelf("1")
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
}
func TestGetBookshelfOfInvalidId(t *testing.T) {
	mockRepo := new(MockBookRepository)
	bookService := service.NewBookService(mockRepo)
	var result []model.Book
	mockRepo.On("GetBooks").Return(result, errors.New("invalid id"))
	_, err := bookService.GetBookshelf("1")
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, "invalid id")
}
func TestBookmark(t *testing.T) {
	mockRepo := new(MockBookRepository)
	bookService := service.NewBookService(mockRepo)
	mockRepo.On("Bookmark").Return(nil)
	err := bookService.Bookmark("1212", "1", "30")
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
}
