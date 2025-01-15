package main

import (
	"io"
	"net"

	"github.com/rivo/tview"
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

type chattify struct {
	pages *tview.Pages
	term  *tview.Application
	conn  *net.Conn
}

func newChattify() chattify {
	pages := tview.NewPages()

	return chattify{pages, tview.NewApplication().SetRoot(pages, true), nil}
}

func (app *chattify) connectToHost(username string, host string, token string, done chan bool) {
	conn, err := net.Dial("tcp", host)

	if err != nil {
		done <- false
		return
	}

	app.conn = &conn
	done <- true
}

func (app *chattify) loadPages() {
	loginPage := newLoginPage()
	loginPage.build(app)

	chatPage := newChatPage()
	chatPage.build(app.pages)

	app.pages.AddPage(LOGIN_PAGE, loginPage, true, true)
	app.pages.AddPage(CHAT_PAGE, chatPage, true, false)
}

func (app *chattify) Run() error {
	app.loadPages()

	if err := app.term.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	app := newChattify()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
