package service

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"virtual-bookshelf/model"
	"virtual-bookshelf/repository"
)

type AuthService interface {
	Login(email, password string) (string, error)
	Register(email, password string) (string, error)
}
type service struct {
}

var repo repository.AuthRepository

func NewAuthService(repository repository.AuthRepository) AuthService {
	repo = repository
	return &service{}
}
func (*service) Login(email, password string) (string, error) {
	user, err := repo.GetUser(email)
	if err != nil {
		return "", err
	}
	hashedPassword, _ := json.Marshal(user["password"])
	s, _ := strconv.Unquote(string(hashedPassword))

	err = bcrypt.CompareHashAndPassword([]byte(s), []byte(password))
	if err != nil {
		return "", err
	}
	id, err := json.Marshal(user["id"])
	if err != nil {
		return "", err
	}
	return string(id), nil
}

func (*service) Register(email, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	// Generating uuid
	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	document := model.User{Id: uuid, Email: email, Password: string(hashedPassword)}
	err = repo.SaveUser(uuid, document)
	if err != nil {
		return "", err
	}
	return uuid, nil
}
