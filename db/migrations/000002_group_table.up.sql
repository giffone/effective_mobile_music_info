CREATE TABLE IF NOT EXISTS music.group (
	id SERIAL,
	group_name varchar(100) NOT NULL,
	CONSTRAINT group_pkey PRIMARY KEY (id)
);

-- Permissions

-- ALTER TABLE music.group OWNER TO postgres;
-- GRANT ALL ON TABLE music.group TO postgres;