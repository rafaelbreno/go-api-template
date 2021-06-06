package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/go-api-template/api/auth"
	"github.com/rafaelbreno/go-api-template/api/internal/entity"
	"github.com/rafaelbreno/go-api-template/api/internal/repository"
	"github.com/rafaelbreno/go-api-template/api/utils"
	"gorm.io/gorm"
)

type listHandler struct {
	repo repository.ListRepositoryDB
	user auth.UserDTO
}

func NewListHandler(userDTO auth.UserDTO) listHandler {
	return listHandler{
		repo: repository.NewListRepositoryDB(),
		user: userDTO,
	}
}

func (l listHandler) FindAll(c *gin.Context) {
	lists, err := l.repo.FindAll()

	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": lists,
	})
}

func (l listHandler) FindByID(c *gin.Context) {
	lists, err := l.repo.FindByID(c.Param("id"))

	if err == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": lists,
	})
}

func (l listHandler) Create(c *gin.Context) {
	var listInput entity.List

	if err := c.ShouldBindJSON(&listInput); err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	list, err := l.repo.Create(listInput)

	if err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": list,
	})
	return
}

func (l listHandler) Update(c *gin.Context) {
	var listInput entity.List
	if err := c.ShouldBindJSON(&listInput); err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := utils.StringToUint(c.Param("id"))

	if err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	listInput.ID = id

	list, err := l.repo.Update(listInput)

	if err == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": list,
	})
	return
}

func (l listHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	list, err := l.repo.Delete(id)

	if err == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": list,
	})
	return
}
