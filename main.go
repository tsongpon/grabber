package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	found := 0

	c.OnHTML(".post .inner", func(e *colly.HTMLElement) {
		if found == 0 {
			content := e.Text
			fmt.Println(content)
			found++
		}
	})

	c.OnHTML("title", func(e *colly.HTMLElement) { // Title
		title := e.Text
		fmt.Println(title)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://xonly8.com/index.php?topic=253295.0")
}
