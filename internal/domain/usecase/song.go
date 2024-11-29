package usecase

import "song-lib/internal/domain/entity"

type SongService interface {
	List(filters entity.Filters, paginate entity.Paginator) ([]*entity.Song, error)
	Get(id int) (*entity.Song, error)
	GetText(songID int, paginate entity.Paginator) ([]*entity.SongVerse, error)
	Create(in *entity.SongMinimal) (*entity.Song, error)
	Update(id int, in *entity.Song) (*entity.Song, error)
	Delete(id int) error
}
