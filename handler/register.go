package handler

import "github.com/gofiber/fiber/v2"
import "virtual-bookshelf/database"
import "virtual-bookshelf/model"

func GetRegister(c *fiber.Ctx) error {
	return c.Render("sign_up", fiber.Map{
		"Title": "register",
	})
}
func PostRegister(c *fiber.Ctx) error {
	document := model.User{Email: c.FormValue("email"), Password: c.FormValue("password")}
	_, err := database.GetCollection().Insert("document-key", &document, nil)
	if err != nil {
		panic(err)
	}
	return c.Redirect("/")
}
