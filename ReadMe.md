# Go Web Crawler

A simple, concurrent web crawler built in Go. It fetches pages, extracts links, and recursively crawls through them. This crawler uses goroutines for efficient crawling and ensures each URL is visited only once.

## Features

- Concurrent crawling using Go's goroutines.
- Extracts links (`<a href="...">`) from HTML.
- Avoids revisiting URLs.
- 1-second delay between requests to avoid overloading the server.

## Prerequisites

- Go 1.18+.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/avanimalviya/go-webcrawler.git
   cd go-webcrawler
   ```

2. Initialize Go modules (if needed):

   ```bash
   go mod init webcrawler
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

1.  Update the starting URL in main.go:

    ```go
    startURL := "http://example.com"
    ```

2.  Run the crawler:

    ```bash
    go run main.go
    ```

    The crawler will start at the given URL and recursively crawl through all discovered links.
