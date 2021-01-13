package test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"testing"
	"virtual-bookshelf/service"
)

func TestHomeTest(t *testing.T) {
	app := fiber.New()

	// Create route with GET method for test:
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("index page")
	})

	// http.Request
	req := httptest.NewRequest("GET", "http://127.0.0.1:3000", nil)
	req.Header.Set("X-Custom-Header", "hi")

	// http.Response
	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, string(body), "index page")
}
func TestGetBookTest(t *testing.T) {
	q := "sherlock"
	books := service.Search(q)
	fmt.Println(len(books) != 0)
}
func TestLogin(t *testing.T) {
	email := "can@gmail.com"
	password := "123456"
	req := httptest.NewRequest("POST", "/login", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	form, _ := url.ParseQuery(req.URL.RawQuery)
	form.Add("email", email)
	form.Add("password", password)
	//req.Form.Add("password", password)
	//req.Form.Add("email",email)
	//req.Form.Encode()
	req.URL.RawQuery = form.Encode()
	assert.Equal(t, nil, req.Response)
}
func TestAddBookTest(t *testing.T) {
	app := fiber.New()
	app.Get("/add-book/:bookId", func(c *fiber.Ctx) error {

		return c.SendString("book added")
	})
	req := httptest.NewRequest("GET", "http://127.0.0.1:3000/add-book/1", nil)
	// http.Response
	resp, _ := app.Test(req)

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, string(body), "book added")
}
