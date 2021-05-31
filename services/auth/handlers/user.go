package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/go-api-template/services/auth/auth"
	"github.com/rafaelbreno/go-api-template/services/auth/entity"
	"github.com/rafaelbreno/go-api-template/services/auth/repository"
)

type userHandler struct {
	repo repository.UserRepositoryDB
}

func NewUserHandler() userHandler {
	return userHandler{repository.NewUserRepositoryDB()}
}

func (u userHandler) Create(c *fiber.Ctx) error {
	var user entity.User

	if err := c.BodyParser(&user); err != nil {
		c.
			Status(http.StatusServiceUnavailable).
			JSON(map[string]string{
				"message": err.Error(),
			})
		return err
	}

	if err := user.EncryptPassword(); err != nil {
		c.
			Status(http.StatusUnprocessableEntity).
			JSON(map[string]string{
				"message": err.Error(),
			})
		return err
	}

	user, err := u.repo.Create(user)
	if err != nil {
		c.
			Status(http.StatusUnprocessableEntity).
			JSON(map[string]string{
				"message": err.Error(),
			})
		return err
	}

	userAuth, err := auth.GetUserToken(user)
	if err != nil {
		c.
			Status(http.StatusUnprocessableEntity).
			JSON(map[string]string{
				"message": err.Error(),
			})
		return err
	}

	c.
		Status(http.StatusCreated).
		JSON(userAuth)
	return nil
}

func (u userHandler) SignIn(c *fiber.Ctx) error {
	var user entity.User
	var err error

	if err = c.BodyParser(&user); err != nil {
		c.
			Status(http.StatusServiceUnavailable).
			JSON(map[string]string{
				"message": err.Error(),
			})
		return err
	}

	if user, err = u.repo.SignIn(user); err != nil {
		c.
			Status(http.StatusNotFound).
			JSON(map[string]string{
				"message": err.Error(),
			})
		return err
	}

	userAuth, err := auth.GetUserToken(user)
	if err != nil {
		c.
			Status(http.StatusUnprocessableEntity).
			JSON(map[string]string{
				"message": err.Error(),
			})
		return err
	}

	c.
		Status(http.StatusCreated).
		JSON(userAuth)

	return nil
}
