CREATE TABLE IF NOT EXISTS song(
    id BIGSERIAL PRIMARY KEY,
    group_name VARCHAR(150) NOT NULL,
    song_name VARCHAR(200) NOT NULL,
    release_date DATE NOT NULL,
    link VARCHAR(500) NOT NULL
);

CREATE TABLE IF NOT EXISTS song_verse(
    id BIGSERIAL PRIMARY KEY,
    order_num INT,
    verse TEXT NOT NULL,
    song_id INT,
    CONSTRAINT fk_song FOREIGN KEY (song_id)
    REFERENCES song (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);