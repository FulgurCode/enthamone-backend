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
			if msg.Category == message.CONNECT_REQ {
				c.NewConnection()
			} else if msg.Category == message.DISCONNECT_REQ {
				c.DisconnectChan <- true
				Clients[c.ConnectedUser].DisconnectChan <- true
			} else if msg.Category == message.ICE_SIGNAL {
				Clients[msg.To].MessageChan <- msg
			}

		} else if msg.MessageType == message.OFFER {
			Clients[msg.To].MessageChan <- msg
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
				continue
			}

		case _ = <-c.UnRegister:
			if !c.LookingConn {
				Clients[c.ConnectedUser].DisconnectChan <- true
			}
			delete(Clients, c.Id)
			c.Conn.Close()
			return

		case id := <-c.ConnectChan:
			var msg = message.Message{
				To:          c.Id,
				MessageType: message.SIGNAL,
				Category:    message.CONNECT_SIGNAL,
				Content:     id,
			}

			c.Conn.WriteJSON(msg)

		case _ = <-c.DisconnectChan:
			c.deleteConnection()

			var msg = message.Message{
				To:          c.Id,
				MessageType: message.SIGNAL,
				Category:    message.DISCONNECT_SIGNAL,
				Content:     "DISCONNECT",
			}
			var err = c.Conn.WriteJSON(msg)
			if err != nil {
				continue
			}
		}
	}
}
