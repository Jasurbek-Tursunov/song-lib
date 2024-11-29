package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"song-lib/internal/domain/entity"
	"time"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) List(filters entity.Filters, paginate entity.Paginator) ([]*entity.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, group_name, song_name, release_date, link FROM song
           WHERE (LOWER(group_name)=LOWER($1) OR $1 = '')
			AND (LOWER(song_name)=LOWER($2) OR $2 = '')
			AND (release_date=CAST(NULLIF($3, '') AS DATE) OR $3 = '')
			AND (LOWER(link)=LOWER($4) OR $4 = '')
			ORDER BY id ASC
			LIMIT $5 OFFSET $6`

	args := []any{
		filters.Group,
		filters.Song,
		filters.ReleaseDate,
		filters.Link,
		paginate.PageSize,
		(paginate.Page - 1) * paginate.PageSize,
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

	query := `SELECT id, group_name, song_name, release_date, link FROM song WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)

	var song entity.Song
	err := row.Scan(
		&song.ID,
		&song.Group,
		&song.Song,
		&song.ReleaseDate,
		&song.Link,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.NotFoundError
		}
		return nil, err
	}

	return &song, nil
}

func (r *SongRepository) Create(in *entity.CreateSong) (*entity.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO song(group_name, song_name, release_date, link) VALUES ($1, $2, $3, $4) RETURNING id`

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

	query := `UPDATE song SET group_name=$1, song_name=$2, release_date=$3, link=$4 WHERE id=$5`

	args := []any{
		in.Group,
		in.Song,
		in.ReleaseDate,
		in.Link,
		id,
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return nil, entity.NotFoundError
	}

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
		return entity.NotFoundError
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

	query := "INSERT INTO song_verse(song_id, order_num, verse) VALUES ($1, $2, $3) RETURNING id"

	args := []any{
		in.SongID,
		in.Order,
		in.Verse,
	}

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *SongVerseRepository) GetText(songID int, paginate entity.Paginator) ([]*entity.SongVerse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, order_num, verse FROM song_verse
           WHERE song_id=$1 ORDER BY order_num ASC LIMIT $2 OFFSET $3`

	args := []any{
		songID,
		paginate.PageSize,
		(paginate.Page - 1) * paginate.PageSize,
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var text []*entity.SongVerse

	for rows.Next() {
		var verse entity.SongVerse
		err := rows.Scan(
			&verse.ID,
			&verse.Order,
			&verse.Verse,
		)

		if err != nil {
			return nil, err
		}
		text = append(text, &verse)
	}
	return text, nil
}
