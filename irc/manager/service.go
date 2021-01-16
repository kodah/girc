package manager

import (
	"net"
	"sync"
	"time"

	"github.com/guark/guark/log"
	"gopkg.in/irc.v3"

	"github.com/kodah/girc/store"
)

var Service *ConnectionManager

func New() *ConnectionManager {
	Service = new(ConnectionManager)

	Service.Log = log.New("connection-pooler")

	return Service
}

type ConnectionManager struct {
	Log     log.Logger
	Clients sync.Map
}

func (c *ConnectionManager) NewClient(nick, pass, user, name, address string, frequency, timeout, limit time.Duration, burst int) (*irc.Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	config := irc.ClientConfig{
		Nick:          nick,
		Pass:          pass,
		User:          user,
		Name:          name,
		PingFrequency: frequency,
		PingTimeout:   timeout,
		SendLimit:     limit,
		SendBurst:     burst,
		Handler: irc.HandlerFunc(func(client *irc.Client, message *irc.Message) {
			switch message.Command {
			// welcome event
			case "001":
				client.Write("JOIN #girc-test")
			default:
				c.Log.Debug("Unknown command ", "command: ", message.Command, " user: ", message.User, " name: ", message.Name, " host: ", message.Host, " params: ", message.Params, " prefix: ", message.Prefix, " tags: ", message.Tags)

				_ = store.Service.AddMessage(message.Name, message.User, message.Host, message.Command)
			}
		}),
	}

	client := irc.NewClient(conn, config)

	// store the client connection under the nick@address (eg: kodah@chat.freenode.net)
	c.Clients.Store(nick+"@"+address, client)

	return client, client.Run()
}
