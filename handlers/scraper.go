package handlers

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func WebScraper() (string, error) {
	c := colly.NewCollector()
	var text string
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting:", request.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Status:", response.StatusCode)
	})

	c.OnHTML("div.fc-text-section_wrapperClassName__Sgd9D", func(element *colly.HTMLElement) {
		fmt.Println(element.Text)
		text = element.Text
	})

	err := c.Visit("https://www.fifa.com/tournaments/mens/worldcup/canadamexicousa2026/tickets")
	if err != nil {
		log.Fatal("Failed to visit url")
	}

	return text, nil
}
