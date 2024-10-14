package service

import (
	"context"
	"errors"
	"music_info/internal/dto"
	"music_info/internal/model"
	"music_info/internal/repo"
	"music_info/internal/view"
	"strconv"
	"time"
)

type Service interface {
	CreateSong(ctx context.Context, song dto.Song) error
	UpdateSong(ctx context.Context, song dto.UpdateSong) error
	GetInfoByGroupAndSong(ctx context.Context, filter dto.Song) (view.SongDetail, error)
}

func New(storage repo.Storage) Service {
	return &service{storage: storage}
}

type service struct {
	storage repo.Storage
}

func (s *service) CreateSong(ctx context.Context, song dto.Song) error {
	// can add validate text max length or other
	if song.Empty() {
		return errors.Join(model.ErrBadData, errors.New("empty data"))
	}

	songValidated := model.Song{
		GroupName: song.Group,
		SongName:  song.Song,
	}

	// create song
	return s.storage.CreateSong(ctx, songValidated)
}

func (s *service) UpdateSong(ctx context.Context, song dto.UpdateSong) error {
	// can add validate text max length or other
	if song.Empty() {
		return errors.Join(model.ErrBadData, errors.New("empty data"))
	}

	// check group id
	groupID, err := strconv.Atoi(song.GroupID)
	if err != nil || groupID == 0 {
		return errors.Join(model.ErrBadData, err)
	}

	// check date
	t, err := time.Parse(model.DateFormat, song.ReleaseDate)
	if err != nil {
		return errors.Join(model.ErrBadData, err)
	}

	songValidated := model.UpdateSong{
		GroupID:  groupID,
		SongName: song.Song,
		SongDetail: model.SongDetail{
			ReleaseDate: t,
			Lyrics:      song.Text,
			Link:        song.Link,
		},
	}

	// create song
	return s.storage.UpdateSong(ctx, songValidated)
}

func (s *service) GetInfoByGroupAndSong(ctx context.Context, filter dto.Song) (view.SongDetail, error) {
	// can add validate text max length or other
	if filter.Empty() {
		return view.SongDetail{}, errors.Join(model.ErrBadData, errors.New("empty data"))
	}

	filterValidated := model.Song{
		GroupName: filter.Group,
		SongName:  filter.Song,
	}

	// get song
	song, err := s.storage.GetInfoByGroupAndSong(ctx, filterValidated)
	if err != nil {
		return view.SongDetail{}, err
	}

	// prepare view
	return view.SongDetail{
		ReleaseDate: song.ReleaseDate.Format(model.DateFormat),
		Text:        song.Lyrics,
		Link:        song.Link,
	}, nil
}
