CREATE TABLE IF NOT EXISTS music.song (
	group_id INT NOT NULL,
    song_name VARCHAR(100) NOT NULL,
	lyrics TEXT DEFAULT '',
	link TEXT DEFAULT '',
	release_date DATE,
    CONSTRAINT song_pkey PRIMARY KEY (group_id, song_name),
    CONSTRAINT group_fkey FOREIGN KEY (group_id) REFERENCES music.group(id) ON DELETE CASCADE
);

-- Permissions

-- ALTER TABLE music.song OWNER TO postgres;
-- GRANT ALL ON TABLE music.song TO postgres;