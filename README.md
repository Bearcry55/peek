# Peek

A lightweight terminal-based web search and article reader for Linux/Unix systems.

Search the web and read articles without leaving your terminal.

## Features

- 🔍 Fast web search via DuckDuckGo
- 📖 Clean article extraction and reading
- 🚀 No browser required
- 💾 Zero configuration
- 🎯 Interactive result selection
- 🧹 Removes ads, navigation, and clutter

## Installation

### Prerequisites

- Go 1.16 or higher

### Install
```bash
git clone https://github.com/bearcry55/peek.git
cd peek
go build
```

### Optional: Add to PATH
```bash
# Linux/macOS
sudo mv peek /usr/local/bin/

# Or add to your shell profile
export PATH="$PATH:/path/to/peek"
```

## Usage

### Basic Search
```bash
peek <your search query>
```

### Example
```bash
peek what is linux
```

This will:
1. Display numbered search results
2. Prompt you to select an article number
3. Extract and display the article content
4. Allow you to read another result or quit

### Interactive Commands

- Enter a **number** to read that article
- Enter **q** to quit

## How It Works
```
User Query → DuckDuckGo Search → Numbered Results
     ↓
Select Article Number
     ↓
Fetch & Parse HTML → Extract Main Content → Display in Terminal
```

## Example Session
```bash
$ peek golang concurrency

Searching: golang concurrency

1. Go Wiki: LearnConcurrency - The Go Programming Language
2. Go Concurrency Patterns
3. Understanding Concurrency in Go
...

Enter article number (or 'q' to quit): 1

========================================
Article 1
https://go.dev/wiki/LearnConcurrency
========================================

[Article content displays here]

Enter article number (or 'q' to quit): q
```

## Limitations

- Cannot read JavaScript-heavy sites
- Some sites may block scraping
- Paywalled content not accessible
- Success rate: ~75-80% of sites

## Dependencies

- [colly](https://github.com/gocolly/colly) - Web scraping framework

## Contributing

Contributions are welcome! Feel free to:

- Report bugs
- Suggest features
- Submit pull requests

## License

MIT License - see [LICENSE](LICENSE) file for details

## Author

Built by Deep Narayan Banerjee with Go and curiosity.

## Acknowledgments

- DuckDuckGo for search results
- The Go community for excellent tooling
```

