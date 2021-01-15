package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"virtual-bookshelf/model"
	"virtual-bookshelf/service"
)

func TestReadBook(t *testing.T) {
	//Arrange
	mockRepo := new(MockBookRepository)
	bookService := service.NewBookService(mockRepo)

	//Act
	result, err := bookService.ReadBook("1661")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
	assert.NotNil(t, result)
}
func TestGetBookshelf(t *testing.T) {
	//Arrange
	mockRepo := new(MockBookRepository)
	bookService := service.NewBookService(mockRepo)
	var result []model.Book
	mockRepo.On("GetBooks").Return(result, nil)

	//Act
	_, err := bookService.GetBookshelf("1")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
}
func TestGetBookshelfOfInvalidId(t *testing.T) {
	mockRepo := new(MockBookRepository)
	bookService := service.NewBookService(mockRepo)
	var result []model.Book
	mockRepo.On("GetBooks").Return(result, errors.New("invalid id"))

	//Act
	_, err := bookService.GetBookshelf("1")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, "invalid id")
}
func TestBookmark(t *testing.T) {
	//Arrange
	mockRepo := new(MockBookRepository)
	bookService := service.NewBookService(mockRepo)
	mockRepo.On("Bookmark").Return(nil)

	//Act
	err := bookService.Bookmark("1212", "1", "30")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
}
func TestAddBook(t *testing.T) {
	//Arrange
	mockRepo := new(MockBookRepository)
	bookService := service.NewBookService(mockRepo)
	mockRepo.On("AddBook").Return(nil)
	var book model.Book
	book.Id = "1"
	book.Bookmark = "1"
	book.Image = "https://www.gutenberg.org/cache/epub/1661/pg1661.cover.medium.jpg"
	book.Title = "The Adventures of Sherlock Holmes by Arthur Conan Doyle"

	//Act
	err := bookService.AddBook("1212", book)

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
}
func TestRemoveBook(t *testing.T) {
	//Arrange
	mockRepo := new(MockBookRepository)
	bookService := service.NewBookService(mockRepo)
	mockRepo.On("RemoveBook").Return(nil)

	//Act
	err := bookService.RemoveBook("1", "1212")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
}
