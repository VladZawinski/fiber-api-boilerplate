package main

import (
	"futbook/cmd/server"
	"futbook/pkg/config"
)

func main() {
	// setup various configuration for app
	config.LoadAllConfigs(".env")
	server.Serve()
}
