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

	r.Use(authMiddleware)

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
