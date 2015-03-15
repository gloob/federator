package telegramio

import (
  "github.com/gloob/go-telegram/tgl"
)

type Conn struct {
  Host string
}

func createConnection(host string) *Conn {
  c := new(Conn)
  c.Host = host

  tgl_state := new(C.struct_tgl_state)
  tgl_state = tgl.login(tgl_state)

  return c
}
