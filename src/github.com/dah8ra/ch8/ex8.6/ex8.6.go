package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopl.io/ch5/links"
)

type URL struct {
	url   string
	depth int
}

func crawl(url URL) map[int]URL {
	fmt.Println(url.url)
	list, err := links.Extract(url.url)
	if err != nil {
		log.Print(err)
	}
	m := make(map[int]URL)
	for index, link := range list {
		m[index] = URL{url: link, depth: url.depth}
	}

	return m
}

func main() {
	input := os.Args[1:2]
	fmt.Printf("DEPTH ------> %v\n", input)
	max, _ := strconv.Atoi(input[0])
	worklist := make(chan URL)    // lists of URLs, may have duplicates
	unseenLinks := make(chan URL) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		for _, url := range os.Args[2:] {
			worklist <- URL{url: url, depth: 0}
		}
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					for _, url := range foundLinks {
						worklist <- URL{url: url.url, depth: url.depth + 1}
					}
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for lists := range worklist {
		link := lists.url
		depth := lists.depth
		if depth > max-1 {
			fmt.Printf("Reached max depth: %d\n", depth)
		} else if !seen[link] {
			seen[link] = true
			unseenLinks <- lists
		}
	}
}
