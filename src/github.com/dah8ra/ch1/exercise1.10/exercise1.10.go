package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"bufio"
	"net/http"
	"os"
	"time"
	"strconv"
)

func main() {
	for i := 1 ; i < 3 ; i++ {
		fetchall(i)
	}
}

func fetchall(filenumber int){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		fmt.Println("*** " + url)
		go fetch(url, ch)
	}
	index := strconv.Itoa(filenumber)
	filename := "/workspace/go/golangtraining/src/github.com/dah8ra/ch1/exercise1.10/result" + index + ".txt"
	file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}	
	var writer *bufio.Writer
	writer = bufio.NewWriter(file)
	for range os.Args[1:] {		
		writer.WriteString(<-ch)
		//fmt.Println(<-ch)
		writer.Flush()
	}
	file.Close()
	
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
