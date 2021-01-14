package test

import (
	"github.com/stretchr/testify/mock"
	"virtual-bookshelf/model"
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
