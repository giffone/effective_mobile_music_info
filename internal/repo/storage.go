package repo

import (
	"context"
	"fmt"
	"music_info/internal/dto"
	"music_info/internal/model"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	GetInfoByGroupAndSong(ctx context.Context, filter dto.InfoByGroupAndSong) (model.SongDetail, error)
}

func New(pool *pgxpool.Pool) Storage {
	return &storage{pool: pool}
}

type storage struct {
	pool *pgxpool.Pool
}

func (s *storage) GetInfoByGroupAndSong(ctx context.Context, filter dto.InfoByGroupAndSong) (model.SongDetail, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT s.release_date, s.lyrics, s.link
FROM music.song s
JOIN music.group g ON s.group_id = g.id
WHERE g.group_name = $1 
AND s.song_name = $2;`

	var song model.SongDetail

	err := s.pool.QueryRow(ctx, query, filter.Group, filter.Song).Scan(
		&song.ReleaseDate,
		&song.Text,
		&song.Link,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return song, err
		}
		return song, fmt.Errorf("db: getInfoByGroupAndSong: %w", err)
	}

	return song, nil
}
