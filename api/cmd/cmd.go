package cmd

import (
	_ "github.com/rafaelbreno/go-api-template/api/cmd/config"
	router "github.com/rafaelbreno/go-api-template/api/routes"
)

// Boostrap
func Boostrap() {
	router.Listen()
}
