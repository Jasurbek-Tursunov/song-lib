package di

import (
	"song-lib/internal/config"
	"song-lib/internal/data/repository/database"
	"song-lib/internal/data/repository/external"
	"song-lib/internal/presentation"
	"song-lib/internal/usecase"
	"song-lib/pkg/postgres"
)

func Init(cfg *config.Config) {
	dbCfg := postgres.Config{
		Host:     cfg.DbHost,
		Port:     cfg.DbPort,
		Name:     cfg.DbName,
		User:     cfg.DbUser,
		Password: cfg.DbPassword,
	}
	db := postgres.MustLoad(&dbCfg)

	// DATA Layer
	songRepository := database.NewSongRepository(db)
	songVerseRepository := database.NewSongVerseRepository(db)
	songInfoRepository := external.NewSongInfoRepository("")

	// DOMAIN Layer
	songService := usecase.NewSongService(songRepository, songVerseRepository, songInfoRepository)

	// PRESENTATION Layer
	handler := presentation.NewHandler(songService)
	server := presentation.NewServer(handler.Routes(), cfg)
	server.MustRun()
}
