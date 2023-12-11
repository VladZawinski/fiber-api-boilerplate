package main

import (
	"fiber-api-boilerplate/cmd/server"
	"fiber-api-boilerplate/pkg/config"
)

func main() {
	// setup various configuration for app
	config.LoadAllConfigs(".env")
	server.Serve()
}
