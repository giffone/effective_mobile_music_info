package service

import (
	"context"
	"music_info/internal/dto"
	"music_info/internal/repo"
	"music_info/internal/view"
)

type Service interface {
	GetInfoByGroupAndSong(ctx context.Context, filter dto.InfoByGroupAndSong) (view.SongDetail, error)
}

func New(storage repo.Storage) Service {
	return &service{storage: storage}
}

type service struct {
	storage repo.Storage
}

func (s *service) GetInfoByGroupAndSong(ctx context.Context, filter dto.InfoByGroupAndSong) (view.SongDetail, error) {
	// can add validate text max length or other

	// get song
	song, err := s.storage.GetInfoByGroupAndSong(ctx, filter)
	if err != nil {
		return view.SongDetail{}, err
	}

	// prepare view
	return view.SongDetail{
		ReleaseDate: song.ReleaseDate.Format("02.01.2006"),
		Text:        song.Text,
		Link:        song.Link,
	}, nil
}
