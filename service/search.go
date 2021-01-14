package service

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
	"virtual-bookshelf/model"
)

func Search(query string) []model.Book {
	var books []model.Book
	urlPrefix := "http://www.gutenberg.org"
	c := colly.NewCollector()
	c.OnHTML(".booklink", func(e *colly.HTMLElement) {
		temp := model.Book{}
		imageUrl := e.ChildAttr(".cover-thumb", "src")
		if imageUrl == "" {
			return
		}
		temp.Title = e.ChildText(".title")
		temp.Subtitle = e.ChildText(".subtitle")
		temp.Id = strings.Split(e.ChildAttr("a[href]", "href"), "/ebooks/")[1]
		temp.Url = "/books/" + temp.Id
		temp.Description = get(urlPrefix + "/ebooks/" + temp.Id)
		temp.Image = strings.Replace(urlPrefix+imageUrl, "small", "medium", 1)
		books = append(books, temp)
	})
	url := fmt.Sprintf("http://www.gutenberg.org/ebooks/search/?query=%v&submit_search=Go%21", query)
	err := c.Visit(url)
	if err != nil {
		return nil
	}
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
	err := c.Visit(url)
	if err != nil {
		return ""
	}
	return description

}
