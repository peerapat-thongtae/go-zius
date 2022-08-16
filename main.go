package main

import (
	"example.com/go-zius/configs"
	"example.com/go-zius/routes"
)

func main() {
	// Connect DB
	configs.ConnectDB()

	// Init route V1
	routes.InitV1Route()
}
