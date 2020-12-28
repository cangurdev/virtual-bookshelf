package test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHomeTest(t *testing.T) {
	app := fiber.New()
	// Create route with GET method for test:

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.BaseURL())              // => http://google.com
		fmt.Println(c.Get("X-Custom-Header")) // => hi

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
