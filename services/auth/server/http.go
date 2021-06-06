package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-api-template/services/auth/auth"
	"github.com/rafaelbreno/go-api-template/services/auth/entity"
	"github.com/rafaelbreno/go-api-template/services/auth/handlers"
	"github.com/rafaelbreno/go-api-template/services/auth/repository"
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

var user entity.User

func routes() {

	apiGroup := srv.Group("/auth")

	apiGroup.Get("/health-check", func(c *fiber.Ctx) error {
		return c.
			Status(http.StatusOK).
			JSON(map[string]string{
				"message": "Everything working fine",
			})
	})
	apiGroup.Get("/ping", func(c *fiber.Ctx) error {
		return c.
			Status(http.StatusOK).
			JSON(map[string]string{
				"message": "Pong! Auth",
			})
	})

	uh := handlers.NewUserHandler()
	apiGroup.Post("/signup", uh.Create)
	apiGroup.Post("/signin", uh.SignIn)

	authGroup := srv.Group("", authMiddleware)

	authGroup.Get("/check", func(c *fiber.Ctx) error {
		return c.
			Status(http.StatusOK).
			JSON(map[string]interface{}{
				"token": c.Get("Authorization"),
				"user":  user.ToDTO(),
			})
	})

	authGroup.Get("/auth-health-check", func(c *fiber.Ctx) error {
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

	sql := repository.
		NewUserRepositoryDB().
		DB.
		First(&user, jwtClaim.ID)

	err = sql.Error

	if err != nil {
		return c.
			Status(http.StatusNotFound).
			JSON(map[string]string{
				"message": err.Error(),
			})
	}

	return c.Next()
}
