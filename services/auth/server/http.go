package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-api-template/services/auth/auth"
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

func Listen() {
	srv = fiber.New(getConfig())

	routes()

	log.Fatal(srv.Listen(":8080"))
}

func routes() {
	srv.Get("/health-check", func(c *fiber.Ctx) error {
		return c.
			Status(http.StatusOK).
			JSON(map[string]string{
				"message": "Everything working fine",
			})
	})
	srv.Get("/ping", func(c *fiber.Ctx) error {
		return c.
			Status(http.StatusOK).
			JSON(map[string]string{
				"message": "Pong! Auth",
			})
	})

	uh := handlers.NewUserHandler()

	srv.Post("/signup", uh.Create)
	srv.Post("/signin", uh.SignIn)

	srv.Use(authMiddleware)

	srv.Get("/auth-health-check", func(c *fiber.Ctx) error {
		return c.
			Status(http.StatusOK).
			JSON(map[string]string{
				"message": "Everything working fine",
				"token":   c.Get("Authorization"),
			})
	})
}

func authMiddleware(c *fiber.Ctx) error {
	bearerToken := c.Get("Authorization")

	if bearerToken == "" {
		err := fmt.Errorf("Authorization key isn't set")
		return c.
			Status(http.StatusForbidden).
			JSON(map[string]string{
				"message": err.Error(),
			})
	}

	tokens := strings.Split(bearerToken, " ")

	token := tokens[len(tokens)-1]

	jwtWrapper := auth.Wrapper{
		Secret: "super-secret",
		Issuer: "AuthService",
	}

	jwtClaim, err := jwtWrapper.ValidateToken(token)
	if err != nil {
		return c.
			Status(http.StatusForbidden).
			JSON(map[string]string{
				"message": err.Error(),
			})
	}

	c.Set("username", jwtClaim.Username)

	return c.Next()
}
