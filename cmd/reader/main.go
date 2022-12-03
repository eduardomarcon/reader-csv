package main

import (
	"reader/internal/configs"
	server "reader/pkg"
)

func main() {
	configs.LoadEnvs()
	server.UP()
}
