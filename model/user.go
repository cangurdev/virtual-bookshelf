package model

type User struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	Nickname string   `json:"nickname"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Books    []string `json:"books"`
}
