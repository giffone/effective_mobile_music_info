package repo

import (
	"context"
	"fmt"
	"log"
	"music_info/internal/model"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	CreateSong(ctx context.Context, song model.Song) (err error)
	UpdateSong(ctx context.Context, song model.UpdateSong) error
	GetInfoByGroupAndSong(ctx context.Context, filter model.Song) (model.SongDetail, error)
}

func New(pool *pgxpool.Pool) Storage {
	return &storage{pool: pool}
}

type storage struct {
	pool *pgxpool.Pool
}

func (s *storage) CreateSong(ctx context.Context, song model.Song) (err error) {
	ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	queryGroup := `INSERT INTO
music.group (group_name)
VALUES ($1)
ON CONFLICT (group_name)
DO UPDATE SET
group_name = EXCLUDED.group_name
RETURNING id;`

	querySong := `INSERT INTO
music.song (group_id, song_name)
VALUES ($1, $2)
ON CONFLICT (group_id, song_name)
DO NOTHING;`

	var groupID int

	tx, err := s.pool.Begin(ctx2)
	if err != nil {
		return fmt.Errorf("db: createSong: tx: %w", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx2); rbErr != nil {
				log.Printf("db: createSong: rollback err: %s", rbErr)
			}
		}
	}()

	// create group
	if err = tx.QueryRow(ctx2, queryGroup, song.GroupName).Scan(&groupID); err != nil {
		return fmt.Errorf("db: createSong: create group: %w", err)
	}

	// create song
	if _, err = tx.Exec(ctx2, querySong, groupID, song.SongName); err != nil {
		return fmt.Errorf("db: createSong: create song: %w", err)
	}

	// fix changes
	if err = tx.Commit(ctx2); err != nil {
		return fmt.Errorf("db: createSong: tx: commit: %w", err)
	}

	return nil
}

func (s *storage) UpdateSong(ctx context.Context, song model.UpdateSong) error {
	ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `UPDATE music.song
SET release_date = $1,
lyrics = $2,
link = $3
WHERE group_id = $4
AND song_name = $5;`

	// update song
	if _, err := s.pool.Exec(ctx2, query,
		song.ReleaseDate,
		song.Lyrics,
		song.Link,
		song.GroupID,
		song.SongName,
	); err != nil {
		return fmt.Errorf("db: updateSong: exec: %w", err)
	}

	return nil
}

func (s *storage) GetInfoByGroupAndSong(ctx context.Context, filter model.Song) (model.SongDetail, error) {
	ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT s.release_date, s.lyrics, s.link
FROM music.song s
JOIN music.group g ON s.group_id = g.id
WHERE g.group_name = $1 
AND s.song_name = $2;`

	var song model.SongDetail

	err := s.pool.QueryRow(ctx2, query, filter.GroupName, filter.SongName).Scan(
		&song.ReleaseDate,
		&song.Lyrics,
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
