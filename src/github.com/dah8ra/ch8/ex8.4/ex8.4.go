package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer func() {
		fmt.Println("Called Done!")
		wg.Done()
	}()
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	//	c.Close()
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

var m map[int]net.Conn

func main() {
	m = make(map[int]net.Conn)
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	index := 0
	for {
		wg.Add(1)
		index++
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
		m[index] = conn
	}
	go func() {
		wg.Wait()
		fmt.Println("Finished to wait...")
		for _, c := range m {
			shutdownWrite(c)
		}
	}()
}
