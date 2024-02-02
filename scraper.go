package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// Regular expression to match stock tickers (example: AAPL, GOOGL, TSLA)
	tickerRegex := regexp.MustCompile(`\b[A-Z]{2,5}\b`)

	// Set up a callback to be called when a visited HTML element is found
	c.OnHTML("body", func(e *colly.HTMLElement) {
		text := e.Text
		matches := tickerRegex.FindAllString(text, -1)

		for _, match := range matches {
			fmt.Println("Found Ticker:", match)
		}
	})

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping a sample news website
	err := c.Visit("https://example.com/news")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
