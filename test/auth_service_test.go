package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"virtual-bookshelf/service"
)

func TestLoginUSer(t *testing.T) {
	//Arrange
	mockRepo := new(MockRepository)
	testService := service.NewAuthService(mockRepo)
	user := map[string]interface{}{"id": "123", "email": "can@gmail.com", "password": "$2a$10$rINQfxKohoAq74Rj7eSD.O1PY2fMu48CXeyYT9mQLw.h3SGkOjDzi"}
	mockRepo.On("GetUser").Return(user, nil)

	//Act
	id, err := testService.Login("can@gmail.com", "123456")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
	assert.Equal(t, "123", id)
}
func TestLoginInvalidUser(t *testing.T) {
	//Arrange
	mockRepo := new(MockRepository)
	testService := service.NewAuthService(mockRepo)
	var user map[string]interface{}
	mockRepo.On("GetUser").Return(user, errors.New("no result was available"))

	//Act
	id, err := testService.Login("deneme@gmail.com", "123456")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, "no result was available")
	assert.Equal(t, "", id)
}
func TestRegisterUser(t *testing.T) {
	//Arrange
	mockRepo := new(MockRepository)
	testService := service.NewAuthService(mockRepo)
	mockRepo.On("SaveUser").Return(nil)

	//Act
	id, err := testService.Register("deneme@gmail.com", "123456")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
	assert.NotNil(t, id)
}
func TestRegisterEmptyPasswordUser(t *testing.T) {
	//Arrange
	mockRepo := new(MockRepository)
	testService := service.NewAuthService(mockRepo)
	mockRepo.On("SaveUser").Return(errors.New("invalid user"))

	//Act
	_, err := testService.Register("deneme@gmail.com", "")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, "invalid user")
}
func TestRegisterEmptyEmailUser(t *testing.T) {
	//Arrange
	mockRepo := new(MockRepository)
	testService := service.NewAuthService(mockRepo)
	mockRepo.On("SaveUser").Return(errors.New("invalid user"))

	//Act
	_, err := testService.Register("", "123456")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, "invalid user")
}
func TestRegisterEmptyUser(t *testing.T) {
	//Arrange
	mockRepo := new(MockRepository)
	testService := service.NewAuthService(mockRepo)
	mockRepo.On("SaveUser").Return(errors.New("invalid user"))

	//Act
	_, err := testService.Register("", "")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, "invalid user")
}
