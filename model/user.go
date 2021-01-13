package model

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Books    []Book `json:"books"`
}
