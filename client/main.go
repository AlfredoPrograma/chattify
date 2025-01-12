package main

import (
	"os"
)

const BUF_SIZE = 1024 // Around 256 characters

type client struct{}

func (c *client) Run() {
	stdin := os.Stdin
	stdout := os.Stdout

	for {
		buf := make([]byte, BUF_SIZE)
		_, err := stdin.Read(buf)

		if err != nil {
			panic(err)
		}

		stdout.WriteString(string(buf))
	}
}

func main() {
	c := client{}

	c.Run()
}
