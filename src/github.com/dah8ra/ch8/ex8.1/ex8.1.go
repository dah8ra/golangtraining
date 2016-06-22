package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
Please use following command.
$ go run ex8.1.go 8010 8020 8030
*/

var timeZone = map[string]int{
	"US/Eastern":    -5 * 60 * 60,
	"Asia/Tokyo":    9 * 60 * 60,
	"Europe/London": 1 * 60 * 60,
}

var portMapping = map[string]string{
	"8010": "US/Eastern",
	"8020": "Asia/Tokyo",
	"8030": "Europe/London",
}

func handleConn(c net.Conn, location string) {
	defer c.Close()
//	io.WriteString(c, location)
//	io.WriteString(c, "\n")
	for {
		now := time.Now()
		utc := now.UTC()
		locale := time.FixedZone(location, timeZone[location])
		localeTime := utc.In(locale)
		_, err := io.WriteString(c, localeTime.Format("15:04:05"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func createListener(port string) {
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	location := portMapping[port]

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn, location)
	}
}

func main() {
	for _, port := range os.Args[1:] {
		go createListener(port)
	}
	time.Sleep(time.Minute)
}
