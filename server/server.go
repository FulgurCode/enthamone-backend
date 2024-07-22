package server

import "github.com/gofiber/fiber/v2"

func Run(port string) {
	var app = fiber.New()

	app.Get("/test", func (c *fiber.Ctx) error {
		return c.SendString("test route")
	})

	app.Listen(":" + port)
}
