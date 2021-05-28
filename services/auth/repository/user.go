package repository

import (
	"github.com/rafaelbreno/go-api-template/services/auth/config"
	"github.com/rafaelbreno/go-api-template/services/auth/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
}

type UserRepositoryDB struct {
	DB *gorm.DB
}

func (ur UserRepositoryDB) Create(user entity.User) (entity.User, error) {
	if err := ur.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepositoryDB() UserRepositoryDB {
	return UserRepositoryDB{config.DB}
}
