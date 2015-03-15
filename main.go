package main

import (
  "log"
  "os"
  "time"

  "github.com/gloob/federator/irc"
)

func main() {
  log.SetOutput(os.Stdout)
  log.SetFlags(log.Ltime)

  ircConn := irc.createConnection("irc.irc-hispano.org")

  ircConn.dial()

  for {
    time.Sleep(1 * time.Second)
  }
}
