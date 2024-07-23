package ws

import (
	"sync"

	"github.com/VAISHAKH-GK/atta-backend/pkg/message"
	"github.com/gofiber/contrib/websocket"
)

var Clients map[string]*Client = map[string]*Client{}
var mu sync.Mutex

type Client struct {
	Id             string
	Conn           *websocket.Conn
	LookingConn    bool
	MessageChan    chan message.Message
	UnRegister     chan bool
	ConnectionChan chan string
}

func NewClient(conn *websocket.Conn, id string) *Client {
	var c = Client{
		Id:             id,
		Conn:           conn,
		LookingConn:    true,
		MessageChan:    make(chan message.Message),
		UnRegister:     make(chan bool),
		ConnectionChan: make(chan string),
	}

	Clients[id] = &c

	return &c
}

// Creating new connection between clients
func (c *Client) NewConnection() {
	for id, client := range Clients {
		if id == c.Id {
			continue
		}

		if c.LookingConn {
			mu.Lock()
			c.LookingConn = true
			client.LookingConn = true
			// Sending connected id client id
			var msg = message.Message{
				To:          c.Id,
				MessageType: message.SIGNAL,
				Category:    message.CONNECTED_SIGNAL,
				Content:     id,
			}

			c.MessageChan <- msg
			msg.To = id

			msg.Content = c.Id
			client.MessageChan <- msg

			mu.Unlock()
			return
		}
	}
}
