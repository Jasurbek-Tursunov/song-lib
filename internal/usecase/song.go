package usecase

import (
	"song-lib/internal/domain/entity"
	"song-lib/internal/domain/repository"
	"strings"
)

type SongService struct {
	SongRepository      repository.SongRepository
	SongVerseRepository repository.SongVerseRepository
	SongInfoRepository  repository.SongInfoRepository
}

func NewSongService(
	song repository.SongRepository,
	verse repository.SongVerseRepository,
	info repository.SongInfoRepository,
) *SongService {
	return &SongService{
		SongRepository:      song,
		SongVerseRepository: verse,
		SongInfoRepository:  info,
	}
}

func (s *SongService) List(filters entity.Filters) ([]*entity.Song, error) {
	return s.SongRepository.List(filters)
}

func (s *SongService) Get(id int) (*entity.Song, error) {
	return s.SongRepository.Get(id)
}

func (s *SongService) Create(in *entity.SongMinimal) (*entity.Song, error) {
	detail, err := s.SongInfoRepository.GetInfo(in)
	if err != nil {
		return nil, err
	}

	text := strings.Split(detail.Text, "\n\n")
	create := entity.CreateSong{
		Group:       in.Group,
		Song:        in.Song,
		ReleaseDate: detail.ReleaseDate,
		Link:        detail.Link,
	}

	song, err := s.SongRepository.Create(&create)
	if err != nil {
		return nil, err
	}

	for i, value := range text {
		verse := entity.CreateSongVerse{
			SongID: song.ID,
			Order:  i + 1,
			Verse:  value,
		}
		err = s.SongVerseRepository.Create(&verse)
		if err != nil {
		}
	}

	return song, nil
}

func (s *SongService) Update(id int, in *entity.Song) (*entity.Song, error) {
	return s.SongRepository.Update(id, in)
}

func (s *SongService) Delete(id int) error {
	return s.SongRepository.Delete(id)
}
