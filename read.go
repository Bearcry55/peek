package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func openArticle(num int) {
	articleURL := results[num-1]
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
		colly.AllowURLRevisit(),
	)

	fmt.Println("\n" + Cyan + "========================================" + Reset)
	fmt.Println(Bold + Cyan + "Article " + fmt.Sprint(num) + Reset)
	fmt.Println(Blue + articleURL + Reset)
	fmt.Println(Cyan + "========================================" + Reset)

	var allContent []string
	found := false

	c.OnHTML("article, [role='main'], main, .post-content, .entry-content, .article-content, #content", func(e *colly.HTMLElement) {
		if !found {
			found = true
			extractContent(e, &allContent)
		}
	})

	if !found {
		c.OnHTML("div", func(e *colly.HTMLElement) {
			text := strings.TrimSpace(e.Text)
			if len(text) > 1000 && strings.Count(text, "\n") > 5 {
				if !found {
					found = true
					extractContent(e, &allContent)
				}
			}
		})
	}

	if err := c.Visit(articleURL); err != nil {
		fmt.Println(Red + "Failed to fetch: " + err.Error() + Reset)
		return
	}

	if !found || len(allContent) == 0 {
		fmt.Println(Yellow + "Could not extract readable content." + Reset)
		fmt.Println("Try visiting the URL directly in your browser.")
		return
	}

	seen := make(map[string]bool)
	for _, text := range allContent {
		cleanText := stripColors(text)
		if !seen[cleanText] && len(cleanText) > 0 {
			fmt.Println(text)
			seen[cleanText] = true
		}
	}
}

func extractContent(e *colly.HTMLElement, content *[]string) {
	e.ForEach("h1, h2, h3, h4, p, pre, code, li", func(_ int, el *colly.HTMLElement) {
		text := strings.TrimSpace(el.Text)
		text = normalizeWhitespace(text)

		if len(text) < 20 {
			return
		}
		if strings.Contains(strings.ToLower(text), "cookie") && len(text) < 100 {
			return
		}
		if strings.Contains(strings.ToLower(text), "subscribe") && len(text) < 100 {
			return
		}

		formatted := formatByTag(el.Name, text)
		*content = append(*content, formatted)
	})
}

func normalizeWhitespace(text string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(text, " ")
}

func formatByTag(tag string, text string) string {
	switch tag {
	case "h1":
		return "\n" + Bold + Cyan + text + Reset + "\n"
	case "h2":
		return "\n" + Bold + Blue + text + Reset
	case "h3":
		return "\n" + Bold + Yellow + text + Reset
	case "h4":
		return Bold + White + text + Reset
	case "pre", "code":
		return Gray + text + Reset
	case "li":
		return "  • " + text
	default:
		return text
	}
}

func stripColors(text string) string {
	re := regexp.MustCompile(`\033\[[0-9;]*m`)
	return re.ReplaceAllString(text, "")
}