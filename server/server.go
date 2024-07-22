package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/websocket"
)

func Run(port string) {
	var app = fiber.New()

	app.Get("/test", func (c *fiber.Ctx) error {
		return c.SendString("test route")
	})

	app.Use("/ws", func (c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Listen(":" + port)
}
