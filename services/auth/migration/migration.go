package migration

import (
	"log"

	"github.com/rafaelbreno/go-api-template/services/auth/config"
	"github.com/rafaelbreno/go-api-template/services/auth/entity"
)

func init() {
	log.Println("[Migrations] Running...")
	config.
		DB.
		AutoMigrate(entity.User{})

	log.Println("[Migrations] Done!")
}
