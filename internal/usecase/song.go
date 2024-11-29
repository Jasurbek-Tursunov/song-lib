package usecase

import (
	"fmt"
	"song-lib/internal/domain/entity"
	"song-lib/internal/domain/repository"
	"strings"
	"time"
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

func (s *SongService) List(filters entity.Filters, paginate entity.Paginator) ([]*entity.Song, error) {
	return s.SongRepository.List(filters, paginate)
}

func (s *SongService) Get(id int) (*entity.Song, error) {
	return s.SongRepository.Get(id)
}

func (s *SongService) GetText(songID int, paginate entity.Paginator) ([]*entity.SongVerse, error) {
	return s.SongVerseRepository.GetText(songID, paginate)
}

func (s *SongService) Create(in *entity.SongMinimal) (*entity.Song, error) {
	const op = "SongService.Create"

	detail, err := s.SongInfoRepository.GetInfo(in)
	if err != nil {
		return nil, err
	}

	releaseDate, err := time.Parse("02.01.2006", detail.ReleaseDate)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	text := strings.Split(detail.Text, "\n\n")
	create := entity.CreateSong{
		Group:       in.Group,
		Song:        in.Song,
		ReleaseDate: releaseDate,
		Link:        detail.Link,
	}

	song, err := s.SongRepository.Create(&create)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, value := range text {
		verse := entity.CreateSongVerse{
			SongID: song.ID,
			Order:  i + 1,
			Verse:  value,
		}

		err = s.SongVerseRepository.Create(&verse)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
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
