package repository

import (
	"github.com/rafaelbreno/go-api-template/api/cmd/storage"
	"github.com/rafaelbreno/go-api-template/api/internal/entity"
	"gorm.io/gorm"
)

type ListRepository interface {
	FindAll() ([]entity.List, error)
	FindByID(id string, user_id string) (entity.List, error)
	Create(list entity.List) (entity.List, error)
	Update(list entity.List) (entity.List, error)
	Delete(id string, user_id string) (entity.List, error)
}

type ListRepositoryDB struct {
	DBConn *gorm.DB
}

func (lr ListRepositoryDB) FindAll(user_id string) ([]entity.List, error) {
	var lists []entity.List

	if err := lr.DBConn.Where("user_id = ?", user_id).Find(&lists).Error; err != nil {
		return lists, err
	}

	return lists, nil
}

func (lr ListRepositoryDB) FindByID(id string, user_id string) (entity.List, error) {
	var list entity.List

	if err := lr.DBConn.First(&list, "user_id = ? AND id = ?", user_id, id).Error; err != nil {
		return list, err
	}

	if err := lr.DBConn.Where("list_id = ?", list.ID).Find(&list.Tasks).Error; err != nil {
		return list, err
	}

	return list, nil
}

func (lr ListRepositoryDB) Create(l entity.List) (entity.List, error) {
	if err := lr.DBConn.Create(&l).Error; err != nil {
		return entity.List{}, err
	}
	return l, nil
}

func (lr ListRepositoryDB) Update(l entity.List) (entity.List, error) {
	var list entity.List

	if err := lr.DBConn.First(&list, "user_id = ? AND id = ?", l.UserID, l.ID).Error; err != nil {
		return list, err
	}

	if l.Title != "" {
		list.Title = l.Title
	}

	if l.Description != "" {
		list.Description = l.Description
	}

	if l.Status != list.Status {
		list.Status = l.Status
	}

	if err := lr.DBConn.Save(&list).Error; err != nil {
		return entity.List{}, err
	}

	return list, nil
}

func (lr ListRepositoryDB) Delete(id string, user_id string) (entity.List, error) {
	var list entity.List

	if err := lr.DBConn.First(&list, "user_id = ? AND id = ?", user_id, id).Error; err != nil {
		return list, err
	}

	if err := lr.DBConn.Delete(&list, id).Error; err != nil {
		return list, err
	}

	return list, nil
}
func NewListRepositoryDB() ListRepositoryDB {
	return ListRepositoryDB{
		DBConn: storage.DBConn,
	}
}
