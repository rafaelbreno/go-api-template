package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-api-template/services/auth/handlers"
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
var r fiber.Router

func Listen() {
	srv = fiber.New(getConfig())

	r = srv.Group("/auth")

	routes()

	log.Fatal(srv.Listen(":3000"))
}

func routes() {
	srv.Get("/health-check", func(c *fiber.Ctx) error {
		return c.
			Status(http.StatusOK).
			JSON(map[string]string{
				"message": "Everything working fine",
			})
	})

	uh := handlers.NewUserHandler()

	r.Post("/signup", uh.Create)
	r.Post("/signin", uh.SignIn)
}
