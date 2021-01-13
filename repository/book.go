package repository

import (
	"github.com/couchbase/gocb/v2"
	"virtual-bookshelf/database"
	"virtual-bookshelf/model"
)

type BookRepository interface {
	AddBook(email string, document model.Book) error
}
type bookRepository struct {
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}
func (*bookRepository) AddBook(id string, document model.Book) error {
	updateGetResult, err := database.GetCollection().Get(id, nil)
	if err != nil {
		return err
	}
	var doc model.User
	err = updateGetResult.Content(&doc)
	if err != nil {
		return err
	}
	doc.Books = append(doc.Books, document)
	_, err = database.GetCollection().Replace(id, doc, &gocb.ReplaceOptions{
		Cas: updateGetResult.Cas()})
	return nil
}

/*
func (*bookRepository) GetBook(id string) (model.Book, error) {
	query := fmt.Sprintf("SELECT users.* FROM users WHERE email = '%s'", id)
	res, err := database.GetCluster().Query(query, nil)
	var book map[string]interface{}
	err = res.One(&user)
	if err != nil {
		return nil, err
	}
	err = res.Err()
	if err != nil {
		return nil, err
	}
	return user, err
}*/
