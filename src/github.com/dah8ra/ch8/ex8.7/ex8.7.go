package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type URL struct {
	url   string
	depth int
	text  []byte
}

var index int = 0

func crawl(url URL) map[int]URL {
	fmt.Println(url.url)
	list, err := extract(url.url)
	if err != nil {
		log.Print(err)
	}
	text, err1 := getText(url.url)
	if err1 != nil {
		log.Print(err1)
	}

	m := make(map[int]URL)
	for index, link := range list {
		m[index] = URL{url: link, depth: url.depth, text: text}
	}

	return m
}

// First Args: depth, Second: domain, Third: url
func main() {
	input := os.Args[1:2]
	domains := os.Args[2:3]
	fmt.Printf("DEPTH ------> %v\nDOMAIN ------> %v\n", input, domains)
	domain := domains[0]
	r := regexp.MustCompile(domain)
	max, _ := strconv.Atoi(input[0])
	worklist := make(chan URL)    // lists of URLs, may have duplicates
	unseenLinks := make(chan URL) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		initbyte := []byte("")
		for _, url := range os.Args[3:] {
			worklist <- URL{url: url, depth: 0, text: initbyte}
		}
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					for _, url := range foundLinks {
						path := createDirectory("./")
						saveData(path, url)
						worklist <- URL{url: url.url, depth: url.depth + 1, text: url.text}
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
		if r.MatchString(link) {
			depth := lists.depth
			if depth > max-1 {
				fmt.Printf("Reached max depth: %d\n", depth)
			} else if !seen[link] {
				seen[link] = true
				unseenLinks <- lists
			}
		} else {
			fmt.Printf("Not match the domain: %s\n", link)
		}
	}
}

func getText(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	text, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return nil, err1
	}
	return text, nil
}

////////////////////////////
// Copy links.go from ch5 //
////////////////////////////
func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err1 := html.Parse(resp.Body)
	if err1 != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err1)
	}
	resp.Body.Close()
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func saveData(path string, url URL) {
	if strings.Contains("javascript", path) {
		return
	}
	if !isExist(path) {
		return
	}

	//	fmt.Printf("@@@@@@@@@@@@@F %s\n", fname)
	//	dir := headercut(fname)
	//	fmt.Printf("@@@@@@@@@@@@@D %s\n", dir)
	//	fmt.Printf("@@@@@@@@@@@@@P %s\n", path)
	//	t := path + "/" + dir + ".html"
	//	fmt.Printf("@@@@@@@@@@@@@T %s\n", t)
	//	fmt.Println()
	//	fmt.Printf("@@@@@@@@@@@@@ %s : %s : %s\n", t, fname, path)
	//	t := path + "/" + strconv.Itoa(index) + ".html"
	t := path + "/" + url.url + ".html"
	ioutil.WriteFile(t, url.text, os.ModePerm)
	index++
}

func separateSlash(dir string) string {
	sep := strings.Split(dir, "/")
	return sep[0]
}

func headercut(dirname string) string {
	sep1 := strings.Split(dirname, "http://")
	var sep2 []string
	if len(sep1) == 1 {
		sep2 = strings.Split(dirname, "https://")
		if len(sep2) == 1 {
			return ""
		}
	}
	if sep2 == nil {
		sep2 = sep1
	}
	withHtml := sep2[1]
	sep3 := strings.Split(withHtml, ".html")
	return sep3[0]
}

func createDirectory(dirname string) string {
	dir := headercut(dirname)
	sep := separateSlash(dir)
	p, _ := os.Getwd()
	dname := p + "/" + sep
	os.Mkdir(dname, 0777)
	//	if err := os.Mkdir(dname, 0777); err != nil {
	//		fmt.Printf("[ERROR] %s\n", err)
	//	}
	//	fmt.Printf("@@@@@@@@@@@@@Directory %s\n", dname)
	return dname
}
