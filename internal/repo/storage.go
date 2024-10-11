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

	query := ``

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
