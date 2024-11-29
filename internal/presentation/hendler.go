package presentation

import (
	"encoding/json"
	"errors"
	"net/http"
	"song-lib/internal/domain/entity"
	"song-lib/internal/domain/usecase"
	"song-lib/pkg/http/rest"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service usecase.SongService
}

func NewHandler(service usecase.SongService) *Handler {
	return &Handler{service: service}
}

// @Summary List
// @Tags song
// @Description list songs data
// @Accept json
// @Produce json
// @Param group query string  false  "Filter by group name"
// @Param song query string  false  "Filter by song name"
// @Param release-date query string  false  "Filter by release-date name"
// @Param link query string  false  "Filter by link name"
// @Param limit query int false  "Limit for paginate"
// @Param page query int false  "Page for paginate"
// @Success 200 {object} entity.Songs
// @Failure 400 {object} rest.Err
// @Failure 500 {object} rest.Err
// @Router /songs [get]
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	filters := entity.Filters{
		Group:       params.Get("group"),
		Song:        params.Get("song"),
		ReleaseDate: params.Get("release-date"),
		Link:        params.Get("link"),
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	page, err := strconv.Atoi(params.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	paginate := entity.Paginator{
		PageSize: limit,
		Page:     page,
	}

	songList, err := h.service.List(filters, paginate)
	if err != nil {
		rest.ErrorResponse(w, r, 500, err)
		return
	}

	rest.Encode(w, &entity.Songs{Songs: songList})
}

// @Summary Get
// @Tags song
// @Description get song data
// @Accept json
// @Produce json
// @Param id path int  true  "Song ID"
// @Success 200 {object} entity.Song
// @Failure 404 {object} rest.Err
// @Failure 500 {object} rest.Err
// @Router /songs/{id} [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		rest.NotFoundResponse(w, r)
		return
	}

	song, err := h.service.Get(id)
	if err != nil {
		if errors.Is(err, entity.NotFoundError) {
			rest.NotFoundResponse(w, r)
			return
		}
		rest.ErrorResponse(w, r, 500, err)
		return
	}

	rest.Encode(w, song)
}

// @Summary Song list verse
// @Tags song-text
// @Description Song list verse
// @Accept json
// @Produce json
// @Param id path int  true  "Song ID"
// @Param limit query int false  "Limit for paginate"
// @Param page query int false  "Page for paginate"
// @Success 200 {object} entity.SongText
// @Failure 404 {object} rest.Err
// @Failure 400 {object} rest.Err
// @Failure 500 {object} rest.Err
// @Router /songs/{id}/text [get]
func (h *Handler) GetText(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		rest.NotFoundResponse(w, r)
		return
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	page, err := strconv.Atoi(params.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	paginate := entity.Paginator{
		PageSize: limit,
		Page:     page,
	}

	verses, err := h.service.GetText(id, paginate)
	if err != nil {
		rest.ErrorResponse(w, r, 500, err)
		return
	}

	rest.Encode(w, &entity.SongText{Text: verses})
}

// @Summary Create
// @Tags song
// @Description create song
// @Accept json
// @Produce json
// @Param input body entity.SongMinimal  true  "Song struct"
// @Success 200 {object} entity.Song
// @Failure 400 {object} rest.Err
// @Failure 404 {object} rest.Err
// @Failure 500 {object} rest.Err
// @Router /songs [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var songMinimal entity.SongMinimal
	err := json.NewDecoder(r.Body).Decode(&songMinimal)
	if err != nil {
		rest.BadRequestResponse(w, r, err)
		return
	}

	song, err := h.service.Create(&songMinimal)
	if err != nil {
		rest.ErrorResponse(w, r, 500, err)
		return
	}

	rest.Encode(w, song)
}

// @Summary Update
// @Tags song
// @Description update song data
// @Accept json
// @Produce json
// @Param id path int  true  "Song ID"
// @Param input body entity.Song  true  "Song struct"
// @Success 200 {object} entity.Song
// @Failure 400 {object} rest.Err
// @Failure 404 {object} rest.Err
// @Failure 500 {object} rest.Err
// @Router /songs/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		rest.NotFoundResponse(w, r)
		return
	}

	var song entity.Song
	err = json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		rest.BadRequestResponse(w, r, err)
		return
	}

	songNew, err := h.service.Update(id, &song)
	if err != nil {
		if errors.Is(err, entity.NotFoundError) {
			rest.NotFoundResponse(w, r)
			return
		}
		rest.ErrorResponse(w, r, 500, err)
		return
	}

	rest.Encode(w, songNew)
}

// @Summary Delete
// @Tags song
// @Description delete song data
// @Accept json
// @Produce json
// @Param id path int  true  "Song ID"
// @Success 200
// @Failure 404 {object} rest.Err
// @Failure 500 {object} rest.Err
// @Router /songs/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		rest.NotFoundResponse(w, r)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		if errors.Is(err, entity.NotFoundError) {
			rest.NotFoundResponse(w, r)
			return
		}
		rest.ErrorResponse(w, r, 500, err)
		return
	}

	rest.Encode(w, nil)
}
