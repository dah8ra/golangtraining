package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/dah8ra/ch9/ex93"
)

func httpGetBody(url string, done chan struct{}) (interface{}, error) {
	os.Setenv("HTTP_PROXY", "http://usrname:password@proxyurl:8080")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done chan struct{}) (interface{}, error)
}

/////////////////////////////
// Added main function
/////////////////////////////
func main() {
	m := ex93.New(httpGetBody)
	//	Sequential(m)
	Concurrent(m)
}

func Sequential(m M) {
	for url := range incomingURLs() {
		start := time.Now()
		done := make(chan struct{})
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		////////////////////////
		// Detect cancel action
		////////////////////////
		if value == nil && err == nil {
			fmt.Println("Cancel action...")
			break
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			done := make(chan struct{})
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
				return
			}
			////////////////////////
			// Detect cancel action
			////////////////////////
			if value == nil && err == nil {
				fmt.Println("Cancel action...")
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))

		}(url)
	}
	n.Wait()
}
