package main

import (
	_ "github.com/rafaelbreno/go-api-template/services/auth/config"
	_ "github.com/rafaelbreno/go-api-template/services/auth/migration"
	"github.com/rafaelbreno/go-api-template/services/auth/server"
)

func main() {
	server.Listen()
}
