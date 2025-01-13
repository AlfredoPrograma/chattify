package main

import (
	"io"
	"net"
	"os"
)

const BUF_SIZE = 1024 // Around 256 characters

type client struct {
	stdin  io.Reader
	stdout io.Writer
}

func (c *client) readFromServer(conn net.Conn) {
	for {
		buf := make([]byte, BUF_SIZE)
		_, err := conn.Read(buf)

		if err != nil {
			panic(err)
		}

		c.stdout.Write(buf)
	}
}

func (c *client) writeToServer(conn net.Conn) {
	for {
		buf := make([]byte, BUF_SIZE)
		_, err := c.stdin.Read(buf)

		if err != nil {
			panic(err)
		}

		conn.Write(buf)
	}
}

func (c *client) Run() {
	conn, err := net.Dial("tcp", ":9999")

	if err != nil {
		panic(err)
	}

	go c.writeToServer(conn)
	go c.readFromServer(conn)

	select {}
}

func main() {
	c := client{os.Stdin, os.Stdout}

	c.Run()
}
