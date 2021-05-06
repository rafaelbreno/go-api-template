package handler

import (
	"github.com/gin-gonic/gin"
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
