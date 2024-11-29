package main

import (
	"song-lib/internal/config"
	"song-lib/internal/di"

	_ "song-lib/api/swagger"
)

// @title Song Library API
// @version 1.0
// @host localhost:8081
// @BasePath /
func main() {
	cfg := config.MustLoad()
	di.Init(cfg)
}
