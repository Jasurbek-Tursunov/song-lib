package presentation

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"song-lib/internal/domain/entity"
	"song-lib/internal/domain/usecase"
	"strconv"
)

type Handler struct {
	service usecase.SongService
}

func NewHandler(service usecase.SongService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	filters := entity.Filters{}

	songList, err := h.service.List(filters)
	if err != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(songList)
	if err != nil {

	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 0, 64)
	if err != nil {

	}

	song, err := h.service.Get(int(id))
	if err != nil {
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(song)
	if err != nil {
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var songMinimal entity.SongMinimal
	err := json.NewDecoder(r.Body).Decode(&songMinimal)
	if err != nil {

	}

	song, err := h.service.Create(&songMinimal)
	if err != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(song)
	if err != nil {
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 0, 64)
	if err != nil {

	}

	var song *entity.Song
	err = json.NewDecoder(r.Body).Decode(song)
	if err != nil {
	}

	song, err = h.service.Update(int(id), song)
	if err != nil {
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(song)
	if err != nil {
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 0, 64)
	if err != nil {

	}

	err = h.service.Delete(int(id))

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(nil)
	if err != nil {

	}
}
