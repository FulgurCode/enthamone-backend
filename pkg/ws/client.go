package ws

import "github.com/gofiber/contrib/websocket"

var Clients map[string]*Client

type Client struct {
	Conn *websocket.Conn
}

func NewClient(conn *websocket.Conn,id string) *Client {
	var c = Client{
		Conn: conn,
	}

	Clients[id] = &c

	return &c
}
