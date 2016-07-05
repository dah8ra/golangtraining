package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	comm chan string // an outgoing message channel
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.comm <- msg
			}

		case newclient := <-entering:
			clients[newclient] = true
			for existing, _ := range clients {
				if newclient.name != existing.name {
					newclient.comm <- existing.name + " is in the room."
				}
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.comm)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	who := conn.RemoteAddr().String()
	c := client{comm: ch, name: who}
	go clientWriter(conn, c)

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- c

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- c
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, c client) {
	for msg := range c.comm {
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
