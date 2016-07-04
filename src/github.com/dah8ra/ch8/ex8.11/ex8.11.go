package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

// go run ex8.11.go http://www.yahoo.co.jp/ http://www.google.co.jp ...
func main() {
	fmt.Println(mirroredQuery())
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	urls := os.Args[1:]
	for _, url := range urls {
		fmt.Println(url)
		done := make(chan struct{})
		go func(url string, done chan struct{}) { responses <- request(url, done) }(url, done)

		select {
		case <-responses:
			fmt.Println("Closed...")
			close(done)
			break
		default:
		}
	}

	return <-responses
}

func request(hostname string, done chan struct{}) (response string) {

	req, _ := http.NewRequest("GET", hostname, nil)
	req.Cancel = done

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(resp.Body)
	response = bufbody.String()
	return response
}
