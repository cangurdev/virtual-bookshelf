package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"virtual-bookshelf/model"
	"virtual-bookshelf/service"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetUser(email string) (map[string]interface{}, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(map[string]interface{}), args.Error(1)
}

func (mock *MockRepository) SaveUser(uuid string, document model.User) error {
	args := mock.Called()
	return args.Error(0)
}

func TestLoginUSer(t *testing.T) {
	mockRepo := new(MockRepository)
	user := map[string]interface{}{"email": "can@gmail.com", "password": "$2a$10$rINQfxKohoAq74Rj7eSD.O1PY2fMu48CXeyYT9mQLw.h3SGkOjDzi"}
	mockRepo.On("GetUser").Return(user, nil)
	testService := service.NewAuthService(mockRepo)
	result := testService.Login("can@gmail.com", "123456")

	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, result)
}
func TestLoginInvalidUser(t *testing.T) {
	mockRepo := new(MockRepository)
	var user map[string]interface{}
	mockRepo.On("GetUser").Return(user, errors.New("no result was available"))
	testService := service.NewAuthService(mockRepo)
	result := testService.Login("deneme@gmail.com", "123456")

	mockRepo.AssertExpectations(t)
	assert.EqualError(t, result, "no result was available")
}
func TestRegisterUser(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("SaveUser").Return(nil)
	testService := service.NewAuthService(mockRepo)
	result := testService.Register("deneme@gmail.com", "123456")
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, result)
}
func TestRegisterEmptyPasswordUser(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("SaveUser").Return(errors.New("invalid user"))
	testService := service.NewAuthService(mockRepo)
	result := testService.Register("deneme@gmail.com", "")
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, result, "invalid user")
}
func TestRegisterEmptyEmailUser(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("SaveUser").Return(errors.New("invalid user"))
	testService := service.NewAuthService(mockRepo)
	result := testService.Register("", "123456")
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, result, "invalid user")
}
func TestRegisterEmptyUser(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("SaveUser").Return(errors.New("invalid user"))
	testService := service.NewAuthService(mockRepo)
	result := testService.Register("", "")
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, result, "invalid user")
}
