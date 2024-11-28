package database

import (
	"context"
	"database/sql"
	"errors"
	"song-lib/internal/domain/entity"
	"time"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) List(filters entity.Filters) ([]*entity.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, group, song, release_date, link 
           WHERE (LOWER(group)=LOWER($1) OR $1 = '')
			AND (LOWER(song)=LOWER($2) OR $2 = '')
			AND (release_date=$3 OR $3 = '')
			AND (LOWER(link)=LOWER($4) OR $4 = '')`

	args := []any{
		filters.Group,
		filters.Song,
		filters.ReleaseDate,
		filters.Link,
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*entity.Song

	for rows.Next() {
		var song entity.Song
		err := rows.Scan(
			&song.ID,
			&song.Group,
			&song.Song,
			&song.ReleaseDate,
			&song.Link,
		)

		if err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}
	return songs, nil
}

func (r *SongRepository) Get(id int) (*entity.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, group, song, release_date, link FROM song WHERE id = $1`

	r.db.QueryRowContext(ctx, query, id)
	return &entity.Song{}, nil
}

func (r *SongRepository) Create(in *entity.CreateSong) (*entity.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO song(group, song, release_date, link) VALUES ($1, $2, $3, $4) RETURNING id`

	args := []any{
		in.Group,
		in.Song,
		in.ReleaseDate,
		in.Link,
	}
	out := entity.Song{
		Group:       in.Group,
		Song:        in.Song,
		ReleaseDate: in.ReleaseDate,
		Link:        in.Link,
	}

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&out.ID)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *SongRepository) Update(id int, in *entity.Song) (*entity.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE song SET group=$1, song=$2, release_date=$3, link=$4 WHERE id=$5`

	args := []any{
		in.Group,
		in.Song,
		in.ReleaseDate,
		in.Link,
		id,
	}

	r.db.ExecContext(ctx, query, args...)
	return in, nil
}

func (r *SongRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM song WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("record not found")
	}
	return nil
}

type SongVerseRepository struct {
	db *sql.DB
}

func NewSongVerseRepository(db *sql.DB) *SongVerseRepository {
	return &SongVerseRepository{db: db}
}

func (r *SongVerseRepository) Create(in *entity.CreateSongVerse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO song_text(song_id, order, verse) VALUES ($1, $2, $3) RETURNING id"

	args := []any{
		in.SongID,
		in.Order,
		in.Verse,
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("record not found")
	}

	return nil
}
