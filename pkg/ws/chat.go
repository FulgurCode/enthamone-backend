package ws

import (
	"github.com/VAISHAKH-GK/atta-backend/pkg/message"
)

func (c *Client) ListenMsg() {
	// Disconnecting client
	defer func() {
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
			if client, exist := Clients[msg.To]; exist {
				client.MessageChan <- msg
			}
		} else if msg.MessageType == message.SIGNAL {
			if msg.Category == message.DISCONNECTED_SIGNAL {
				c.mu.Lock()
				c.ConnectedUser = ""
				c.LookingConn = true
				c.mu.Unlock()
				Clients[c.ConnectedUser].DisconnectChan <- true
			}
		}
	}
}

// Send message to client
func (c *Client) WriteMsg() {
	for {
		select {
		case msg := <-c.MessageChan:
			var err = c.Conn.WriteJSON(msg)
			if err != nil {
				return
			}
		case _ = <-c.UnRegister:
			if !c.LookingConn {
				Clients[c.ConnectedUser].DisconnectChan <- true
			}
			delete(Clients,c.Id)
			c.Conn.Close()
			return
		case id := <-c.ConnectChan:
			c.LookingConn = false
			c.ConnectedUser = id
			c.mu.Unlock()

			var msg = message.Message{
				To:          c.Id,
				MessageType: message.SIGNAL,
				Category:    message.CONNECTED_SIGNAL,
				Content:     id,
			}

			c.Conn.WriteJSON(msg)
		case _ = <-c.DisconnectChan:
			c.mu.Lock()
			c.ConnectedUser = ""
			c.LookingConn = true
			c.mu.Unlock()
		}
	}
}
