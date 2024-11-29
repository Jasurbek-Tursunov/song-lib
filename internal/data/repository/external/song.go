package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"song-lib/internal/domain/entity"
)

type SongInfoRepository struct {
	ExternalBaseURL string
}

func NewSongInfoRepository(url string) *SongInfoRepository {
	return &SongInfoRepository{ExternalBaseURL: url}
}

func (c *SongInfoRepository) GetInfo(in *entity.SongMinimal) (*entity.SongDetail, error) {
	const op = "data.external.SongInfoRepository"

	url := fmt.Sprintf("%s/info?group=%s&song=%s", c.ExternalBaseURL, in.Group, in.Song)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var songDetail entity.SongDetail
		if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return &songDetail, nil

	case http.StatusBadRequest:
		return nil, fmt.Errorf("%s: %w", op, entity.BadRequestError)

	default:
		return nil, fmt.Errorf("%s: %w", op, entity.InternalError)
	}
}
