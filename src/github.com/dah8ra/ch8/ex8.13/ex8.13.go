package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	message chan string // an outgoing message channel
	name    string
	updated chan struct{}
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[string]client) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for _, cli := range clients {
				cli.message <- msg
			}
		case newclient := <-entering:
			clients[newclient.name] = newclient
			for _, existing := range clients {
				if newclient.name != existing.name {
					newclient.message <- existing.name + " is in the room."
				}
			}
		case cli := <-leaving:
			delete(clients, cli.name)
			close(cli.message)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	who := conn.RemoteAddr().String()
	c := client{message: ch, name: who}
	go clientWriter(conn, c)

	ch <- "You are " + who
	c.message <- who + " has arrived"
	entering <- c

	input := bufio.NewScanner(conn)
	c.updated = make(chan struct{})
	go sessionTimer(conn, c)

	for input.Scan() {
		messages <- who + ": " + input.Text()
		c.updated <- struct{}{}
	}

}

func sessionTimer(conn net.Conn, c client) {
	flag := false
	for {
		select {
		// Decide timeout for a client.
		case <-time.After(10 * time.Second):
			fmt.Printf("close conn -> %s\n", c.name)
			leaving <- c
			messages <- c.name + " has left"
			conn.Close()
			flag = true
			break
		case <-c.updated:
			fmt.Println("Updated")
			break
		}
		if flag {
			break
		}
	}
}

func clientWriter(conn net.Conn, c client) {
	for msg := range c.message {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
