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

  ircConn.Dial()
  _ = plugin.NewDispatcher(ircConn, config, db)

  for {
    time.Sleep(1 * time.Second)
  }
}
