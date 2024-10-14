package model

import "time"

const DateFormat = "02.01.2006"

type SongDetail struct {
	ReleaseDate time.Time
	Lyrics      string
	Link        string
}

type Song struct {
	GroupName string
	SongName  string
}

type UpdateSong struct {
	GroupID  int
	SongName string
	SongDetail
}
