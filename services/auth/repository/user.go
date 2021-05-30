package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/rafaelbreno/go-api-template/services/auth/config"
	"github.com/rafaelbreno/go-api-template/services/auth/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	SignIn(user entity.User) (entity.User, error)
}

type UserRepositoryDB struct {
	DB  *gorm.DB
	Rdb *redis.Client
}

func (ur UserRepositoryDB) Create(user entity.User) (entity.User, error) {
	if err := ur.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur UserRepositoryDB) SignIn(user entity.User) (entity.User, error) {
	var userDB entity.User

	if err := ur.DB.Where("username = ?", user.Username).First(&userDB).Error; err != nil {
		return entity.User{}, err
	}

	if err := user.CheckPassword(userDB.Password); err != nil {
		return user, err
	}

	return userDB, nil
}

func NewUserRepositoryDB() UserRepositoryDB {
	return UserRepositoryDB{
		DB:  config.DB,
		Rdb: config.Rdb,
	}
}
