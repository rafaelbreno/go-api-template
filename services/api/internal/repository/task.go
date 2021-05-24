package repository

import (
	"github.com/rafaelbreno/go-api-template/api/cmd/storage"
	"github.com/rafaelbreno/go-api-template/api/internal/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	FindAll() ([]entity.Task, error)
	FindByID(id string) (entity.Task, error)
	Create(task entity.Task) (entity.Task, error)
}

type TaskRepositoryDB struct {
	DBConn *gorm.DB
}

func (tr TaskRepositoryDB) FindAll() ([]entity.Task, error) {
	var tasks []entity.Task

	if err := tr.DBConn.Find(&tasks).Error; err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (tr TaskRepositoryDB) FindByID(id string) (entity.Task, error) {
	var task entity.Task

	if err := tr.DBConn.First(&task, id).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (tr TaskRepositoryDB) Create(t entity.Task) (entity.Task, error) {
	if result := tr.DBConn.Create(&t); result.Error != nil {
		return entity.Task{}, result.Error
	}

	return t, nil
}

func NewTaskRepositoryDB() TaskRepositoryDB {
	return TaskRepositoryDB{storage.DBConn}
}
