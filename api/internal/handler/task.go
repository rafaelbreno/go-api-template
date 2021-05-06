package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelbreno/go-api-template/api/internal/entity"
	"github.com/rafaelbreno/go-api-template/api/internal/repository"
)

type taskHandler struct {
	repo repository.TaskRepositoryDB
}

func NewTaskHandler() taskHandler {
	return taskHandler{repository.NewTaskRepositoryDB()}
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
