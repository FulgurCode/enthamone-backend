package server

import (
	"github.com/VAISHAKH-GK/atta-backend/internal/handlers"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Run(port string) {
	var app = fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/start", websocket.New(handlers.HandleConnection))

	app.Listen(":" + port)
}
