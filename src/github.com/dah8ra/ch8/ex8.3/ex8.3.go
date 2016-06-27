package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	go waitDone(done)
	go closeStdin()
	mustCopy(conn, os.Stdin)
	shutdownWrite(conn)
	time.Sleep(20 * time.Second)
}

func closeStdin() {
	time.Sleep(5 * time.Second)
	err := os.Stdin.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func waitDone(done chan struct{}) {
	<-done
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

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
