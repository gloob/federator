package telegramio

import (
	_ "github.com/gloob/go-telegram"
)

type Conn struct {
	Host string
}

func createConnection(host string) *Conn {
	c := new(Conn)
	c.Host = host

	state := NewState()
	defer state.Destroy()

	//err := state.Dial()

	return c
}
