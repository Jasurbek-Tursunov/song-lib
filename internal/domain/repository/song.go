package repository

import "song-lib/internal/domain/entity"

type SongRepository interface {
	List(filters entity.Filters, paginate entity.Paginator) ([]*entity.Song, error)
	Get(id int) (*entity.Song, error)
	Create(in *entity.CreateSong) (*entity.Song, error)
	Update(id int, in *entity.Song) (*entity.Song, error)
	Delete(id int) error
}

type SongVerseRepository interface {
	Create(in *entity.CreateSongVerse) error
	GetText(songID int, paginate entity.Paginator) ([]*entity.SongVerse, error)
}

type SongInfoRepository interface {
	GetInfo(in *entity.SongMinimal) (*entity.SongDetail, error)
}
