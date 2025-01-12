package main

import (
	"fmt"
	"net"
)

const BUF_SIZE = 1024 // Around 256 characters

type server struct {
	port string
}

func (s *server) handleConn(conn net.Conn) {
	for {
		buf := make([]byte, BUF_SIZE)
		conn.Read(buf)
		fmt.Println(string(buf))
	}
}

func (s *server) Run() {
	ln, err := net.Listen("tcp", s.port)

	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		s.handleConn(conn)
	}
}

func main() {
	s := server{port: ":9999"}

	s.Run()
}
