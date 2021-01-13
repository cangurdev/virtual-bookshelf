package repository

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"virtual-bookshelf/database"
	"virtual-bookshelf/model"
)

func GetUser(email string) (*gocb.QueryResult, error) {
	query := fmt.Sprintf("SELECT users.* FROM users WHERE email = '%s'", email)
	res, err := database.GetCluster().Query(query, nil)
	return res, err
}

func SaveUser(uuid string, document model.User) error {
	_, err := database.GetCollection().Insert(uuid, &document, nil)
	return err
}
