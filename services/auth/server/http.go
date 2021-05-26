package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// More https://docs.gofiber.io/api/fiber#config
func getConfig() fiber.Config {
	return fiber.Config{
		CaseSensitive: true,
		StrictRouting: false,
		Concurrency:   256 * 1024,
		WriteTimeout:  time.Duration(time.Second * 45), // 45 seconds
	}
}

var srv *fiber.App

func Listen() {
	srv = fiber.New()

	routes()

	srv.Listen(":3000")
}

func routes() {
	srv.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
