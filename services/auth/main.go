package main

import (
	_ "github.com/rafaelbreno/go-api-template/services/auth/config"
	"github.com/rafaelbreno/go-api-template/services/auth/server"
)

func main() {
	server.Listen()
}
