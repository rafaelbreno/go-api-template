package migration

import (
	"fmt"

	"github.com/rafaelbreno/go-api-template/services/auth/config"
	"github.com/rafaelbreno/go-api-template/services/auth/entity"
)

func init() {
	fmt.Println("[Migrations] Running...")
	config.
		DB.
		AutoMigrate(entity.User{})

	fmt.Println("[Migrations] Done!")
}
