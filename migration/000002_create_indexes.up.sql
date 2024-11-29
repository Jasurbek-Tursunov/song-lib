CREATE INDEX IF NOT EXISTS song_id_idx ON song (id);
CREATE INDEX IF NOT EXISTS song_verse_id_idx ON song_verse(id);
CREATE INDEX IF NOT EXISTS song_verse_song_id_idx ON song_verse(song_id);   