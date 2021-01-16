package manager

import (
	"net"
	"time"

	"github.com/guark/guark/log"
	"gopkg.in/irc.v3"

	"github.com/kodah/girc/store"
)

type User struct {
	Nick         string
	Host         string
	UserModes    string
	ChannelModes map[string]string
}

type Channel struct {
	Modes []string
	Users []*User
	Topic string
}

type Connection struct {
	Log      log.Log
	Client   *irc.Client
	Channels []Channel
	Users    []*User
}

func New(nick, pass, user, name, address string, frequency, timeout, limit time.Duration, burst int) (*Connection, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	connection := new(Connection)
	connection.Log = log.New(nick + "@" + address)

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
				connection.Log.Debug("Unknown command ", "command: ", message.Command, " user: ", message.User, " name: ", message.Name, " host: ", message.Host, " params: ", message.Params, " prefix: ", message.Prefix, " tags: ", message.Tags)

				_ = store.Service.AddMessage(message.Name, message.User, message.Host, message.Command)
			}
		}),
	}

	connection.Client = irc.NewClient(conn, config)

	return connection, connection.Client.Run()
}
