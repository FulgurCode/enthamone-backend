package ws

import (
	"sync"

	"github.com/VAISHAKH-GK/enthamone-backend/pkg/message"
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

		client.mu.Lock()
		c.mu.Lock()
		if client.LookingConn && c.LookingConn {
			c.saveConnection(client)
			c.ConnectChan <- client.Id
			return
		}
		client.mu.Unlock()
		c.mu.Unlock()
	}

	var msg = message.Message{
		To: c.Id,
		MessageType: message.SIGNAL,
		Category: message.CONNECT_FAIL,
	}
	c.MessageChan <- msg
}

func (c *Client) saveConnection(client *Client) {
	c.LookingConn = false
	c.ConnectedUser = client.Id

	client.LookingConn = false
	client.ConnectedUser = c.Id

	c.mu.Unlock()
	client.mu.Unlock()
}

func (c *Client) deleteConnection() {
	c.mu.Lock()

	c.LookingConn = true
	c.ConnectedUser = ""

	c.mu.Unlock()
}
