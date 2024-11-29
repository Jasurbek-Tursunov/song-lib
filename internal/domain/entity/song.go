package entity

import "time"

type SongMinimal struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type CreateSongVerse struct {
	SongID int
	Order  int
	Verse  string
}

type CreateSong struct {
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate time.Time `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

type Song struct {
	ID          int       `json:"id"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate time.Time `json:"releaseDate"`
	Link        string    `json:"link"`
}

type Songs struct {
	Songs []*Song `json:"songs"`
}

type SongVerse struct {
	ID    int    `json:"id"`
	Order int    `json:"order"`
	Verse string `json:"verse"`
}

type SongText struct {
	Text []*SongVerse `json:"text"`
}

type Filters struct {
	Group       string
	Song        string
	ReleaseDate string
	Link        string
}

type Paginator struct {
	PageSize int
	Page     int
}
