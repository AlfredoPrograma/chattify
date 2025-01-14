package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	LOGIN_PAGE = "LOGIN_PAGE"
	CHAT_PAGE  = "CHAT_PAGE"
)

type loginPage struct {
	username string
	host     string
	token    string
	tview.Primitive
}

func newLoginPage() loginPage {
	return loginPage{
		username:  "",
		host:      "",
		token:     "",
		Primitive: tview.NewFlex(),
	}
}

func (p *loginPage) build(pages *tview.Pages) {
	root, ok := p.Primitive.(*tview.Flex)
	form := tview.NewForm()

	if !ok {
		panic("invalid primitive for login page")
	}

	form.AddInputField("Username", p.username, 0, nil, func(username string) {
		p.username = username
	})
	form.AddInputField("Host", p.host, 0, nil, func(host string) {
		p.host = host
	})
	form.AddInputField("Token", p.token, 0, nil, func(token string) {
		p.token = token
	})

	form.AddButton("Sign in", func() {
		pages.SwitchToPage(CHAT_PAGE)
	})

	form.Box.SetBackgroundColor(0).SetBorder(true)
	root.AddItem(form, 0, 1, true)
}

type chatPage struct {
	messages []string
	input    string
	tview.Primitive
}

func newChatPage() chatPage {
	return chatPage{
		messages:  []string{},
		input:     "",
		Primitive: tview.NewFlex(),
	}
}

func (p *chatPage) build(pages *tview.Pages) {
	root, ok := p.Primitive.(*tview.Flex)

	if !ok {
		panic("invalid primitive for chat page")
	}

	root.SetDirection(tview.FlexRow)
	messagesContainer := tview.NewFlex()
	messagesContainer.Box.SetBackgroundColor(tcell.Color155)

	input := tview.NewInputField()
	input.SetText(p.input)
	input.SetChangedFunc(func(text string) {
		p.input = text
	})

	root.AddItem(messagesContainer, 0, 90, false)
	root.AddItem(input, 0, 10, true)
}
