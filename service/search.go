package service

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
	"virtual-bookshelf/model"
)

func Search(query string) []model.Book {
	var books []model.Book
	c := colly.NewCollector()
	c.OnHTML(".booklink", func(e *colly.HTMLElement) {
		temp := model.Book{}
		temp.Title = e.ChildText(".title")
		temp.Subtitle = e.ChildText(".subtitle")
		temp.Url = "http://www.gutenberg.org/" + e.ChildAttr("a[href]", "href")
		temp.Description = get(temp.Url)
		id := strings.Split(temp.Url, "/")[5]
		temp.Image = fmt.Sprintf("http://www.gutenberg.org/cache/epub/%s/pg%s.cover.medium.jpg", id, id)
		books = append(books, temp)
	})
	url := fmt.Sprintf("http://www.gutenberg.org/ebooks/search/?query=%v&submit_search=Go%21", query)
	c.Visit(url)

	return books
}

func get(url string) string {
	c := colly.NewCollector()
	description := ""
	c.OnHTML(".bibrec", func(e *colly.HTMLElement) {
		description += e.ChildText("tr[itemprop=inLanguage]")
		description += "\n"
		description += e.ChildText("tr[datatype=\"xsd:date\"]")
	})
	c.Visit(url)
	return description

}
