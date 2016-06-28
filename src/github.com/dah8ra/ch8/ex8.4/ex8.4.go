package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

//var m map[int]net.Conn

func main() {
	//	m = make(map[int]net.Conn)
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func closeStdin() {
	time.Sleep(5 * time.Second)
	err := os.Stdin.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second)
	}
	wg.Wait()
	fmt.Println("Close conn.")
	shutdownWrite(c)
	//	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	defer func() {
		fmt.Println("Called Done!")
		wg.Done()
	}()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func shutdownWrite(conn net.Conn) {
	switch i := conn.(type) {
	case *net.TCPConn:
		fmt.Println("CLOSE -> *net.TCPConn")
		i.CloseWrite()
	case *net.UnixConn:
		fmt.Println("CLOSE -> *net.UnixConn")
		i.CloseWrite()
	}
}
