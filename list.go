package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

var results []string

func search(query string) {
	searchURL := "https://duckduckgo.com/html/?q=" + strings.ReplaceAll(query, " ", "+")
	c := colly.NewCollector(
		colly.UserAgent("PEEK/0.1"),
		colly.AllowURLRevisit(),
	)

	fmt.Println("Searching:", query)

	c.OnHTML(".result__a", func(e *colly.HTMLElement) {
		href := e.Attr("href")

		// Extract actual URL from DuckDuckGo redirect
		if strings.Contains(href, "uddg=") {
			parts := strings.Split(href, "uddg=")
			if len(parts) > 1 {
				encoded := strings.Split(parts[1], "&")[0]
				decoded, err := url.QueryUnescape(encoded)
				if err == nil {
					href = decoded
				}
			}
		}

		results = append(results, href)
		fmt.Printf("%d. %s\n", len(results), strings.TrimSpace(e.Text))
	})

	if err := c.Visit(searchURL); err != nil {
		fmt.Println("Search failed:", err)
	}
}