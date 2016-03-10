package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/dah8ra/ch4/xkcdcom"
)

var word = flag.String("w", "default", "Search word")

const preurl = "https://xkcd.com/"
const sufurl = "/info.0.json"

var x xkcdcom.Xkcd

func main() {
	m := make(map[string]string)
	for i := 570; i < 572; i++ {
		url := preurl + strconv.Itoa(i) + sufurl
		fmt.Printf("-------> %s\n", url)
		r, _ := xkcdcom.Get(url)
		m[url] = r.Transcript
		fmt.Printf("#%-5d %s %.10s\n", r.Num, r.Title, r.Transcript)
	}
	fmt.Println("@@@@@@@@@@@@")
	flag.Parse()
	//	search := []byte(*word)
	for url, trans := range m {
		if strings.Contains(trans, *word) {
			fmt.Printf("%s\n%s\n@@@\n", url, trans)
		}
	}
}
