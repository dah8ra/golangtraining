package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/dah8ra/ch4/omdbapi"
)

const baseurl = "http://omdbapi.com/?"
const basefilename = "/workspace/go/golangtraining/src/github.com/dah8ra/ch4/ex4.13/"

var word = flag.String("w", "frozen", "Search word")

func main() {
	flag.Parse()
	values := url.Values{}
	values.Add("t", *word)
	values.Add("r", "json")
	url := baseurl + values.Encode()
	fmt.Printf("-------> %s\n", url)
	r, _ := omdbapi.Get(url)
	if r != nil {
		fmt.Printf("#%s %s %s\n", r.Title, r.Year, r.Poster)
	} else {
		fmt.Println("Failed to download...")
	}

	if r != nil {
		body, _ := omdbapi.Download(r.Poster)
		filename := basefilename + *word + ".jpeg"
		fmt.Printf("Saved ==> %s\n", filename)
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)

		if err != nil {
			fmt.Println(err)
		}

		defer file.Close()

		file.Write(body)
	} else {
		fmt.Println("Failed to download...")
	}
}
