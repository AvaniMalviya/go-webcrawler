package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func crawl( urlToCrawl string, visited map[string]bool, mu *sync.Mutex, wg *sync.WaitGroup){
	defer wg.Done()
	mu.Lock()
	
	if visited[urlToCrawl]{
		mu.Unlock()
		return	
	}
	visited[urlToCrawl] = true
	mu.Unlock()
	
	fmt.Println("Crawling", urlToCrawl)
	resp,err := http.Get(urlToCrawl)
	if err != nil{
		log.Println("Error while fetching the url", err)
		return
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return
	}
	
	var extractLinks func(*html.Node)
	
	extractLinks = func(n *html.Node){
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr{
				if attr.Key == "href"{
					link := attr.Val
					parsedUrl, err := url.Parse(link)
					if err == nil && parsedUrl.IsAbs(){
						mu.Lock()
						if !visited[parsedUrl.String()]{
							wg.Add(1)
							go crawl(parsedUrl.String(), visited, mu, wg)
						}
						mu.Unlock()
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}

	extractLinks(doc)

	time.Sleep(1 * time.Second)
}


func main(){
	startUrl := "https://go.dev/"
	visited := make(map[string]bool)

	var mu sync.Mutex

	var wg sync.WaitGroup

	wg.Add(1)

	go crawl(startUrl,visited, &mu, &wg)

	wg.Wait()
	
	fmt.Println("Crawl complete.")
}
