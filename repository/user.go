package repository

import (
	"fmt"
	"virtual-bookshelf/database"
	"virtual-bookshelf/model"
)

type AuthRepository interface {
	GetUser(email string) (map[string]interface{}, error)
	SaveUser(uuid string, document model.User) error
}
type repository struct {
}

func NewAuthRepository() AuthRepository {
	return &repository{}
}
func (*repository) GetUser(email string) (map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT users.* FROM users WHERE email = '%s'", email)
	res, err := database.GetCluster().Query(query, nil)
	var user map[string]interface{}
	err = res.One(&user)
	if err != nil {
		return nil, err
	}
	err = res.Err()
	if err != nil {
		return nil, err
	}
	return user, err
}
func (repository) SaveUser(uuid string, document model.User) error {
	_, err := database.GetCollection().Insert(uuid, &document, nil)
	return err
}
