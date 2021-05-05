package cmd

import (
	_ "github.com/rafaelbreno/go-api-template/api/cmd/config"
	"github.com/rafaelbreno/go-api-template/api/cmd/storage"
	"github.com/rafaelbreno/go-api-template/api/internal/entity"
	router "github.com/rafaelbreno/go-api-template/api/routes"
)

// Boostrap
func Boostrap() {
	migration()

	router.Listen()
}

// Migrate tables
func migration() {
	storage.
		Migrator(entity.Task{})
}
