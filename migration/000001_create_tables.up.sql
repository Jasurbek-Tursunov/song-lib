CREATE TABLE IF NOT EXISTS song(
    id BIGSERIAL PRIMARY KEY,
    group VARCHAR(150) NOT NULL,
    song VARCHAR(200) NOT NULL,
    release_date DATE NOT NULL,
    link VARCHAR(500) NOT NULL
);

CREATE TABLE IF NOT EXISTS song_verse(
    id BIGSERIAL PRIMARY KEY,
    order INT,
    verse TEXT NOT NULL,
    song_id INT,
    CONSTRAINT fk_song FOREIGN KEY (song_id)
    REFERENCES songs (song_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);