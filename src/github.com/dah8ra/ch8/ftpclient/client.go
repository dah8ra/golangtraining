package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

func main() {
	c := NewClient(1)
	c.Connect()
}

type Client struct {
	reader     *textproto.Reader
	writer     *textproto.Writer
	conn       net.Conn
	passive    net.Conn
	lastMsg    string
	passReader *bufio.Reader
	passWriter *bufio.Writer
	id         int
}

func NewClient(id int) *Client {
	c := Client{}
	c.id = id
	return &c
}

func (c *Client) read(print bool) {
	if c.reader == nil {
		return
	}
	code, msg, err := c.reader.ReadResponse(0)
	if print {
		fmt.Println(code, msg)
	}
	c.lastMsg = msg
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) send(text string) {
	if c.writer == nil {
		return
	}
	err := c.writer.PrintfLine(text)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) Connect() {
	var err error
	c.conn, err = net.DialTimeout("tcp", "127.0.0.1:2121", 10000000)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.reader = textproto.NewReader(bufio.NewReader(c.conn))
	c.writer = textproto.NewWriter(bufio.NewWriter(c.conn))

	c.Pwd()

	c.read(false)
	c.send("USER bad")
	c.read(false)
	c.send("PASS security")
	c.read(false)
}

func (c *Client) Pwd() {
	c.send("PWD")
	c.read(true)
}

func (c *Client) Quit() {
	c.send("QUIT")
	c.read(true)
}
