package usecase

import "song-lib/internal/domain/entity"

type SongService interface {
	List(filters entity.Filters) ([]*entity.Song, error)
	Get(id int) (*entity.Song, error)
	Create(in *entity.SongMinimal) (*entity.Song, error)
	Update(id int, in *entity.Song) (*entity.Song, error)
	Delete(id int) error
}
