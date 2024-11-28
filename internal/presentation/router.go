package presentation

import (
	"github.com/go-chi/chi/v5"
)

func (h *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/songs", func(r chi.Router) {
		r.Get("/", h.List)
		r.Post("/", h.Create)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.Get)
			r.Put("/", h.Update)
			r.Delete("/", h.Delete)
		})
	})
	return router
}
