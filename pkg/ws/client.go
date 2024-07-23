package ws

import (
	"github.com/VAISHAKH-GK/atta-backend/pkg/message"
	"github.com/gofiber/contrib/websocket"
)

var Clients map[string]*Client = map[string]*Client{}

type Client struct {
	Conn        *websocket.Conn
	Connected   bool
	MessageChan chan message.Message
	UnRegister  chan bool
}

func NewClient(conn *websocket.Conn, id string) *Client {
	var c = Client{
		Conn:      conn,
		Connected: false,
	}

	Clients[id] = &c

	return &c
}
