package repository

import (
	"github.com/rafaelbreno/go-api-template/api/auth"
	"github.com/rafaelbreno/go-api-template/api/cmd/storage"
	"github.com/rafaelbreno/go-api-template/api/internal/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	FindAll() ([]entity.Task, error)
	FindByID(id string) (entity.Task, error)
	Create(task entity.Task) (entity.Task, error)
	Update(task entity.Task) (entity.Task, error)
	Delete(id string) (entity.Task, error)
}

type TaskRepositoryDB struct {
	DBConn *gorm.DB
	User   auth.UserDTO
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
	var tempList entity.List

	if err := tr.DBConn.First(&tempList, t.ListID).Error; err != nil {
		return entity.Task{}, err
	}

	if err := tr.DBConn.Create(&t).Error; err != nil {
		return entity.Task{}, err
	}

	return t, nil
}

func (tr TaskRepositoryDB) Update(t entity.Task) (entity.Task, error) {
	var task entity.Task

	if err := tr.DBConn.First(&task, t.ID).Error; err != nil {
		return task, err
	}

	if t.Title == "" {
		task.Title = t.Title
	}

	if t.Description == "" {
		task.Description = t.Description
	}

	if t.Status != task.Status {
		task.Status = t.Status
	}

	if result := tr.DBConn.Save(&task); result.Error != nil {
		return entity.Task{}, result.Error
	}

	return t, nil
}

func (tr TaskRepositoryDB) Delete(id string) (entity.Task, error) {
	var task entity.Task

	if err := tr.DBConn.Delete(&task, id).Error; err != nil {
		return task, err
	}

	return task, nil
}

func NewTaskRepositoryDB(userDTO auth.UserDTO) TaskRepositoryDB {
	return TaskRepositoryDB{
		DBConn: storage.DBConn,
		User:   userDTO,
	}
}
