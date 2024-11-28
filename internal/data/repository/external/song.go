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
	url := fmt.Sprintf("%s/info?group=%s&song=%s", c.ExternalBaseURL, in.Group, in.Song)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var songDetail entity.SongDetail
		if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
			return nil, err
		}
		return &songDetail, nil

	case http.StatusBadRequest:
		return nil, BadRequestError

	default:
		return nil, InternalServerError
	}
}
