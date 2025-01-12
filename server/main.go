package main

import (
	"errors"
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
		bytes, err := conn.Read(buf)

		if err != nil || bytes == 0 {
			if errors.Is(err, io.EOF) || bytes == 0 {
				log.Printf("connection with peer %v closed\n", conn.RemoteAddr())
				break
			}
		}

		_, err = conn.Write(buf)

		if err != nil {
			log.Printf("cannot write over peer %v\n", conn.RemoteAddr())
			panic(err)
		}
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

		go s.handleConn(conn)
	}
}

func main() {
	s := server{port: ":9999"}

	s.Run()
}
