package main

import (
	"log"
	"os"
	"time"

	"github.com/gloob/federator/ircio"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime)

	host := "irc.irc-hispano.org"
	ircConn := ircio.CreateConnection(host)

	ircConn.Dial()

	ircConn.Loop()
}
