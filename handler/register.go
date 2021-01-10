package handler

import (
	"crypto/rand"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"virtual-bookshelf/database"
	"virtual-bookshelf/model"
)

func GetRegister(c *fiber.Ctx) error {
	return c.Render("sign_up", fiber.Map{
		"Title": "register",
	})
}
func PostRegister(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	document := model.User{Email: email, Password: string(hashedPassword)}
	// Generating uuid
	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	_, err = database.GetCollection().Insert(uuid, &document, nil)
	if err != nil {
		panic(err)
	}
	return c.Redirect("/home")
}
