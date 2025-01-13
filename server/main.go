package main

import (
	"errors"
	"io"
	"log"
	"net"
	"sync"
)

const BUF_SIZE = 1024 // Around 256 characters

const (
	DISCONNECT = iota
	FORCED_DISCONNECT
)

type server struct {
	port        string
	connections map[string]net.Conn
	mu          sync.Mutex
}

func (s *server) registerConn(conn net.Conn) {
	log.Printf("connection with peer %v started\n", conn.RemoteAddr())
	s.connections[conn.RemoteAddr().String()] = conn
}

func (s *server) closeConn(addr string, disconnectKind int) {
	s.mu.Lock()

	conn, ok := s.connections[addr]

	if !ok {
		log.Printf("cannot close connection with unregistered peer %v\n", addr)
		return
	}

	if disconnectKind == DISCONNECT {
		err := conn.Close()

		if err != nil {
			log.Printf("cannot close connection with peer %v\n", addr)
		}
	}

	delete(s.connections, addr)

	log.Printf("connection with peer %v closed\n", addr)
	defer s.mu.Unlock()
}

func (s *server) handleConn(conn net.Conn) {
	for {
		buf := make([]byte, BUF_SIZE)
		bytes, err := conn.Read(buf)

		if err != nil || bytes == 0 {
			if errors.Is(err, io.EOF) || bytes == 0 {
				s.closeConn(conn.RemoteAddr().String(), FORCED_DISCONNECT)
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
		s.registerConn(conn)

		if err != nil {
			panic(err)
		}

		go s.handleConn(conn)
	}
}

func main() {
	s := server{port: ":9999", connections: map[string]net.Conn{}}

	s.Run()
}
