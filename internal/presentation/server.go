package presentation

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"song-lib/internal/config"
)

type Server struct {
	router *chi.Mux
	cfg    *config.Config
}

func NewServer(router *chi.Mux, cfg *config.Config) *Server {
	return &Server{router: router, cfg: cfg}
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	const op = "http.Server.Run"
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", s.cfg.Port),
		Handler: s.router,
	}

	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	fmt.Printf("Server started on addr: %s", srv.Addr)
	return nil
}
