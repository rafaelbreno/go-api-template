package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/go-api-template/api/auth"
	"github.com/rafaelbreno/go-api-template/api/internal/entity"
	"github.com/rafaelbreno/go-api-template/api/internal/repository"
	"github.com/rafaelbreno/go-api-template/api/utils"
	"gorm.io/gorm"
)

type taskHandler struct {
	repo repository.TaskRepositoryDB
	user auth.UserDTO
}

func NewTaskHandler(userDTO auth.UserDTO) taskHandler {
	return taskHandler{
		repo: repository.NewTaskRepositoryDB(),
		user: userDTO,
	}
}

func (t taskHandler) FindAll(c *gin.Context) {
	tasks, err := t.repo.FindAll()

	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": tasks,
	})
}

func (t taskHandler) FindById(c *gin.Context) {
	id := c.Param("id")

	task, err := t.repo.FindByID(id)

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
		"data": task,
	})
	return
}

func (t taskHandler) Create(c *gin.Context) {
	var taskInput entity.Task

	if err := c.ShouldBindJSON(&taskInput); err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	task, err := t.repo.Create(taskInput)

	if err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": task,
	})
	return
}

func (t taskHandler) Update(c *gin.Context) {
	var taskInput entity.Task
	if err := c.ShouldBindJSON(&taskInput); err != nil {
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

	taskInput.ID = id

	task, err := t.repo.Update(taskInput)

	if err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": task,
	})
	return
}

func (t taskHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	task, err := t.repo.Delete(id)

	if err != nil {
		c.JSON(402, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": task,
	})
	return
}
