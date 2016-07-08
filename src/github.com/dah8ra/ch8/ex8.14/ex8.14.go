package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
	"strconv"
)

type client struct {
	writer http.ResponseWriter // an outgoing message channel
	name   string
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
				cli.writer.Write([]byte(msg + "\r\n"))
			}
		case newclient := <-entering:
			clients[newclient] = true
			for existing, _ := range clients {
				if newclient.name != existing.name {
					newclient.writer.Write([]byte(existing.name + " is in the room.\r\n"))
				}
			}
		case <-leaving:
			//			delete(clients, cli)
			//			close(cli.writer)
		}
	}
}

func handleConn(w http.ResponseWriter, r *http.Request) {
	//	ch := make(chan ) // outgoing client messages
	rand.Seed(time.Now().UnixNano())
	who := strconv.Itoa(rand.Intn(100)) // Need to check same number
	c := client{writer: w, name: who}
	go clientWriter(w, r)

	w.Write([]byte("You are " + who + "\r\n"))
	messages <- who + " has arrived"
	entering <- c

	//	leaving <- c
	//	messages <- who + " has left"
}

func clientWriter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.PostFormValue("key")))
}

func communication(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	go broadcaster()
	http.HandleFunc("/", handleConn)
	http.HandleFunc("/comm", communication)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
