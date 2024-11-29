package presentation

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	swagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

func (h *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(CorsMiddleware)

	router.Get("/swagger/*", swagger.WrapHandler)

	router.Route("/songs", func(r chi.Router) {
		r.Get("/", h.List)
		r.Post("/", h.Create)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.Get)
			r.Put("/", h.Update)
			r.Delete("/", h.Delete)
			r.Get("/text", h.GetText)
		})
	})

	return router
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
