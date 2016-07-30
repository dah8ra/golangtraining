package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/dah8ra/ch8/links"
)

const auth = "username:password"

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	//	depth := flag.Int("depth", 1, "extract depth")
	//	url := flag.String("url", "", "root url")
	//	fmt.Printf("depth: %d\n", *depth)

	os.Setenv("HTTP_PROXY", "proxy.ricoh.co.jp:8080")
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	//	go func() {
	//		url := []string{*url}
	//		urls := make([]string, url)
	//		worklist <- urls
	//	}()
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
