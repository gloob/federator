package telegramio

import (
	"github.com/gloob/go-telegram/tgl"
)

type Conn struct {
	Host string
}

func CreateConnection() *Conn {
	c := new(Conn)
	//c.Host = host

	state := tgl.NewState()
	defer state.Destroy()

	//err := state.Dial()

	return c
}
