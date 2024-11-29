package presentation

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"os/signal"
	"song-lib/internal/config"
	"syscall"
	"time"
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

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		sig := <-quit
		fmt.Printf("Shutting down server, signal: %s\n", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	fmt.Printf("Server starting on addr: %s\n", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	return nil
}
