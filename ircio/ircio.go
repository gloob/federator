package ircio

import (
	"flag"
	"log"

	"github.com/nickvanw/ircx"
	"github.com/sorcix/irc"
)

var (
	name     = flag.String("name", "federator", "Nick to use in IRC")
	server   = flag.String("server", "irc.irc-hispano.org:6667", "Host:Port to connect to")
	channels = flag.String("chan", "#unix-people", "Channels to join")
)

type IrcConn struct {
	Host string
	Bot  *ircx.Bot
}

func CreateConnection(host string) *IrcConn {
	flag.Parse()
	ic := new(IrcConn)
	ic.Host = host
	ic.Bot = ircx.Classic(*server, *name)
	return ic
}

func (ic *IrcConn) Dial() {
	if err := ic.Bot.Connect(); err != nil {
		log.Panicln("Unable to dial IRC Server ", err)
	}
	RegisterHandlers(ic.Bot)
}

func (ic *IrcConn) Loop() {
	ic.Bot.CallbackLoop()
	log.Println("Exiting..")
}

func RegisterHandlers(bot *ircx.Bot) {
	bot.AddCallback(irc.RPL_WELCOME, ircx.Callback{Handler: ircx.HandlerFunc(RegisterConnect)})
	bot.AddCallback(irc.PING, ircx.Callback{Handler: ircx.HandlerFunc(PingHandler)})
}

func RegisterConnect(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command: irc.JOIN,
		Params:  []string{*channels},
	})
}

func PingHandler(s ircx.Sender, m *irc.Message) {
	s.Send(&irc.Message{
		Command:  irc.PONG,
		Params:   m.Params,
		Trailing: m.Trailing,
	})
}
