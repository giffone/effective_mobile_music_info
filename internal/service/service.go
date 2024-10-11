package service

import "music_info/internal/repo"

type Service interface {
}

func New(storage repo.Storage) Service {
	return &service{storage: storage}
}

type service struct {
	storage repo.Storage
}
