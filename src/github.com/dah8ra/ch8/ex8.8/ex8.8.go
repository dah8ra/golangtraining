package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	fmt.Println("Called handleConn")
	tick := time.Tick(10 * time.Second)
	input := bufio.NewScanner(c)
	token := make(chan struct{})
	for input.Scan() {
		token <- struct{}{}
		go echo(c, input.Text(), 1*time.Second)
	}

	select {
	case <-tick:
		c.Close()
		fmt.Println("Close connection...")
	case <-token:
		tick = time.Tick(10 * time.Second)
		fmt.Println("Keep connection...")
	}

}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("@@@")
		conn, err := l.Accept()
		fmt.Println("@@@")
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		fmt.Println("@@@")
		fmt.Println("Call handleConn")
		go handleConn(conn)
	}
}
