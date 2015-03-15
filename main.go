package main

import (
  "log"
  "os"
  "time"

  "github.com/gloob/federator/ircio"
  "github.com/gloob/federator/telegramio"
)

func main() {
  log.SetOutput(os.Stdout)
  log.SetFlags(log.Ltime)

  ircConn := ircio.createConnection("irc.irc-hispano.org")

  ircConn.dial()

  for {
    time.Sleep(1 * time.Second)
  }
}
