package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: peek <search query>")
		return
	}

	query := strings.Join(os.Args[1:], " ")
	search(query)

	// Interactive loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nEnter article number (or 'q' to quit): ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "q" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil || num < 1 || num > len(results) {
			fmt.Println("Invalid number. Try again.")
			continue
		}

		openArticle(num)
	}
}