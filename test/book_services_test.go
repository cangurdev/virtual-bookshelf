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
	var result []string
	mockRepo.On("GetBook").Return(result, nil)

	//Act
	_, err := bookService.ReadBook("1", "1")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
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
