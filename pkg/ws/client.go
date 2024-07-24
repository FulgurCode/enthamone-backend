package ws

import (
	"sync"

	"github.com/VAISHAKH-GK/atta-backend/pkg/message"
	"github.com/gofiber/contrib/websocket"
)

var Clients map[string]*Client = map[string]*Client{}

type Client struct {
	Id             string
	Conn           *websocket.Conn
	LookingConn    bool
	ConnectedUser  string
	MessageChan    chan message.Message
	UnRegister     chan bool
	ConnectChan    chan string
	DisconnectChan chan bool
	mu             sync.Mutex
}

func NewClient(conn *websocket.Conn, id string) *Client {
	var c = Client{
		Id:             id,
		Conn:           conn,
		LookingConn:    true,
		MessageChan:    make(chan message.Message),
		UnRegister:     make(chan bool),
		ConnectChan:    make(chan string),
		DisconnectChan: make(chan bool),
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

		if client.LookingConn {
			c.mu.Lock()
			client.mu.Lock()
			client.ConnectChan <- c.Id
			c.ConnectChan <- id
			return
		}
	}
}
