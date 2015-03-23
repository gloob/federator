package main

import (
	"log"
	"os"
	"time"

	"github.com/gloob/federator/ircio"
	"github.com/gloob/go-telegram/tgl"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime)

	host := "irc.irc-hispano.org"
	ircConn := ircio.CreateConnection(host)

	ircConn.Dial()

	tglState := tgl.NewState()
	defer state.Destroy()

	for {
		time.Sleep(1 * time.Second)
	}
}
