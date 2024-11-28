package main

import (
	"song-lib/internal/config"
	"song-lib/internal/di"
)

func main() {
	cfg := config.MustLoad()
	di.Init(cfg)
}
