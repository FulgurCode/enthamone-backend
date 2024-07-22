package handlers

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gofiber/contrib/websocket"
)

// Handle websocket connection:w
func HandleConnection(c *websocket.Conn) {
	defer func() {
		c.Conn.Close()
	}()

	var id = strconv.Itoa(rand.Int())
	c.WriteMessage(1, []byte(id))

	for {
		var _, msg, err = c.ReadMessage()
		if err != nil {
			return
		}

		fmt.Println(string(msg))
	}
}
