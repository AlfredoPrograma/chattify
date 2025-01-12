package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const BUF_SIZE = 1024 // Around 256 characters

type server struct {
	port string
}

func (s *server) handleConn(conn net.Conn) {
	for {
		buf := make([]byte, BUF_SIZE)
		_, err := conn.Read(buf)

		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("connection with peer %v closed\n", conn.RemoteAddr())
				break
			}
		}

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
