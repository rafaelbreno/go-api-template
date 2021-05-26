package migration

import (
	"github.com/rafaelbreno/go-api-template/services/auth/config"
	"github.com/rafaelbreno/go-api-template/services/auth/entity"
)

func init() {
	config.
		DB.
		AutoMigrate(entity.User{})
}
