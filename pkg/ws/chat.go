package ws

import (
	"github.com/VAISHAKH-GK/atta-backend/pkg/message"
)

func (c *Client) ListenMsg() {
	// Disconnecting client
	defer func() {
		c.Conn.Close()
		c.UnRegister <- true
	}()

	for {
		// Receiving messages
		var msg message.Message
		var err = c.Conn.ReadJSON(&msg)
		if err != nil {
			return
		}

		if msg.MessageType == message.CHAT {
			if client,exist := Clients[msg.To]; exist {
				client.MessageChan <- msg
			}
		}
	}
}
